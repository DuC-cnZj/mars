package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"reflect"
	"time"
	"unsafe"

	"helm.sh/helm/v3/pkg/downloader"

	"github.com/DuC-cnZj/mars/pkg/mlog"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/storage/driver"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var tmpFileDir = "/tmp"

type DeleteFunc func()

func WriteConfigYamlToTmpFile(data []byte) (string, DeleteFunc, error) {
	fileName := "mars_" + RandomString(20) + ".yaml"
	var fullPath = fmt.Sprintf("%s/%s", tmpFileDir, fileName)
	if FileExists(fullPath) {
		return "", nil, errors.New(fmt.Sprintf("file %s already exists", fullPath))
	}

	openFile, err := os.OpenFile(fullPath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		return "", nil, err
	}
	defer openFile.Close()
	if _, err := openFile.Write(data); err != nil {
		return "", nil, err
	}

	return fullPath, func() {
		mlog.Warning("delete file: " + fullPath)
		if err := os.Remove(fullPath); err != nil {
			mlog.Error("WriteConfigYamlToTmpFile error: ", err)
		}
	}, nil
}

func EncodeConfigToYaml(field string, data interface{}) ([]byte, error) {
	bf := &bytes.Buffer{}
	encoder := yaml.NewEncoder(bf)
	if err := encoder.Encode(map[string]interface{}{
		field: data,
	}); err != nil {
		return nil, err
	}

	return bf.Bytes(), nil
}

func GetSettings(namespace string) *cli.EnvSettings {
	s := cli.New()
	n := (*string)(unsafe.Pointer(s))
	*n = namespace
	config := (*genericclioptions.ConfigFlags)(unsafe.Pointer(s))
	v := namespace
	reflect.ValueOf(config).Elem().FieldByName("Namespace").Set(reflect.ValueOf(&v))
	s.Debug = App().IsDebug()
	mlog.Warningf("%#v", s)

	return s
}

// UpgradeOrInstall TODO
func UpgradeOrInstall(releaseName, namespace string, ch *chart.Chart, valueOpts *values.Options, fn func(format string, v ...interface{})) (*release.Release, error) {
	var settings = GetSettings(namespace)
	mlog.Warning("settings ns", settings.Namespace())
	actionConfig := new(action.Configuration)
	flags := genericclioptions.NewConfigFlags(true)
	flags.Namespace = &namespace

	if Config().KubeConfig != "" {
		*flags.KubeConfig = Config().KubeConfig
	} else {
		host, port := os.Getenv("KUBERNETES_SERVICE_HOST"), os.Getenv("KUBERNETES_SERVICE_PORT")
		settings.KubeAPIServer = "https://" + net.JoinHostPort(host, port)
		token, _ := ioutil.ReadFile(tokenFile)
		settings.KubeToken = string(token)
	}

	if err := actionConfig.Init(flags, namespace, "", fn); err != nil {
		return nil, err
	}
	client := action.NewUpgrade(actionConfig)
	client.Atomic = true
	client.WaitForJobs = true
	client.Install = true
	client.Timeout = 2 * time.Minute
	client.Namespace = namespace
	if valueOpts == nil {
		valueOpts = &values.Options{}
	}

	client.Namespace = namespace
	if client.Install {
		// If a release does not exist, install it.
		histClient := action.NewHistory(actionConfig)
		histClient.Max = 1
		if _, err := histClient.Run(releaseName); err == driver.ErrReleaseNotFound {
			instClient := action.NewInstall(actionConfig)
			instClient.CreateNamespace = true
			instClient.ChartPathOptions = client.ChartPathOptions
			instClient.DryRun = client.DryRun
			instClient.DisableHooks = client.DisableHooks
			instClient.SkipCRDs = client.SkipCRDs
			instClient.Timeout = client.Timeout
			instClient.Wait = client.Wait
			instClient.WaitForJobs = client.WaitForJobs
			instClient.Devel = client.Devel
			instClient.Namespace = client.Namespace
			instClient.Atomic = client.Atomic
			instClient.PostRenderer = client.PostRenderer
			instClient.DisableOpenAPIValidation = client.DisableOpenAPIValidation
			instClient.SubNotes = client.SubNotes
			instClient.Description = client.Description
			mlog.Debug("start install release", valueOpts)
			return runInstall(releaseName, ch, instClient, valueOpts, settings)
		}
	}

	if client.Version == "" && client.Devel {
		mlog.Debug("setting version to >0.0.0-0")
		client.Version = ">0.0.0-0"
	}

	vals, err := valueOpts.MergeValues(getter.All(settings))
	if err != nil {
		return nil, err
	}

	if req := ch.Metadata.Dependencies; req != nil {
		if err := action.CheckDependencies(ch, req); err != nil {
			return nil, err
		}
	}

	if ch.Metadata.Deprecated {
		mlog.Warning("This chart is deprecated")
	}

	return client.Run(releaseName, ch, vals)
}

func UninstallRelease(releaseName, namespace string) error {
	settings := GetSettings(namespace)
	actionConfig := new(action.Configuration)

	if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), "", mlog.Debugf); err != nil {
		return err
	}

	uninstall := action.NewUninstall(actionConfig)
	if _, err := uninstall.Run(releaseName); err != nil {
		mlog.Error(err)
		return err
	}
	return nil
}

func runInstall(releaseName string, chartRequested *chart.Chart, client *action.Install, valueOpts *values.Options, settings *cli.EnvSettings) (*release.Release, error) {
	mlog.Debugf("Original chart version: %q", client.Version)
	if client.Version == "" && client.Devel {
		mlog.Debug("setting version to >0.0.0-0")
		client.Version = ">0.0.0-0"
	}

	client.ReleaseName = releaseName

	p := getter.All(settings)
	vals, err := valueOpts.MergeValues(p)
	if err != nil {
		return nil, err
	}

	// Check chart dependencies to make sure all are present in /charts
	if err := checkIfInstallable(chartRequested); err != nil {
		return nil, err
	}

	if chartRequested.Metadata.Deprecated {
		mlog.Warning("This chart is deprecated")
	}

	client.Namespace = settings.Namespace()
	return client.Run(chartRequested, vals)
}

func checkIfInstallable(ch *chart.Chart) error {
	switch ch.Metadata.Type {
	case "", "application":
		return nil
	}
	return errors.Errorf("%s charts are not installable", ch.Metadata.Type)
}

const (
	StatusUnknown  string = "unknown"
	StatusDeployed string = "deployed"
	StatusFailed   string = "failed"
)

var (
	tokenFile  = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	rootCAFile = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
)

func ReleaseStatus(releaseName, namespace string) (string, error) {
	var settings = GetSettings(namespace)

	mlog.Warning("settings ns", settings.Namespace())
	actionConfig := new(action.Configuration)
	flags := genericclioptions.NewConfigFlags(true)
	flags.Namespace = &namespace
	if Config().KubeConfig != "" {
		*flags.KubeConfig = Config().KubeConfig
	} else {
		host, port := os.Getenv("KUBERNETES_SERVICE_HOST"), os.Getenv("KUBERNETES_SERVICE_PORT")
		settings.KubeAPIServer = "https://" + net.JoinHostPort(host, port)
		token, _ := ioutil.ReadFile(tokenFile)
		settings.KubeToken = string(token)
	}

	if err := actionConfig.Init(flags, namespace, "", log.Printf); err != nil {
		return "", err
	}
	statusClient := action.NewStatus(actionConfig)
	run, err := statusClient.Run(releaseName)
	if err != nil {
		mlog.Error(err)
		return "", err
	}

	switch run.Info.Status {
	case release.StatusDeployed:
		return StatusDeployed, nil
	case release.StatusFailed:
		return StatusFailed, nil
	default:
		return StatusUnknown, nil
	}
}

func PackageChart(path string, destDir string) (string, error) {
	var settings = GetSettings("")

	mlog.Warning("settings ns", settings.Namespace())
	actionConfig := new(action.Configuration)
	flags := genericclioptions.NewConfigFlags(true)
	if Config().KubeConfig != "" {
		*flags.KubeConfig = Config().KubeConfig
	} else {
		host, port := os.Getenv("KUBERNETES_SERVICE_HOST"), os.Getenv("KUBERNETES_SERVICE_PORT")
		settings.KubeAPIServer = "https://" + net.JoinHostPort(host, port)
		token, _ := ioutil.ReadFile(tokenFile)
		settings.KubeToken = string(token)
	}

	if err := actionConfig.Init(flags, "", "", log.Printf); err != nil {
		mlog.Error(err)
		return "", err
	}
	newPackage := action.NewPackage()
	if destDir != "" {
		newPackage.Destination = destDir
	}

	// 更新依赖 dependency, 防止没有依赖文件打包失败
	downloadManager := &downloader.Manager{
		Out:              ioutil.Discard,
		ChartPath:        path,
		Keyring:          newPackage.Keyring,
		Getters:          getter.All(settings),
		Debug:            settings.Debug,
		RepositoryConfig: settings.RepositoryConfig,
		RepositoryCache:  settings.RepositoryCache,
	}

	if err := downloadManager.Update(); err != nil {
		return "", err
	}

	return newPackage.Run(path, nil)
}

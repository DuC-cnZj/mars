package socket

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/gosimple/slug"
	"go.uber.org/config"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/release"

	"github.com/duc-cnzj/mars-client/v4/mars"
	"github.com/duc-cnzj/mars-client/v4/types"
	websocket_pb "github.com/duc-cnzj/mars-client/v4/websocket"
	app "github.com/duc-cnzj/mars/internal/app/helper"
	"github.com/duc-cnzj/mars/internal/contracts"
	"github.com/duc-cnzj/mars/internal/event/events"
	"github.com/duc-cnzj/mars/internal/mlog"
	"github.com/duc-cnzj/mars/internal/models"
	"github.com/duc-cnzj/mars/internal/plugins"
	"github.com/duc-cnzj/mars/internal/utils"
	"github.com/duc-cnzj/mars/internal/utils/recovery"
)

const (
	ResultError             = websocket_pb.ResultType_Error
	ResultSuccess           = websocket_pb.ResultType_Success
	ResultDeployed          = websocket_pb.ResultType_Deployed
	ResultDeployFailed      = websocket_pb.ResultType_DeployedFailed
	ResultDeployCanceled    = websocket_pb.ResultType_DeployedCanceled
	ResultLogWithContainers = websocket_pb.ResultType_LogWithContainers

	WsSetUid             = websocket_pb.Type_SetUid
	WsReloadProjects     = websocket_pb.Type_ReloadProjects
	WsCancel             = websocket_pb.Type_CancelProject
	WsCreateProject      = websocket_pb.Type_CreateProject
	WsUpdateProject      = websocket_pb.Type_UpdateProject
	WsProcessPercent     = websocket_pb.Type_ProcessPercent
	WsClusterInfoSync    = websocket_pb.Type_ClusterInfoSync
	WsInternalError      = websocket_pb.Type_InternalError
	WsHandleExecShell    = websocket_pb.Type_HandleExecShell
	WsHandleExecShellMsg = websocket_pb.Type_HandleExecShellMsg
	WsHandleCloseShell   = websocket_pb.Type_HandleCloseShell
	WsAuthorize          = websocket_pb.Type_HandleAuthorize
	ProjectPodEvent      = websocket_pb.Type_ProjectPodEvent

	// Maximum message size allowed from peer.
	maxMessageSize = 1024 * 1024 * 10 // 10MB
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

var reloadProjectsMessage = &websocket_pb.WsMetadataResponse{Metadata: &websocket_pb.Metadata{Type: WsReloadProjects}}

type WsResponse = websocket_pb.WsMetadataResponse

type SafeWriteMessageCh struct {
	closeable utils.Closeable
	ch        chan contracts.MessageItem
}

func (s *SafeWriteMessageCh) Closed() {
	mlog.Debug("SafeWriteMessageCh closed")
	if s.closeable.Close() {
		close(s.ch)
	}
}

func (s *SafeWriteMessageCh) Chan() <-chan contracts.MessageItem {
	return s.ch
}
func (s *SafeWriteMessageCh) Send(m contracts.MessageItem) {
	if s.closeable.IsClosed() {
		mlog.Debugf("[Websocket]: Drop %s type %s", m.Msg, m.Type)
		return
	}
	s.ch <- m
}

type vars map[string]any

func (v vars) MustGetString(key string) string {
	if v != nil {
		if value, ok := v[key]; ok {
			return fmt.Sprintf("%v", value)
		}
	}

	return ""
}

type Jober struct {
	helmer contracts.Helmer
	done   chan struct{}

	loaders []Loader

	dryRun    bool
	manifests []string

	input    *websocket_pb.CreateProjectInput
	wsType   websocket_pb.Type
	slugName string

	destroyFuncLock sync.RWMutex
	destroyFuncs    []func()

	imagePullSecrets  []string
	vars              vars
	dynamicConfigYaml string
	extraValues       []string
	valuesYaml        string
	chart             *chart.Chart
	valuesOptions     *values.Options
	installer         contracts.ReleaseInstaller
	commit            contracts.CommitInterface

	messageCh contracts.SafeWriteMessageChInterface
	stopCtx   context.Context
	stopOnce  sync.Once
	stopFn    func(error)

	isNew       bool
	config      *mars.Config
	ns          *models.Namespace
	project     *models.Project
	prevProject *models.Project

	percenter contracts.Percentable
	messager  contracts.DeployMsger

	pubsub contracts.PubSub

	user           contracts.UserInfo
	timeoutSeconds int64
}

type Option func(*Jober)

func WithDryRun() Option {
	return func(j *Jober) {
		j.dryRun = true
	}
}

type NewJobFunc func(
	input *websocket_pb.CreateProjectInput,
	user contracts.UserInfo,
	slugName string,
	messager contracts.DeployMsger,
	pubsub contracts.PubSub,
	timeoutSeconds int64,
	opts ...Option,
) contracts.Job

func NewJober(
	input *websocket_pb.CreateProjectInput,
	user contracts.UserInfo,
	slugName string,
	messager contracts.DeployMsger,
	pubsub contracts.PubSub,
	timeoutSeconds int64,
	opts ...Option,
) contracts.Job {
	jb := &Jober{
		helmer:         &DefaultHelmer{},
		loaders:        defaultLoaders(),
		done:           make(chan struct{}),
		user:           user,
		pubsub:         pubsub,
		messager:       messager,
		vars:           vars{},
		valuesOptions:  &values.Options{},
		slugName:       slugName,
		input:          input,
		wsType:         input.Type,
		timeoutSeconds: timeoutSeconds,
		messageCh:      &SafeWriteMessageCh{ch: make(chan contracts.MessageItem, 100)},
		percenter:      newProcessPercent(messager, &realSleeper{}),
	}
	jb.stopCtx, jb.stopFn = utils.NewCustomErrorContext()
	for _, opt := range opts {
		opt(jb)
	}

	return jb
}

func (j *Jober) Manifests() []string {
	return j.manifests
}

func (j *Jober) ID() string {
	return j.slugName
}

func (j *Jober) Logs() []string {
	return j.ReleaseInstaller().Logs()
}

func (j *Jober) IsNew() bool {
	return j.isNew
}

func (j *Jober) IsDryRun() bool {
	return j.dryRun
}

func (j *Jober) Commit() contracts.CommitInterface {
	return j.commit
}

func (j *Jober) User() contracts.UserInfo {
	return j.user
}

func (j *Jober) ProjectModel() *types.ProjectModel {
	if j.project == nil {
		return nil
	}
	return j.project.ProtoTransform()
}

func (j *Jober) Project() *models.Project {
	return j.project
}

func (j *Jober) Namespace() *models.Namespace {
	return j.ns
}

func (j *Jober) Done() <-chan struct{} {
	return j.done
}

func (j *Jober) Finish() {
	mlog.Debug("finished")
	close(j.done)
	j.PubSub().ToAll(reloadProjectsMessage)
}

func (j *Jober) Prune() {
	if j.IsNew() && !j.IsDryRun() {
		mlog.Debug("清理项目")
		app.DB().Delete(&j.project)
	}
}

func (j *Jober) Stop(err error) {
	j.stopOnce.Do(func() {
		mlog.Debugf("stop deploy job, because '%v'", err)
		j.messager.SendMsg("收到取消信号, 开始停止部署~")
		j.stopFn(err)
	})
}

func (j *Jober) IsStopped() bool {
	select {
	case <-j.stopCtx.Done():
		return true
	default:
	}

	return false
}

func (j *Jober) HandleMessage() {
	defer mlog.Debug("HandleMessage exit")
	ch := j.messageCh.Chan()
	for {
		select {
		case <-j.Done():
			return
		case <-app.App().Done():
			return
		case s, ok := <-ch:
			if !ok {
				return
			}
			switch s.Type {
			case contracts.MessageText:
				j.Messager().SendMsgWithContainerLog(s.Msg, s.Containers)
			case contracts.MessageError:
				if j.IsNew() && !j.IsDryRun() {
					app.DB().Delete(&j.project)
				}
				select {
				case <-j.stopCtx.Done():
					j.Messager().SendDeployedResult(ResultDeployCanceled, j.stopCtx.Err().Error(), j.project.ProtoTransform())
				default:
					j.Messager().SendDeployedResult(ResultDeployFailed, s.Msg, j.project.ProtoTransform())
				}
				return
			case contracts.MessageSuccess:
				j.Messager().SendDeployedResult(ResultDeployed, s.Msg, j.project.ProtoTransform())
				return
			}
		}
	}
}

func (j *Jober) CallDestroyFuncs() {
	j.destroyFuncLock.RLock()
	defer j.destroyFuncLock.RUnlock()
	for _, destroyFunc := range j.destroyFuncs {
		destroyFunc()
	}
}

func (j *Jober) AddDestroyFunc(fn func()) {
	j.destroyFuncLock.Lock()
	defer j.destroyFuncLock.Unlock()
	j.destroyFuncs = append(j.destroyFuncs, fn)
}

type userConfig struct {
	Config           string              `yaml:"config"`
	Branch           string              `yaml:"branch"`
	Commit           string              `yaml:"commit"`
	Atomic           bool                `yaml:"atomic"`
	WebUrl           string              `yaml:"web_url"`
	Title            string              `yaml:"title"`
	ExtraValues      []*types.ExtraValue `yaml:"extra_values"`
	FinalExtraValues mergeYamlString     `yaml:"final_extra_values"`
	EnvValues        vars                `yaml:"env_values"`
}

type mergeYamlString []string

func (s mergeYamlString) MarshalYAML() (any, error) {
	var opts []config.YAMLOption
	for _, item := range s {
		opts = append(opts, config.Source(strings.NewReader(item)))
	}
	if len(opts) == 0 {
		return "", nil
	}
	provider, _ := config.NewYAML(opts...)
	var merged map[string]any
	provider.Get("").Populate(&merged)

	bf := &bytes.Buffer{}
	yaml.NewEncoder(bf).Encode(&merged)

	return bf.String(), nil
}

type sortableExtraItem []*types.ExtraValue

func (s sortableExtraItem) Len() int {
	return len(s)
}

func (s sortableExtraItem) Less(i, j int) bool {
	return s[i].Path < s[j].Path
}

func (s sortableExtraItem) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (u userConfig) PrettyYaml() string {
	bf := bytes.Buffer{}
	sort.Sort(sortableExtraItem(u.ExtraValues))
	yaml.NewEncoder(&bf).Encode(&u)
	return bf.String()
}

func toUpdatesMap(p *models.Project) map[string]any {
	return map[string]any{
		"manifest":           p.Manifest,
		"config":             p.Config,
		"git_project_id":     p.GitProjectId,
		"git_commit":         p.GitCommit,
		"git_branch":         p.GitBranch,
		"docker_image":       p.DockerImage,
		"pod_selectors":      p.PodSelectors,
		"override_values":    p.OverrideValues,
		"atomic":             p.Atomic,
		"extra_values":       p.ExtraValues,
		"final_extra_values": p.FinalExtraValues,
		"env_values":         p.EnvValues,
		"deploy_status":      p.DeployStatus,
		"git_commit_title":   p.GitCommitTitle,
		"git_commit_web_url": p.GitCommitWebUrl,
		"git_commit_author":  p.GitCommitAuthor,
		"git_commit_date":    p.GitCommitDate,
		"config_type":        p.ConfigType,
	}
}

func (j *Jober) Run() error {
	done := make(chan struct{}, 1)
	go func() {
		defer func() {
			done <- struct{}{}
			close(done)
		}()
		defer recovery.HandlePanic("[Websocket]: Jober Run")
		j.HandleMessage()
	}()

	err := func() error {
		defer j.messageCh.Closed()
		var (
			result *release.Release
			err    error
		)

		if result, err = j.ReleaseInstaller().Run(j.stopCtx, j.messageCh, j.Percenter(), j.IsNew(), j.Commit().GetTitle()); err != nil {
			mlog.Errorf("[Websocket]: %v", err)
			j.messageCh.Send(contracts.MessageItem{
				Msg:  err.Error(),
				Type: contracts.MessageError,
			})
		} else {
			coalesceValues, _ := chartutil.CoalesceValues(j.ReleaseInstaller().Chart(), result.Config)
			j.project.OverrideValues, _ = coalesceValues.YAML()

			j.manifests = utils.SplitManifests(result.Manifest)
			j.project.Manifest = result.Manifest
			j.project.SetPodSelectors(getPodSelectorsByManifest(j.manifests))
			j.project.DockerImage = matchDockerImage(pipelineVars{
				Pipeline: j.vars.MustGetString("Pipeline"),
				Commit:   j.vars.MustGetString("Commit"),
				Branch:   j.vars.MustGetString("Branch"),
			}, result.Manifest)
			j.project.GitCommitTitle = j.Commit().GetTitle()
			j.project.GitCommitWebUrl = j.Commit().GetWebURL()
			j.project.GitCommitAuthor = j.Commit().GetAuthorName()
			j.project.GitCommitDate = j.Commit().GetCommittedDate()
			j.project.ConfigType = j.config.GetConfigFileType()

			if len(j.input.ExtraValues) > 0 {
				marshal, _ := json.Marshal(j.input.ExtraValues)
				j.project.ExtraValues = string(marshal)
			}
			if len(j.extraValues) > 0 {
				marshal, _ := json.Marshal(j.extraValues)
				j.project.FinalExtraValues = string(marshal)
			}
			if len(j.vars) > 0 {
				marshal, _ := json.Marshal(j.vars)
				j.project.EnvValues = string(marshal)
			}

			var (
				p                *models.Project
				oldConf, newConf userConfig
			)
			if j.prevProject != nil && j.prevProject.ID > 0 {
				p = j.prevProject
				j.project.ID = p.ID
				if !j.IsDryRun() {
					app.DB().Model(j.project).Updates(toUpdatesMap(j.project))
				}
				var (
					ev  []*types.ExtraValue
					fev mergeYamlString
				)
				if len(p.ExtraValues) > 0 {
					json.Unmarshal([]byte(p.ExtraValues), &ev)
				}
				if len(p.FinalExtraValues) > 0 {
					json.Unmarshal([]byte(p.FinalExtraValues), &fev)
				}

				oldConf = userConfig{
					Config:           p.Config,
					Branch:           p.GitBranch,
					Commit:           p.GitCommit,
					Atomic:           p.Atomic,
					ExtraValues:      ev,
					FinalExtraValues: fev,
					EnvValues:        vars(p.GetEnvValues()),
				}
				commit, err := plugins.GetGitServer().GetCommit(fmt.Sprintf("%d", p.GitProjectId), p.GitCommit)
				if err == nil {
					oldConf.WebUrl = commit.GetWebURL()
					oldConf.Title = commit.GetTitle()
				}
			} else {
				if !j.IsDryRun() {
					app.DB().Model(j.project).Updates(toUpdatesMap(j.project))
				}
			}
			newConf = userConfig{
				Config:           j.project.Config,
				Branch:           j.project.GitBranch,
				Commit:           j.project.GitCommit,
				Atomic:           j.project.Atomic,
				WebUrl:           j.project.GitCommitWebUrl,
				Title:            j.project.GitCommitTitle,
				ExtraValues:      j.input.ExtraValues,
				FinalExtraValues: mergeYamlString(j.extraValues),
				EnvValues:        j.vars,
			}
			if !j.IsDryRun() {
				app.Event().Dispatch(events.EventProjectChanged, &events.ProjectChangedData{
					Project:  j.project,
					Username: j.User().Name,
				})
			}
			var act types.EventActionType = types.EventActionType_Create
			if !j.IsNew() {
				act = types.EventActionType_Update
			}
			if j.IsDryRun() {
				act = types.EventActionType_DryRun
			}
			AuditLogWithChange(j.User().Name, act,
				fmt.Sprintf("%s 项目: %s/%s", act.String(), j.Namespace().Name, j.Project().Name),
				oldConf, newConf)
			j.Percenter().To(100)
			j.messageCh.Send(contracts.MessageItem{
				Msg:  "部署成功",
				Type: contracts.MessageSuccess,
			})
		}
		return err
	}()
	<-done
	return err
}

func (j *Jober) GetStoppedErrorIfHas() error {
	if j.IsStopped() {
		return j.stopCtx.Err()
	}
	return nil
}

func (j *Jober) ReleaseInstaller() contracts.ReleaseInstaller {
	return j.installer
}

func (j *Jober) Messager() contracts.DeployMsger {
	return j.messager
}

func (j *Jober) PubSub() contracts.PubSub {
	return j.pubsub
}

func (j *Jober) Percenter() contracts.Percentable {
	return j.percenter
}

func (j *Jober) Validate() error {
	if !(j.wsType == websocket_pb.Type_CreateProject || j.wsType == websocket_pb.Type_UpdateProject || j.wsType == websocket_pb.Type_ApplyProject) {
		return errors.New("type error: " + j.wsType.String())
	}
	j.Messager().SendMsg("[Start]: 收到请求，开始创建项目")
	j.Percenter().To(5)

	j.Messager().SendMsg("[Check]: 校验名称空间...")

	var ns models.Namespace
	if err := app.DB().Where("`id` = ?", j.input.NamespaceId).First(&ns).Error; err != nil {
		return fmt.Errorf("[FAILED]: 校验名称空间: %w", err)
	}
	j.ns = &ns

	j.Messager().SendMsg("[Loading]: 加载用户配置")
	j.Percenter().To(10)

	marsC, err := utils.GetProjectMarsConfig(j.input.GitProjectId, j.input.GitBranch)
	if err != nil {
		return err
	}
	j.config = marsC
	if j.input.Name == "" {
		j.input.Name = utils.GetProjectName(j.input.GitProjectId, marsC)
	}

	j.project = &models.Project{
		Name:         slug.Make(j.input.Name),
		GitProjectId: int(j.input.GitProjectId),
		GitBranch:    j.input.GitBranch,
		GitCommit:    j.input.GitCommit,
		Config:       j.input.Config,
		NamespaceId:  ns.ID,
		Atomic:       j.input.Atomic,
		ConfigType:   marsC.ConfigFileType,
	}

	j.Messager().SendMsg("[Check]: 检查项目是否存在")

	j.AddDestroyFunc(func() {
		mlog.Debug("update DeployStatus in DestroyFunc")
		if !j.IsDryRun() && j.Project().ID > 0 {
			app.DB().Model(j.Project()).Update("deploy_status", j.helmer.ReleaseStatus(j.Project().Name, j.Namespace().Name))
		}
	})

	var p models.Project
	if app.DB().Where("`name` = ? AND `namespace_id` = ?", j.project.Name, j.project.NamespaceId).First(&p).Error == gorm.ErrRecordNotFound {
		if !j.IsDryRun() {
			j.project.DeployStatus = uint8(types.Deploy_StatusDeploying)
			app.DB().Create(&j.project)
		}
		j.Messager().SendMsg("[Check]: 新建项目")
		j.isNew = true
	} else {
		j.project.ID = p.ID
		j.prevProject = &p
		if p.DeployStatus == uint8(types.Deploy_StatusDeploying) {
			return errors.New("有别人也在操作这个项目，等等哦~")
		}
		if !j.IsDryRun() {
			app.DB().Model(&p).Update("deploy_status", types.Deploy_StatusDeploying)
		}
	}
	if !j.IsDryRun() {
		j.PubSub().ToSelf(reloadProjectsMessage)
	}
	j.imagePullSecrets = j.Namespace().ImagePullSecretsArray()

	commit, err := plugins.GetGitServer().GetCommit(fmt.Sprintf("%d", j.project.GitProjectId), j.project.GitCommit)
	if err != nil {
		return err
	}
	j.commit = commit

	return nil
}

func defaultLoaders() []Loader {
	return []Loader{
		&ChartFileLoader{
			chartLoader: &defaultChartLoader{},
			fileOpener:  &defaultFileOpener{},
		},
		&VariableLoader{},
		&DynamicLoader{},
		&ExtraValuesLoader{},
		&MergeValuesLoader{},
		&ReleaseInstallerLoader{},
	}
}

func (j *Jober) LoadConfigs() error {
	wg := sync.WaitGroup{}
	wg.Add(1)
	ch := make(chan error, 1)
	go func() {
		defer recovery.HandlePanic("LoadConfigs")
		defer wg.Done()
		defer close(ch)
		err := func() error {
			j.Messager().SendMsg("[Check]: 加载项目文件")

			for _, defaultLoader := range j.loaders {
				if err := j.GetStoppedErrorIfHas(); err != nil {
					return err
				}
				if err := defaultLoader.Load(j); err != nil {
					return err
				}
			}

			return nil
		}()
		ch <- err
	}()

	var err error
	select {
	case err = <-ch:
	case <-j.stopCtx.Done():
		err = j.stopCtx.Err()
	}
	wg.Wait()

	return err
}

type Loader interface {
	Load(*Jober) error
}

type helmChartLoader interface {
	LoadDir(dir string) (*chart.Chart, error)
	LoadArchive(in io.Reader) (*chart.Chart, error)
}

type defaultFileOpener struct {
	f *os.File
}

func (d *defaultFileOpener) Open(name string) (io.ReadCloser, error) {
	open, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	d.f = open
	return open, err
}

func (d *defaultFileOpener) Close() error {
	return d.f.Close()
}

type defaultChartLoader struct{}

func (d *defaultChartLoader) LoadArchive(in io.Reader) (*chart.Chart, error) {
	return loader.LoadArchive(in)
}

func (d *defaultChartLoader) LoadDir(dir string) (*chart.Chart, error) {
	return loader.LoadDir(dir)
}

type fileOpener interface {
	Open(name string) (io.ReadCloser, error)
	Close() error
}

type ChartFileLoader struct {
	chartLoader helmChartLoader
	fileOpener  fileOpener
}

func (c *ChartFileLoader) Load(j *Jober) error {
	const loaderName = "[ChartFileLoader]: "
	j.Messager().SendMsg(loaderName + "加载 helm chart 文件")
	j.Percenter().To(20)

	// 下载 helm charts
	split := strings.Split(j.config.LocalChartPath, "|")
	var (
		files        []string
		tmpChartsDir string
		deleteDirFn  func()
		dir          string
	)
	// 如果是这个格式意味着是远程项目, 'uid|branch|path'
	j.Messager().SendMsg(fmt.Sprintf(loaderName+"下载 helm charts path: %s", j.config.LocalChartPath))

	var (
		pid    string = fmt.Sprintf("%d", j.input.GitProjectId)
		branch string = j.input.GitBranch
		path   string = j.config.LocalChartPath
	)
	if utils.IsRemoteChart(j.config) {
		pid = split[0]
		branch = split[1]
		path = split[2]
		files, _ = plugins.GetGitServer().GetDirectoryFilesWithBranch(pid, branch, path, true)
		if len(files) < 1 {
			return errors.New("charts 文件不存在")
		}
		var err error
		tmpChartsDir, deleteDirFn, err = utils.DownloadFiles(pid, branch, files)
		if err != nil {
			return err
		}

		dir = path
		j.Messager().SendMsg(fmt.Sprintf(loaderName+"识别为远程仓库 uid %v branch %s path %s", pid, branch, path))
	} else {
		var err error
		dir = j.config.LocalChartPath
		files, _ = plugins.GetGitServer().GetDirectoryFilesWithSha(fmt.Sprintf("%d", j.input.GitProjectId), j.input.GitCommit, j.config.LocalChartPath, true)
		tmpChartsDir, deleteDirFn, err = utils.DownloadFiles(j.input.GitProjectId, j.input.GitCommit, files)
		if err != nil {
			return err
		}
	}
	j.AddDestroyFunc(deleteDirFn)

	loadDir, _ := c.chartLoader.LoadDir(filepath.Join(tmpChartsDir, dir))
	if loadDir.Metadata.Dependencies != nil && action.CheckDependencies(loadDir, loadDir.Metadata.Dependencies) != nil {
		for _, dependency := range loadDir.Metadata.Dependencies {
			if strings.HasPrefix(dependency.Repository, "file://") {
				depFiles, _ := plugins.GetGitServer().GetDirectoryFilesWithBranch(pid, branch, filepath.Join(path, strings.TrimPrefix(dependency.Repository, "file://")), true)
				_, depDeleteFn, err := utils.DownloadFilesToDir(pid, branch, depFiles, tmpChartsDir)
				if err != nil {
					return err
				}
				j.AddDestroyFunc(depDeleteFn)
				j.Messager().SendMsg(fmt.Sprintf(loaderName+"下载本地依赖 %s", dependency.Name))
			}
		}
	}

	chartDir := filepath.Join(tmpChartsDir, dir)

	j.Percenter().To(30)
	j.Messager().SendMsg(loaderName + "打包 helm charts")
	chart, err := j.helmer.PackageChart(chartDir, chartDir)
	if err != nil {
		return err
	}
	archive, err := c.fileOpener.Open(chart)
	if err != nil {
		return err
	}
	defer archive.Close()

	if j.chart, err = c.chartLoader.LoadArchive(archive); err != nil {
		return err
	}

	return nil
}

type DynamicLoader struct{}

func (d *DynamicLoader) Load(j *Jober) error {
	const loaderName = "[DynamicLoader]: "

	j.Percenter().To(50)
	j.Messager().SendMsg(fmt.Sprintf(loaderName+"%v", "检查到用户传入的配置"))

	if j.input.Config == "" {
		j.Messager().SendMsg(fmt.Sprintf(loaderName+"%v", "未发现用户自定义配置"))
		return nil
	}

	dynamicConfigYaml, err := utils.ParseInputConfig(j.config, j.input.Config)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	j.dynamicConfigYaml = dynamicConfigYaml

	return nil
}

type ExtraValuesLoader struct{}

func (d *ExtraValuesLoader) Load(j *Jober) error {
	const loaderName = "[ExtraValuesLoader]: "

	j.Percenter().To(60)
	j.Messager().SendMsg(fmt.Sprintf(loaderName+"%v", "检查项目额外的配置"))

	if len(j.input.ExtraValues) <= 0 {
		j.Messager().SendMsg(fmt.Sprintf(loaderName+"%v", "未发现项目额外的配置"))
	}

	var validValuesMap = make(map[string]any)
	var useDefaultMap = make(map[string]bool)

	var configElementsMap = make(map[string]*mars.Element)
	for _, element := range j.config.Elements {
		configElementsMap[element.Path] = element
		defaultValue, e := d.typeValue(element, element.Default)
		if e != nil {
			return e
		}
		validValuesMap[element.Path] = defaultValue
		useDefaultMap[element.Path] = true
	}

	// validate
	for _, value := range j.input.ExtraValues {
		var fieldValid bool
		if element, ok := configElementsMap[value.Path]; ok {
			fieldValid = true
			useDefaultMap[value.Path] = false
			typeValue, err := d.typeValue(element, value.Value)
			if err != nil {
				return err
			}
			validValuesMap[value.Path] = typeValue
		}
		if !fieldValid {
			j.Messager().SendMsg(fmt.Sprintf("不允许自定义字段 %s", value.Path))
		}
	}

	j.extraValues = d.deepSetItems(validValuesMap)
	var ds []string
	for k, ok := range useDefaultMap {
		if ok {
			ds = append(ds, k)
		}
	}
	if len(ds) > 0 {
		j.Messager().SendMsg(fmt.Sprintf(loaderName+"已经为 '%s' 设置系统默认值", strings.Join(ds, ",")))
	}

	return nil
}

func (d *ExtraValuesLoader) typeValue(element *mars.Element, input string) (any, error) {
	switch element.Type {
	case mars.ElementType_ElementTypeSwitch:
		if input == "" {
			input = "false"
		}
		v, err := strconv.ParseBool(input)
		if err != nil {
			mlog.Error(err)
			return nil, fmt.Errorf("%s 字段类型不正确，应该为 bool，你传入的是 %s", element.Path, input)
		}
		return v, nil
	case mars.ElementType_ElementTypeInputNumber:
		if input == "" {
			input = "0"
		}
		v, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			mlog.Error(err)
			return nil, fmt.Errorf("%s 字段类型不正确，应该为整数，你传入的是 %s", element.Path, input)
		}
		return v, nil
	case mars.ElementType_ElementTypeRadio:
		fallthrough
	case mars.ElementType_ElementTypeSelect:
		var in bool
		for _, selectValue := range element.SelectValues {
			if input == selectValue {
				in = true
				break
			}
		}
		if !in {
			return nil, fmt.Errorf("%s 必须在 '%v' 里面, 你传的是 %s", element.Path, strings.Join(element.SelectValues, ","), input)
		}

		return input, nil
	default:
		return input, nil
	}
}

func (d *ExtraValuesLoader) deepSetItems(items map[string]any) []string {
	var evs []string
	for k, v := range items {
		ysk, err := utils.YamlDeepSetKey(k, v)
		if err != nil {
			mlog.Error(err)
			continue
		}
		evs = append(evs, string(ysk))
	}
	return evs
}

const (
	leftDelim  = "<"
	rightDelim = ">"

	VarImagePullSecrets = "ImagePullSecrets"
	VarBranch           = "Branch"
	VarCommit           = "Commit"
	VarPipeline         = "Pipeline"
	VarClusterIssuer    = "ClusterIssuer"
	VarHost             = "Host"
	VarTlsSecret        = "TlsSecret"
)

var tagRegex = regexp.MustCompile(leftDelim + `\s*(\.Branch|\.Commit|\.Pipeline)\s*` + rightDelim)

type VariableLoader struct {
	values vars
}

func (v *VariableLoader) Load(j *Jober) error {
	const loaderName = "[VariableLoader]: "
	j.Percenter().To(40)
	j.Messager().SendMsg(fmt.Sprintf(loaderName+"%v", "注入内置环境变量"))

	if j.config.ValuesYaml == "" {
		j.Messager().SendMsg(fmt.Sprintf(loaderName+"%v", "未发现可用的 values.yaml"))
		return nil
	}

	if v.values == nil {
		v.values = vars{}
	}

	//ImagePullSecrets
	parse, e := template.New("ImagePullSecrets").Parse(fmt.Sprintf("[{{- range .%s }}{name: {{ . }}}, {{- end }}]", VarImagePullSecrets))
	if e != nil {
		return e
	}

	renderResult := &bytes.Buffer{}
	if err := parse.Execute(renderResult, struct {
		ImagePullSecrets []string
	}{
		ImagePullSecrets: j.imagePullSecrets,
	}); err != nil {
		return err
	}

	v.values[VarImagePullSecrets] = renderResult.String()

	//Host1...Host10
	sub := utils.GetPreOccupiedLenByValuesYaml(j.config.ValuesYaml)
	mlog.Debug("getPreOccupiedLenByValuesYaml: ", sub)
	for i := 1; i <= 10; i++ {
		v.values[fmt.Sprintf("%s%d", VarHost, i)] = plugins.GetDomainManager().GetDomainByIndex(j.project.Name, j.Namespace().Name, i, sub)
		v.values[fmt.Sprintf("%s%d", VarTlsSecret, i)] = plugins.GetDomainManager().GetCertSecretName(j.project.Name, i)
	}

	//{{.Branch}}{{.Commit}}{{.Pipeline}}
	var (
		pipelineID     int64
		pipelineBranch string = j.project.GitBranch
		pipelineCommit string = j.Commit().GetShortID()
	)

	// 如果存在需要传变量的，则必须有流水线信息
	if pipeline, e := plugins.GetGitServer().GetCommitPipeline(fmt.Sprintf("%d", j.project.GitProjectId), j.project.GitCommit); e == nil {
		pipelineID = pipeline.GetID()
		pipelineBranch = pipeline.GetRef()

		j.Messager().SendMsg(fmt.Sprintf(loaderName+"镜像分支 %s 镜像commit %s 镜像 pipeline_id %d", pipelineBranch, pipelineCommit, pipelineID))
	} else {
		if tagRegex.MatchString(j.config.ValuesYaml) {
			return errors.New("无法获取 Pipeline 信息")
		}
	}

	v.values[VarBranch] = pipelineBranch
	v.values[VarCommit] = pipelineCommit
	v.values[VarPipeline] = pipelineID

	// ingress
	v.values[VarClusterIssuer] = plugins.GetDomainManager().GetClusterIssuer()

	tpl, err := template.New("values_yaml").Delims(leftDelim, rightDelim).Parse(j.config.ValuesYaml)
	if err != nil {
		return err
	}
	bf := bytes.Buffer{}
	tpl.Execute(&bf, v.values)
	j.valuesYaml = bf.String()
	j.vars = v.values

	return nil
}

type MergeValuesLoader struct{}

// Load
// imagePullSecrets 会自动注入到 imagePullSecrets 中
func (m *MergeValuesLoader) Load(j *Jober) error {
	const loaderName = "[MergeValuesLoader]: "
	j.Percenter().To(70)
	j.Messager().SendMsg(fmt.Sprintf(loaderName+"%v", "合并配置文件到 values.yaml"))

	// 自动注入 imagePullSecrets
	var imagePullSecrets = make([]map[string]any, len(j.imagePullSecrets))
	for i, s := range j.imagePullSecrets {
		imagePullSecrets[i] = map[string]any{"name": s}
	}
	var yamlImagePullSecrets []byte
	if len(imagePullSecrets) > 0 {
		yamlImagePullSecrets, _ = yaml.Marshal(map[string]any{
			"imagePullSecrets": imagePullSecrets,
		})
	}

	var opts []config.YAMLOption
	if j.valuesYaml != "" {
		opts = append(opts, config.Source(strings.NewReader(j.valuesYaml)))
	}
	if j.dynamicConfigYaml != "" {
		opts = append(opts, config.Source(strings.NewReader(j.dynamicConfigYaml)))
	}
	if len(yamlImagePullSecrets) != 0 {
		opts = append(opts, config.Source(bytes.NewReader(yamlImagePullSecrets)))
	}

	for _, value := range j.extraValues {
		opts = append(opts, config.Source(strings.NewReader(value)))
	}

	if len(opts) < 1 {
		return nil
	}

	// 5. 用用户传入的yaml配置去合并 `default_values`
	provider, err := config.NewYAML(opts...)
	if err != nil {
		mlog.Error(loaderName, err, j.valuesYaml, j.dynamicConfigYaml)

		return err
	}
	var mergedDefaultAndConfigYamlValues map[string]any
	if err := provider.Get("").Populate(&mergedDefaultAndConfigYamlValues); err != nil {
		mlog.Error(loaderName, mergedDefaultAndConfigYamlValues, err)
		return err
	}

	bf := &bytes.Buffer{}
	encoder := yaml.NewEncoder(bf)
	if err := encoder.Encode(&mergedDefaultAndConfigYamlValues); err != nil {
		return err
	}
	fileData := bf.String()
	//mlog.Debug("fileData", fileData)
	mergedFile, closer, err := utils.WriteConfigYamlToTmpFile([]byte(fileData))
	if err != nil {
		return err
	}
	j.AddDestroyFunc(func() { closer.Close() })
	j.valuesOptions.ValueFiles = append(j.valuesOptions.ValueFiles, mergedFile)

	return nil
}

type ReleaseInstallerLoader struct{}

func (r *ReleaseInstallerLoader) Load(j *Jober) error {
	const loaderName = "[ReleaseInstallerLoader]: "
	j.Messager().SendMsg(loaderName + "worker 已就绪, 准备安装")
	j.Percenter().To(80)
	j.installer = newReleaseInstaller(j.project.Name, j.Namespace().Name, j.chart, j.valuesOptions, j.input.Atomic, j.timeoutSeconds, j.dryRun)
	return nil
}

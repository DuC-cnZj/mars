package domain_manager

import (
	"encoding/base64"
	"errors"
	"strings"
	"sync"

	app "github.com/duc-cnzj/mars/internal/app/helper"
	"github.com/duc-cnzj/mars/internal/mlog"
	"github.com/duc-cnzj/mars/internal/plugins"
	"github.com/duc-cnzj/mars/internal/utils/tls"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
)

var _ plugins.DomainManager = (*SyncSecretDomainManager)(nil)

const SyncSecretSecretName = "mars-tls-sync-secret"

func init() {
	dr := &SyncSecretDomainManager{}
	plugins.RegisterPlugin(dr.Name(), dr)
}

type SyncSecretDomainManager struct {
	nsPrefix       string
	wildcardDomain string
	domainSuffix   string

	secretNamespace string
	secretName      string

	mu     sync.RWMutex
	secret *v1.Secret
}

func (d *SyncSecretDomainManager) SetSecret(s *v1.Secret) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.secret = s
}

func (d *SyncSecretDomainManager) GetSecret() *v1.Secret {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.secret
}

func (d *SyncSecretDomainManager) Name() string {
	return "sync_secret_domain_manager"
}

func (d *SyncSecretDomainManager) Initialize(args map[string]any) error {
	if p, ok := args["ns_prefix"]; ok {
		d.nsPrefix = p.(string)
	}

	if p, ok := args["secret_namespace"]; ok {
		d.secretNamespace = p.(string)
	}

	if p, ok := args["secret_name"]; ok {
		d.secretName = p.(string)
	}

	if wd, ok := args["wildcard_domain"]; ok {
		d.wildcardDomain = wd.(string)
		d.domainSuffix = strings.TrimLeft(d.wildcardDomain, "*.")
	}

	if d.secretNamespace == "" || d.secretName == "" || d.wildcardDomain == "" {
		return errors.New("secret_namespace, secret_name, wildcard_domain required")
	}

	secret, err := app.K8sClient().SecretLister.Secrets(d.secretNamespace).Get(d.secretName)
	if err != nil {
		return err
	}

	if secret.Type != v1.SecretTypeTLS {
		return errors.New("secret not verified")
	}

	var (
		tlsKey, _ = base64.StdEncoding.DecodeString(string(secret.Data["tls.key"]))
		tlsCrt, _ = base64.StdEncoding.DecodeString(string(secret.Data["tls.crt"]))
	)
	err = validateTelsWildcardDomain(tlsKey, tlsCrt, d.wildcardDomain)
	if err != nil {
		return err
	}
	d.SetSecret(secret)

	app.K8sClient().SecretInformer.AddEventHandler(d.eventHandler(d.handleSecretChange))

	mlog.Info("[Plugin]: " + d.Name() + " plugin Initialize...")
	return nil
}

func (d *SyncSecretDomainManager) eventHandler(updateFunc func(oldObj, newObj any)) cache.FilteringResourceEventHandler {
	return cache.FilteringResourceEventHandler{
		FilterFunc: func(obj any) bool {
			sec := obj.(*v1.Secret)
			return sec.Namespace == d.secretNamespace && sec.Name == d.secretName
		},
		Handler: cache.ResourceEventHandlerFuncs{
			UpdateFunc: updateFunc,
		},
	}
}

func (d *SyncSecretDomainManager) handleSecretChange(oldObj any, newObj any) {
	oldSec := oldObj.(*v1.Secret)
	newSec := newObj.(*v1.Secret)
	if newSec.ResourceVersion != oldSec.ResourceVersion {
		d.SetSecret(newSec)
		// 更新当前的所有 secret
		var (
			oldKey = string(oldSec.Data["tls.key"])
			newKey = string(newSec.Data["tls.key"])

			oldCrt = string(oldSec.Data["tls.crt"])
			newCrt = string(newSec.Data["tls.crt"])
		)
		if oldKey != newKey || oldCrt != newCrt {
			tls.UpdateCertTls(newSec.Name, newKey, newCrt)
		}
	}
}

func (d *SyncSecretDomainManager) Destroy() error {
	mlog.Info("[Plugin]: " + d.Name() + " plugin Destroy...")
	return nil
}

func (d *SyncSecretDomainManager) GetDomainByIndex(projectName, namespace string, index, preOccupiedLen int) string {
	return Subdomain{
		maxLen:       maxDomainLength - preOccupiedLen,
		projectName:  projectName,
		namespace:    namespace,
		index:        index,
		nsPrefix:     d.nsPrefix,
		domainSuffix: d.domainSuffix,
	}.SubStr()
}

func (d *SyncSecretDomainManager) GetDomain(projectName, namespace string, preOccupiedLen int) string {
	return Subdomain{
		maxLen:       maxDomainLength - preOccupiedLen,
		projectName:  projectName,
		namespace:    namespace,
		index:        -1,
		nsPrefix:     d.nsPrefix,
		domainSuffix: d.domainSuffix,
	}.SubStr()
}

func (d *SyncSecretDomainManager) GetCertSecretName(projectName string, index int) string {
	return SyncSecretSecretName
}

func (d *SyncSecretDomainManager) GetClusterIssuer() string {
	return ""
}

func (d *SyncSecretDomainManager) GetCerts() (name, key, crt string) {
	tlsKey, _ := base64.StdEncoding.DecodeString(string(d.secret.Data["tls.key"]))
	tlsCrt, _ := base64.StdEncoding.DecodeString(string(d.secret.Data["tls.crt"]))
	return SyncSecretSecretName, string(tlsKey), string(tlsCrt)
}

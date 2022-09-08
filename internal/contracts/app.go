package contracts

//go:generate mockgen -destination ../mock/mock_app.go -package mock github.com/duc-cnzj/mars/internal/contracts ApplicationInterface
//go:generate mockgen -destination ../mock/mock_tracer.go -package mock go.opentelemetry.io/otel/trace Tracer

import (
	"context"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/duc-cnzj/mars/internal/config"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/oauth2"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

type Callback func(ApplicationInterface)

type Server interface {
	Run(context.Context) error
	Shutdown(context.Context) error
}

type Bootstrapper interface {
	Bootstrap(ApplicationInterface) error
	Tags() []string
}

type OidcConfigItem struct {
	Provider           *oidc.Provider
	Config             oauth2.Config
	EndSessionEndpoint string
}
type OidcConfig map[string]OidcConfigItem

type ApplicationInterface interface {
	IsDebug() bool

	K8sClient() *K8sClient
	SetK8sClient(*K8sClient)

	Auth() AuthInterface
	SetAuth(AuthInterface)

	SetLocalUploader(Uploader)
	LocalUploader() Uploader

	SetUploader(Uploader)
	Uploader() Uploader

	Bootstrap() error
	Config() *config.Config

	DBManager() DBManager
	DB() *gorm.DB

	Oidc() OidcConfig
	SetOidc(OidcConfig)

	AddServer(Server)
	Run() context.Context
	Shutdown()

	Done() <-chan struct{}

	BeforeServerRunHooks(Callback)
	RegisterBeforeShutdownFunc(Callback)
	RegisterAfterShutdownFunc(Callback)

	EventDispatcher() DispatcherInterface
	SetEventDispatcher(DispatcherInterface)

	SetPlugins(map[string]PluginInterface)
	GetPlugins() map[string]PluginInterface
	GetPluginByName(string) PluginInterface

	Singleflight() *singleflight.Group

	SetCache(CacheInterface)
	Cache() CacheInterface

	CacheLock() Locker
	SetCacheLock(Locker)

	SetTracer(trace.Tracer)
	GetTracer() trace.Tracer

	SetCronManager(CronManager)
	CronManager() CronManager
}

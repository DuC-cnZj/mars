package app

import (
	"github.com/duc-cnzj/mars/internal/app/instance"
	"github.com/duc-cnzj/mars/internal/config"
	"github.com/duc-cnzj/mars/internal/contracts"
	"gorm.io/gorm"
	"k8s.io/client-go/kubernetes"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

func App() contracts.ApplicationInterface {
	return instance.App()
}

func Config() *config.Config {
	return App().Config()
}

func DB() *gorm.DB {
	return App().DBManager().DB()
}

func Event() contracts.DispatcherInterface {
	return App().EventDispatcher()
}

func K8sClient() *contracts.K8sClient {
	return App().K8sClient()
}

func K8sClientSet() kubernetes.Interface {
	return App().K8sClient().Client
}

func K8sMetrics() versioned.Interface {
	return App().K8sClient().MetricsClient
}

func Metrics() contracts.Metrics {
	return App().Metrics()
}

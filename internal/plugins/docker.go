package plugins

import (
	"sync"

	app "github.com/duc-cnzj/mars/internal/app/helper"
	"github.com/duc-cnzj/mars/internal/contracts"
)

var dockerOnce = sync.Once{}

type DockerPluginInterface interface {
	contracts.PluginInterface

	ImageNotExists(repo, tag string) bool
}

func GetDockerPlugin() DockerPluginInterface {
	pcfg := app.App().Config().DockerPlugin
	p := app.App().GetPluginByName(pcfg.Name)
	dockerOnce.Do(func() {
		if err := p.Initialize(pcfg.GetArgs()); err != nil {
			panic(err)
		}
		app.App().RegisterAfterShutdownFunc(func(applicationInterface contracts.ApplicationInterface) {
			p.Destroy()
		})
	})

	return p.(DockerPluginInterface)
}

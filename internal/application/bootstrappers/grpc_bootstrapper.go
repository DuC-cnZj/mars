package bootstrappers

import (
	"github.com/duc-cnzj/mars/v4/internal/application"
	"github.com/duc-cnzj/mars/v4/internal/server"
)

type GrpcBootstrapper struct{}

func (g *GrpcBootstrapper) Tags() []string {
	return []string{"api", "grpc"}
}

func (g *GrpcBootstrapper) Bootstrap(app application.App) error {
	app.AddServer(server.NewGrpcRunner(app.Config().GrpcPort, app))

	return nil
}

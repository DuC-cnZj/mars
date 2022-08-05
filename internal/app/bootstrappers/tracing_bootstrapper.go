package bootstrappers

import (
	"context"
	"time"

	"github.com/duc-cnzj/mars/internal/contracts"
	"github.com/duc-cnzj/mars/internal/mlog"
	"github.com/duc-cnzj/mars/version"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

const serviceName = "mars"

type TracingBootstrapper struct{}

func (t *TracingBootstrapper) Bootstrap(app contracts.ApplicationInterface) error {
	cfg := app.Config()
	if cfg.JaegerAgentHostPort != "" {
		jaeexp, err := newJaegerExporter(cfg.JaegerAgentHostPort, cfg.JaegerUser, cfg.JaegerPassword)
		if err != nil {
			return err
		}
		tp := trace.NewTracerProvider(
			trace.WithBatcher(jaeexp),
			trace.WithResource(newResource()),
		)
		otel.SetTracerProvider(tp)
		app.RegisterAfterShutdownFunc(func(app contracts.ApplicationInterface) {
			mlog.Info("shutdown tracer")
			timeout, cancelFunc := context.WithTimeout(context.TODO(), 3*time.Second)
			defer cancelFunc()
			if err := tp.Shutdown(timeout); err != nil {
				mlog.Error(err)
			}
		})
	}
	tracer := otel.Tracer("mars")
	app.SetTracer(tracer)

	return nil
}

func newResource() *resource.Resource {
	v := version.GetVersion()
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(serviceName),
		semconv.ServiceVersionKey.String(v.String()),
		attribute.String("system.build_date", v.BuildDate),
		attribute.String("system.git_commit", v.GitCommit),
		attribute.String("system.git_branch", v.GitBranch),
		attribute.String("system.go_version", v.GoVersion),
		attribute.String("system.platform", v.Platform),
	)
}

func newJaegerExporter(hostAndPort, user, pwd string) (trace.SpanExporter, error) {
	return jaeger.New(
		jaeger.WithCollectorEndpoint(
			jaeger.WithUsername(user),
			jaeger.WithPassword(pwd),
			jaeger.WithEndpoint(hostAndPort),
		),
	)
}

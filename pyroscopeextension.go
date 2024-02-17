package pyroscopeextension

import (
	"context"
	"errors"
	"sync/atomic"

	"go.opentelemetry.io/collector/component"
        "github.com/grafana/pyroscope-go"
)

var running = &atomic.Bool{}

type pyroscopeExtension struct {
	config            Config
	telemetrySettings component.TelemetrySettings
}

func (p *pyroscopeExtension) Start(_ context.Context, _ component.Host) error {
	if !running.CompareAndSwap(false, true) {
		return errors.New("only a single pyroscope extension instance can be running per process")
	}

	var startErr error
	defer func() {
		if startErr != nil {
			running.Store(false)
		}
	}()

        pyroscope.Start(pyroscope.Config{
            ApplicationName: p.config.ApplicationName,
            ServerAddress: p.config.ServerAddress,
            BasicAuthUser: p.config.User,
            BasicAuthPassword: p.config.Password,
        })

	return startErr
}

func (p *pyroscopeExtension) Shutdown(context.Context) error {
	defer running.Store(false)
	return nil
}

func newServer(config Config, params component.TelemetrySettings) *pyroscopeExtension {
	return &pyroscopeExtension{
		config:            config,
		telemetrySettings: params,
	}
}

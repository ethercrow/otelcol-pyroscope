package pyroscopeextension

import (
	"go.opentelemetry.io/collector/component"
)

type Config struct {
    ApplicationName string `mapstructure:"application_name"`
    ServerAddress string `mapstructure:"server_address"`
    User string `mapstructure:"user"`
    Password string `mapstructure:"password"`
}

var _ component.Config = (*Config)(nil)

func (cfg *Config) Validate() error {
	return nil
}

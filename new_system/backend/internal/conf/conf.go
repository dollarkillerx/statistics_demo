package conf

import "github.com/dollarkillerx/common/pkg/config"

type Config struct {
	ServiceConfiguration       config.ServiceConfiguration
	PostgresConfiguration      config.PostgresConfiguration
	RedisConfiguration         config.RedisConfiguration
	LoggerConfiguration        config.LoggerConfig
	OpenTelemetryConfiguration config.OpenTelemetryConfig
}

package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// MySQLConfiguration  configuration for MySQL database connection
type MySQLConfiguration struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	LogMode  MySQLLogMode
	Charset  string
}

// PostgresConfiguration  configuration for Postgres database connection
type PostgresConfiguration struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  bool
	TimeZone string
	LogMode  MySQLLogMode
}

// ServiceConfiguration  configuration for service
type ServiceConfiguration struct {
	Port  string
	Debug bool
}

// RedisConfiguration ...
type RedisConfiguration struct {
	Addr     string
	Db       int
	Password string
}

// LoggerConfig configuration for logger
type LoggerConfig struct {
	Filename string
	MaxSize  int // MB
}

// MySQLLogMode ...
type MySQLLogMode string

// Console 使用 gorm 的 logger，打印漂亮的sql到控制台
// SlowQuery 使用自定义 logger.Logger,记录慢查询sql到日志
// None 关闭 logs 功能
const (
	Console   MySQLLogMode = "console"
	SlowQuery MySQLLogMode = "slow_query"
	None      MySQLLogMode = "none"
)

type OpenTelemetryConfig struct {
	Traces OpenTelemetryTracesConfig
	Logs   OpenTelemetryLogsConfig
}

type OpenTelemetryTracesConfig struct {
	HTTPEndpoint  string
	Path          string
	Authorization string
	ServerName    string
}

type OpenTelemetryLogsConfig struct {
	HTTPEndpoint string
	User         string
	Password     string

	File string // log file path
}

type MetricsConfig struct {
	HTTPEndpoint      string
	User              string
	Password          string
	MaxSamplesPerSend string
}

// InitConfiguration initializes the configuration by first attempting to read from environment variables,
// and then falling back to configuration files if the environment variables are not set.
func InitConfiguration(configName string, configPaths []string, config interface{}) error {
	vp := viper.New()
	vp.SetConfigName(configName)
	vp.AutomaticEnv()

	// Read in configuration files
	for _, configPath := range configPaths {
		vp.AddConfigPath(configPath)
	}

	if err := vp.ReadInConfig(); err != nil {
		// Only return an error if no config file was found, but continue if it's not found
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return errors.WithStack(err)
		}
	}

	// Unmarshal the config into the provided config struct
	if err := vp.Unmarshal(config); err != nil {
		return errors.WithStack(err)
	}

	// Bind all keys to environment variables to ensure environment variables take precedence
	for _, key := range vp.AllKeys() {
		if err := vp.BindEnv(key); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

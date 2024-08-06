package open_telemetry

import (
	"github.com/dollarkillerx/common/pkg/config"
	"github.com/dollarkillerx/common/pkg/logs"
	"github.com/rs/zerolog/log"

	"testing"
	"time"
)

func TestLog(t *testing.T) {
	cf := config.LoggerConfig{
		Filename: "my.log",
		MaxSize:  20,
	}

	logs.InitLog(cf)

	go func() {
		for {
			log.Info().Msg("hello world")
			time.Sleep(time.Second)
		}
	}()

	InitLog(config.OpenTelemetryLogsConfig{
		HTTPEndpoint: "http://127.0.0.1:5081/api/default/log1/_json",
		User:         "google@google.com",
		Password:     "tuqrqMH8WZsTn2Km",
		File:         "my.log",
	})
}

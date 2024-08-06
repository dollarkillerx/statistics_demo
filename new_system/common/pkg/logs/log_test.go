package logs

import (
	"github.com/dollarkillerx/common/pkg/config"
	"github.com/rs/zerolog/log"
	"testing"
)

func TestLogs(t *testing.T) {
	conf := config.LoggerConfig{
		Filename: "my.log",
		MaxSize:  20,
	}

	InitLog(conf)
	log.Info().Str("Code", "500").Msg("hello world")
}

package internal

import (
	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"os"
	"time"
)

func InitLog() {
	// Configure the lumberjack logger
	rotatingLogger := &lumberjack.Logger{
		Filename:   "tv.log", // Log file name
		MaxSize:    25,       // Max size in MB before rotating
		MaxBackups: 1,        // Max number of old logs files to keep
		MaxAge:     28,       // Max number of days to retain old logs files
		Compress:   true,     // Compress/zip old logs files
	}

	// Configure zerolog to write to the lumberjack logger
	log.Logger = zerolog.New(rotatingLogger).With().Caller().Logger()

	// Optionally, configure zerolog to write to both console and file
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	multi := zerolog.MultiLevelWriter(consoleWriter, rotatingLogger)
	log.Logger = zerolog.New(multi).With().Caller().Logger()
	log.Info().Msg("Logger initialized")
}

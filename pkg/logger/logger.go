package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {

	// Set the global logger to use JSON format.
	zerolog.TimeFieldFormat = time.RFC3339

	// Get the log file name from the environment variable.
	logFile := os.Getenv("CM_BEETLE_LOG_FILE")
	if logFile == "" {
		// If the environment variable is not set, use the default file name.
		logFile = "cm-beetle.log"
	}

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open log file")
	}

	// Multi-writer setup: logs to both file and console
	multi := zerolog.MultiLevelWriter(
		file,
		zerolog.ConsoleWriter{Out: os.Stdout},
	)

	// Set the log level
	logLevel := os.Getenv("CM_BEETLE_LOG_LEVEL")
	switch logLevel {
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel) // default level
	}

	// Check the execution environment from the environment variable
	env := os.Getenv("CM_BEETLE_ENV")

	// Configure the log output
	if env == "production" {
		// Apply multi-writer to the global logger
		// Set the global logger to use JSON format.
		log.Logger = zerolog.New(file).With().Timestamp().Caller().Logger()
	} else {
		// Apply file to the global logger
		log.Logger = zerolog.New(multi).With().Timestamp().Caller().Logger()
	}

	// Log a message
	log.Info().Msgf("Logger initialized: level=%s, file=%s, environment=%s", logLevel, logFile, env)
	if env == "production" {
		log.Info().Msg("Single-write setup: logs to file only")
	} else {
		log.Info().Msg("Multi-writes setup: logs to both file and console")
	}
}

package logger

import (
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {

	// Set the global logger to use JSON format.
	zerolog.TimeFieldFormat = time.RFC3339

	// Get log-related environment variables
	// Default: cm-beetle.log
	logFilePath := os.Getenv("CM_BEETLE_LOG_PATH")
	if logFilePath == "" {
		logFilePath = "cm-beetle.log"
	}

	// Default: 10 MB
	maxSize, err := strconv.Atoi(os.Getenv("CM_BEETLE_LOG_MAX_SIZE"))
	if err != nil {
		log.Fatal().Msg("Invalid CM_BEETLE_LOG_MAX_SIZE value")
	}

	// Default: 3 backups
	maxBackups, err := strconv.Atoi(os.Getenv("CM_BEETLE_LOG_MAX_BACKUPS"))
	if err != nil {
		log.Fatal().Msg("Invalid CM_BEETLE_LOG_MAX_BACKUPS value")
	}

	// Default: 30 days
	maxAge, err := strconv.Atoi(os.Getenv("CM_BEETLE_LOG_MAX_AGE"))
	if err != nil {
		log.Fatal().Msg("Invalid CM_BEETLE_LOG_MAX_AGE value")
	}

	// Default: false
	compress, err := strconv.ParseBool(os.Getenv("CM_BEETLE_LOG_COMPRESS"))
	if err != nil {
		log.Fatal().Msg("Invalid CM_BEETLE_LOG_COMPRESS value")
	}

	// lumberjack log file setup (lumberjack is used for log rotation)
	logFile := &lumberjack.Logger{
		Filename:   logFilePath, // log file path
		MaxSize:    maxSize,     // max size in megabytes before log is rotated
		MaxBackups: maxBackups,  // max number of old log files to keep
		MaxAge:     maxAge,      // max number of days to retain log files
		Compress:   compress,    // compress/logrotate log files
	}

	// Multi-writer setup: logs to both file and console
	multi := zerolog.MultiLevelWriter(
		logFile,
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
		log.Fatal().Msgf("Invalid CM_BEETLE_LOG_LEVEL value: %s", logLevel)
	}

	// Check the execution environment from the environment variable
	env := os.Getenv("CM_BEETLE_APP_ENV")

	// Configure the log output
	if env == "production" {
		// Apply multi-writer to the global logger
		log.Logger = zerolog.New(logFile).With().Timestamp().Caller().Logger()
	} else {
		// Apply file to the global logger
		log.Logger = zerolog.New(multi).With().Timestamp().Caller().Logger()
	}

	// Log a message
	log.Info().
		Str("logLevel", logLevel).
		Str("logFilePath", logFilePath).
		Str("env", env).
		Int("maxSize", maxSize).
		Int("maxBackups", maxBackups).
		Int("maxAge", maxAge).
		Bool("compress", compress).
		Msg("Logger initialized")

	if env == "production" {
		log.Info().
			Str("logFilePath", logFilePath).
			Msg("Single-write setup (logs to file only)")
	} else {
		log.Info().
			Str("logFilePath", logFilePath).
			Str("ConsoleWriter", "os.Stdout").
			Msg("Multi-writes setup (logs to both file and console)")
	}
}

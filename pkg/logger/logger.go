package logger

import (
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	sharedLogFile *lumberjack.Logger
)

func init() {

	// Set the global logger to use JSON format.
	zerolog.TimeFieldFormat = time.RFC3339

	// Get log file configuration from environment variables
	logFilePath, maxSize, maxBackups, maxAge, compress := getLogFileConfig()

	// Initialize a shared log file with lumberjack to manage log rotation
	sharedLogFile = &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}

	// Set the log level
	logLevel := os.Getenv("CM_BEETLE_LOG_LEVEL")
	var level zerolog.Level
	switch logLevel {
	case "trace":
		level = zerolog.TraceLevel
	case "debug":
		level = zerolog.DebugLevel
	case "info":
		level = zerolog.InfoLevel
	case "warn":
		level = zerolog.WarnLevel
	case "error":
		level = zerolog.ErrorLevel
	case "fatal":
		level = zerolog.FatalLevel
	case "panic":
		level = zerolog.PanicLevel
	default:
		log.Warn().Msgf("Invalid CM_BEETLE_LOG_LEVEL value: %s. Using default value: info", logLevel)
		level = zerolog.InfoLevel
	}

	logger := NewLogger(level)

	// Set global logger
	log.Logger = *logger

	// Check the execution environment from the environment variable
	env := os.Getenv("CM_BEETLE_APP_ENV")

	// Log a message
	log.Info().
		Str("logLevel", level.String()).
		Str("logFilePath", sharedLogFile.Filename).
		Str("env", env).
		Int("maxSize", maxSize).
		Int("maxBackups", maxBackups).
		Int("maxAge", maxAge).
		Bool("compress", compress).
		Msg("Global logger initialized")

	if env == "production" {
		log.Info().
			Str("logFilePath", sharedLogFile.Filename).
			Msg("Single-write setup (logs to file only)")
	} else {
		log.Info().
			Str("logFilePath", sharedLogFile.Filename).
			Str("ConsoleWriter", "os.Stdout").
			Msg("Multi-writes setup (logs to both file and console)")
	}
}

// Create a new logger
func NewLogger(level zerolog.Level) *zerolog.Logger {

	// Multi-writer setup: logs to both file and console
	multi := zerolog.MultiLevelWriter(
		sharedLogFile,
		zerolog.ConsoleWriter{Out: os.Stdout},
	)

	var logger zerolog.Logger

	// Check the execution environment from the environment variable
	env := os.Getenv("CM_BEETLE_APP_ENV")

	// Configure the log output
	if env == "production" {
		// Apply multi-writer to the global logger
		logger = zerolog.New(sharedLogFile).Level(level).With().Timestamp().Caller().Logger()
	} else {
		// Apply file to the global logger
		logger = zerolog.New(multi).Level(level).With().Timestamp().Caller().Logger()
	}

	// Log a message
	logger.Info().
		Str("logLevel", level.String()).
		Msg("New logger created")

	if env == "production" {
		logger.Info().
			Str("logFilePath", sharedLogFile.Filename).
			Msg("Single-write setup (logs to file only)")
	} else {
		logger.Info().
			Str("logFilePath", sharedLogFile.Filename).
			Str("ConsoleWriter", "os.Stdout").
			Msg("Multi-writes setup (logs to both file and console)")
	}

	return &logger
}

// Get log file configuration from environment variables
func getLogFileConfig() (string, int, int, int, bool) {

	// Default: cm-beetle.log
	logFilePath := os.Getenv("CM_BEETLE_LOG_PATH")
	if logFilePath == "" {
		log.Info().Msg("CM_BEETLE_LOG_PATH is not set. Using default value: cm-beetle.log")
		logFilePath = "cm-beetle.log"
	}

	// Default: 10 MB
	maxSize, err := strconv.Atoi(os.Getenv("CM_BEETLE_LOG_MAX_SIZE"))
	if err != nil {
		log.Warn().Msgf("Invalid CM_BEETLE_LOG_MAX_SIZE value: %s. Using default value: 10 MB", os.Getenv("CM_BEETLE_LOG_MAX_SIZE"))
		maxSize = 10
	}

	// Default: 3 backups
	maxBackups, err := strconv.Atoi(os.Getenv("CM_BEETLE_LOG_MAX_BACKUPS"))
	if err != nil {
		log.Warn().Msgf("Invalid CM_BEETLE_LOG_MAX_BACKUPS value: %s. Using default value: 3 backups", os.Getenv("CM_BEETLE_LOG_MAX_BACKUPS"))
		maxBackups = 3
	}

	// Default: 30 days
	maxAge, err := strconv.Atoi(os.Getenv("CM_BEETLE_LOG_MAX_AGE"))
	if err != nil {
		log.Warn().Msgf("Invalid CM_BEETLE_LOG_MAX_AGE value: %s. Using default value: 30 days", os.Getenv("CM_BEETLE_LOG_MAX_AGE"))
		maxAge = 30
	}

	// Default: false
	compress, err := strconv.ParseBool(os.Getenv("CM_BEETLE_LOG_COMPRESS"))
	if err != nil {
		log.Warn().Msgf("Invalid CM_BEETLE_LOG_COMPRESS value: %s. Using default value: false", os.Getenv("CM_BEETLE_LOG_COMPRESS"))
		compress = false
	}

	return logFilePath, maxSize, maxBackups, maxAge, compress
}

package logger

import (
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Define context keys
type contextKey string

const (
	TraceIdKey contextKey = "trace_id"
	SpanIdKey  contextKey = "span_id"
)

// Define TracingHook struct
type TracingHook struct{}

// Run method: Executed when a log event occurs
func (h TracingHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	ctx := e.GetCtx()
	traceID := ctx.Value(TraceIdKey)
	spanID := ctx.Value(SpanIdKey)

	if traceID != nil {
		e.Str(string(TraceIdKey), traceID.(string))
	}
	if spanID != nil {
		e.Str(string(SpanIdKey), spanID.(string))
	}
}

var (
	sharedLogFile *lumberjack.Logger
	once          sync.Once
	traceLogger   zerolog.Logger
)

type Config struct {
	LogLevel          string
	LogWriter         string
	LogFilePath       string
	MaxSize           int
	MaxBackups        int
	MaxAge            int
	Compress          bool
	LogMaskingEnabled *bool // if nil, defaults to true (masking enabled)
}

// BoolPtr returns a pointer to the given bool value.
func BoolPtr(b bool) *bool {
	return &b
}

func init() {

	// For consistent log format across different running environments (e.g., local, Docker, Kubernetes)
	// Set the caller field to the relative path from the project root
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(b), "../../") // predict the project root directory from the current file having init() function

	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		// relative path from the project root
		relPath, err := filepath.Rel(projectRoot, file)
		if err != nil {
			return filepath.Base(file) + ":" + strconv.Itoa(line) // return the original file path with line number if the relative path cannot be resolved
		}
		return relPath + ":" + strconv.Itoa(line)
	}
}

// NewLogger initializes a new logger with default values if not provided
func NewLogger(config Config) *zerolog.Logger {
	// Apply default values if not provided
	if config.LogLevel == "" {
		config.LogLevel = "debug"
	}
	if config.LogWriter == "" {
		config.LogWriter = "console"
	}
	if config.LogFilePath == "" {
		config.LogFilePath = "./log/app.log"
	}
	if config.MaxSize == 0 {
		config.MaxSize = 1000 // in MB
	}
	if config.MaxBackups == 0 {
		config.MaxBackups = 3
	}
	if config.MaxAge == 0 {
		config.MaxAge = 30 // in days
	}

	logMaskingEnabled := true
	if config.LogMaskingEnabled != nil {
		logMaskingEnabled = *config.LogMaskingEnabled
	}

	// Initialize shared log file for log rotation once
	once.Do(func() {
		sharedLogFile = &lumberjack.Logger{
			Filename:   config.LogFilePath,
			MaxSize:    config.MaxSize,
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge,
			Compress:   config.Compress,
		}

		// Ensure the log file directory exists before creating the log file
		dir := filepath.Dir(config.LogFilePath)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			// Create the directory if it does not exist
			err = os.MkdirAll(dir, 0755) // Set permissions as needed
			if err != nil {
				log.Fatal().Msgf("Failed to create log directory: %v", err)
			}
		}

		// Ensure the log file exists before changing its permissions
		if _, err := os.Stat(config.LogFilePath); os.IsNotExist(err) {
			// Create the log file if it does not exist
			file, err := os.Create(config.LogFilePath)
			if err != nil {
				log.Fatal().Msgf("Failed to create log file: %v", err)
			}
			file.Close()
		}

		// Change file permissions to -rw-r--r--
		if err := os.Chmod(config.LogFilePath, 0644); err != nil {
			log.Fatal().Msgf("Failed to change file permissions: %v", err)
		}

		// traceLogger = zerolog.New(sharedLogFile).Level(zerolog.TraceLevel).With().Timestamp().Caller().Logger()
		var traceWriter io.Writer = sharedLogFile
		if logMaskingEnabled {
			traceWriter = NewMaskingWriter(sharedLogFile)
		}
		traceLogger = zerolog.New(traceWriter).Level(zerolog.TraceLevel).With().Timestamp().Logger()
		traceLogger.Hook(TracingHook{})
	})

	level := getLogLevel(config.LogLevel)
	logger := configureWriter(config.LogWriter, level, logMaskingEnabled)

	// Add tracing hook to the logger
	logger.Hook(TracingHook{})

	// Log a message to confirm logger setup
	logger.Info().
		Str("logLevel", level.String()).
		Bool("logMaskingEnabled", logMaskingEnabled).
		Msg("New logger created")

	return logger
}

// GetTraceLogger returns the trace logger
func GetTraceLogger() *zerolog.Logger {
	return &traceLogger
}

// getLogLevel returns the zerolog.Level based on the string level
func getLogLevel(logLevel string) zerolog.Level {
	switch logLevel {
	case "trace":
		return zerolog.TraceLevel
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	default:
		log.Warn().Msgf("Invalid log level: %s. Using default value: info", logLevel)
		return zerolog.InfoLevel
	}
}

// configureWriter sets up the logger based on the writer type
func configureWriter(logWriter string, level zerolog.Level, logMaskingEnabled bool) *zerolog.Logger {
	var logger zerolog.Logger
	stdoutWriter := zerolog.ConsoleWriter{Out: os.Stdout}

	// Handle case where sharedLogFile is not yet initialized
	if sharedLogFile == nil {
		var finalWriter io.Writer = stdoutWriter
		if logMaskingEnabled {
			finalWriter = NewMaskingWriter(stdoutWriter)
		}
		logger = zerolog.New(finalWriter).Level(level).With().Timestamp().Caller().Logger()
		logSetupInfo(logger, "stdout", logMaskingEnabled)
		return &logger
	}

	multi := zerolog.MultiLevelWriter(sharedLogFile, stdoutWriter)

	var targetWriter io.Writer
	switch logWriter {
	case "both":
		targetWriter = multi
	case "file":
		targetWriter = sharedLogFile
	case "stdout":
		targetWriter = stdoutWriter
	default:
		log.Warn().Msgf("Invalid log writer: %s. Using default value: both", logWriter)
		targetWriter = multi
	}

	var finalWriter io.Writer = targetWriter
	if logMaskingEnabled {
		finalWriter = NewMaskingWriter(targetWriter)
	}
	logger = zerolog.New(finalWriter).Level(level).With().Timestamp().Caller().Logger()

	logSetupInfo(logger, logWriter, logMaskingEnabled)
	return &logger
}

// logSetupInfo logs the logger setup details
func logSetupInfo(logger zerolog.Logger, logWriter string, logMaskingEnabled bool) {
	if logWriter == "file" {
		logger.Info().
			Str("logFilePath", sharedLogFile.Filename).
			Bool("logMaskingEnabled", logMaskingEnabled).
			Msg("Single-write setup (logs to file only)")
	} else if logWriter == "stdout" {
		logger.Info().
			Str("ConsoleWriter", "os.Stdout").
			Bool("logMaskingEnabled", logMaskingEnabled).
			Msg("Single-write setup (logs to console only)")
	} else {
		logger.Info().
			Str("logFilePath", sharedLogFile.Filename).
			Str("ConsoleWriter", "os.Stdout").
			Bool("logMaskingEnabled", logMaskingEnabled).
			Msg("Multi-writes setup (logs to both file and console)")
	}
}

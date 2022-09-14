package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gohub/pkg/app"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

// Logger global logger object
var Logger *zap.Logger

// InitLogger init logger
func InitLogger(filename string, maxSize, maxBackup, maxAge int,
	compress bool, logType string, level string) {
	// Get logger write media
	writeSyncer := getLogWriter(filename, maxSize, maxBackup, maxAge, compress, logType)

	// Setting log level, detail reference config/log.go file
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		fmt.Println("Logger init error, log level setting incorrect." +
			" please edit /config/log.go file log.level option")
	}

	// Init core
	core := zapcore.NewCore(getEncoder(), writeSyncer, logLevel)

	// Init logger
	Logger = zap.New(core,
		zap.AddCaller(),                   // Call file line, inner use runtime.Caller
		zap.AddCallerSkip(1),              // Remove one layer, runtime.Caller(1)
		zap.AddStacktrace(zap.ErrorLevel), // Error cursor show stacktrace
	)

	// Replace logger for global logger
	// zap.L().Fatal() call will be call our custom Logger
	zap.ReplaceGlobals(Logger)
}

func getEncoder() zapcore.Encoder {
	// Log rule
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller", // Code call, e.g: paginator/paginator.go:300
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,      // Each line log append "\n"
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // Log level, upper case e.g: ERROR, INFO
		EncodeTime:     customTimeEncoder,              // Time format 2006-01-02 15:04:05
		EncodeDuration: zapcore.SecondsDurationEncoder, // Execute time until is second
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Caller short format e.g: types/converter.go:17
	}

	if app.IsLocal() {
		//  Terminal keywords highlight output
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		// Local setting console encoder support stacktrace new line
		return zapcore.NewConsoleEncoder(encoderConfig)
	}

	// Production environment json encoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// customTimeEncoder time format
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int,
	compress bool, logType string) zapcore.WriteSyncer {
	// For daily record log file
	if logType == "daily" {
		logName := time.Now().Format("2006-01-02.log")
		filename = strings.ReplaceAll(filename, "logs.log", logName)
	}

	// Scroll log, detail reference config/log.go
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
		Compress:   compress,
	}

	// Output media
	if app.IsLocal() {
		// Local develop
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),
			zapcore.AddSync(lumberJackLogger))
	} else {
		// Production log
		return zapcore.AddSync(lumberJackLogger)
	}
}

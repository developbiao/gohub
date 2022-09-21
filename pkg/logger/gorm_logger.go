package logger

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"gohub/pkg/helpers"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// GormLogger implement gormlogger.Interface
type GormLogger struct {
	ZapLogger     *zap.Logger
	SlowThreshold time.Duration
}

func NewGormLogger() GormLogger {
	return GormLogger{
		ZapLogger:     Logger,
		SlowThreshold: 200 * time.Millisecond, // Slow query threshold time is 1/1000 second
	}
}

// LogMode implement gormlogger.Interface  LogMode method
func (l GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return GormLogger{
		ZapLogger:     l.ZapLogger,
		SlowThreshold: l.SlowThreshold,
	}
}

// Info implementation gormlogger.Interface Info
func (l GormLogger) Info(ctx context.Context, str string, args ...interface{}) {
	l.logger().Sugar().Debugf(str, args...)
}

// Warn implementation Warn
func (l GormLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	l.logger().Sugar().Warnf(str, args...)
}

// Error implementation Error
func (l GormLogger) Error(ctx context.Context, str string, args ...interface{}) {
	l.logger().Sugar().Errorf(str, args...)
}

// Trace implementation Trace
func (l GormLogger) Trace(ctx context.Context, begin time.Time,
	fc func() (string, int64), err error) {
	// Get elapsed time
	elapsed := time.Since(begin)
	// Get SQL request and return count
	sql, rows := fc()

	// Common fields
	logFields := []zap.Field{
		zap.String("sql", sql),
		zap.String("time", helpers.MicrosecondsStr(elapsed)),
		zap.Int64("rows", rows),
	}

	// Gorm error
	if err != nil {
		// Record not found using warning level
		if errors.Is(err, gorm.ErrRecordNotFound) {
			l.logger().Warn("Database ErrRecordNotFound")
		} else {
			// Other error level
			logFields = append(logFields, zap.Error(err))
			l.logger().Error("Database error", logFields...)
		}
	}

	// Slow query log
	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		l.logger().Warn("Database Slow log", logFields...)
	}

	// Record all sql
	l.logger().Debug("Database Query", logFields...)
}

// logger inner helper method, keep Zap inside message Caller is correct
// e.g paginator/paginator.go:148
func (l GormLogger) logger() *zap.Logger {
	var (
		gormPackage    = filepath.Join("gorm.io", "gorm")
		zapgormPackage = filepath.Join("moul.io", "zapgorm2")
	)
	clone := l.ZapLogger.WithOptions(zap.AddCallerSkip(-2))

	// Reduce one layout and logger init add zap.AddCallerSkip(1)
	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		case strings.Contains(file, gormPackage):
		case strings.Contains(file, zapgormPackage):
		default:
			return clone.WithOptions(zap.AddCallerSkip(i))
		}
	}
	return l.ZapLogger
}

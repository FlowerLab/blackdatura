// +build bd_all bd_gorm gorm

package blackdatura

import (
	"context"
	"errors"
	"strconv"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	gl "gorm.io/gorm/logger"
)

type GormLogger struct {
	log                       *zap.Logger
	Level                     zapcore.Level
	SlowThreshold             time.Duration
	SkipCallerLookup          bool
	IgnoreRecordNotFoundError bool
}

func NewGormLogger() *GormLogger {
	return &GormLogger{
		log:                       With("gorm log"),
		Level:                     zap.WarnLevel,
		SlowThreshold:             100 * time.Millisecond,
		SkipCallerLookup:          false,
		IgnoreRecordNotFoundError: false,
	}
}

func (l *GormLogger) LogMode(level gl.LogLevel) gl.Interface {
	newLog := &GormLogger{
		log:                       l.log,
		Level:                     zap.DPanicLevel,
		SlowThreshold:             l.SlowThreshold,
		SkipCallerLookup:          l.SkipCallerLookup,
		IgnoreRecordNotFoundError: l.IgnoreRecordNotFoundError,
	}
	switch level {
	case gl.Error:
		newLog.Level = zap.ErrorLevel
	case gl.Warn:
		newLog.Level = zap.WarnLevel
	case gl.Info:
		newLog.Level = zap.InfoLevel
	}
	return newLog
}

func (l *GormLogger) Info(_ context.Context, str string, args ...interface{}) {
	if l.Level >= zap.InfoLevel {
		l.log.Info(str, arg2ZapField(args)...)
	}
}

func (l *GormLogger) Warn(_ context.Context, str string, args ...interface{}) {
	if l.Level >= zap.WarnLevel {
		l.log.Warn(str, arg2ZapField(args)...)
	}
}

func (l *GormLogger) Error(_ context.Context, str string, args ...interface{}) {
	if l.Level >= zap.ErrorLevel {
		l.log.Error(str, arg2ZapField(args)...)
	}
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()
	elapsed := time.Since(begin)

	switch {
	case err != nil && l.Level >= zap.ErrorLevel &&
		(!l.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		l.log.Error("trace",
			zap.Error(err),
			zap.Duration("elapsed", elapsed),
			zap.Int64("rows", rows),
			zap.String("sql", sql))
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.Level >= zap.WarnLevel:
		l.log.Warn("trace",
			zap.Duration("elapsed", elapsed),
			zap.Int64("rows", rows),
			zap.String("sql", sql))
	case l.Level == zap.InfoLevel:
		l.log.Info("trace",
			zap.Duration("elapsed", elapsed),
			zap.Int64("rows", rows),
			zap.String("sql", sql))
	}
}

func arg2ZapField(args []interface{}) []zap.Field {
	arr := make([]zap.Field, len(args))
	for i, v := range args {
		arr[i] = zap.Any(strconv.Itoa(i), v)
	}
	return arr
}

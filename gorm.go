package blackdatura

import (
	"context"
	"go.uber.org/zap"
	gl "gorm.io/gorm/logger"
	"strconv"
	"time"
)

type GormLogger struct {
	log *zap.Logger
}

func NewGormLogger() GormLogger {
	return GormLogger{
		log: With("gorm log"),
	}
}

func (l GormLogger) LogMode(gl.LogLevel) gl.Interface {
	return l
}

func (l GormLogger) Info(ctx context.Context, str string, args ...interface{}) {
	s := make([]zap.Field, len(args))
	for i, v := range args {
		s[i] = zap.Any(strconv.Itoa(i), v)
	}
	l.log.Info(str, s...)
}

func (l GormLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	s := make([]zap.Field, len(args))
	for i, v := range args {
		s[i] = zap.Any(strconv.Itoa(i), v)
	}
	l.log.Warn(str, s...)
}

func (l GormLogger) Error(ctx context.Context, str string, args ...interface{}) {
	s := make([]zap.Field, len(args))
	for i, v := range args {
		s[i] = zap.Any(strconv.Itoa(i), v)
	}
	l.log.Error(str, s...)
}

func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()
	l.log.Error("trace", zap.Error(err), zap.Duration("elapsed", time.Since(begin)),
		zap.Int64("rows", rows), zap.String("sql", sql))
}

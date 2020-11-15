package blackdatura

import (
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/url"
)

type LumberjackSink struct {
	*lumberjack.Logger
}

// Sync implements zap.Sink. The remaining methods are implemented
// by the embedded *lumberjack.Logger.
func (LumberjackSink) Sync() error { return nil }

// String return name
func (LumberjackSink) String() string { return "lumberjack" }

// Sink instance
func (l LumberjackSink) Sink(*url.URL) (zap.Sink, error) {
	return l, nil
}

// Lumberjack create lumberjack sink instance
func Lumberjack(logPath string, MaxSize, MaxBackups, MaxAge int, compress bool) LumberjackSink {
	return LumberjackClient(
		&lumberjack.Logger{
			Filename:   logPath,
			MaxSize:    MaxSize,
			MaxBackups: MaxBackups,
			MaxAge:     MaxAge,
			Compress:   compress,
		},
	)
}

// DefaultLumberjack return default config
func DefaultLumberjack() LumberjackSink {
	return Lumberjack("/var/logs/log", 1024, 30, 90, true)
}

// Lumberjack create lumberjack sink instance
func LumberjackClient(l *lumberjack.Logger) LumberjackSink {
	return LumberjackSink{Logger: l}
}

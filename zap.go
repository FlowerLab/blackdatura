package blackdatura

import (
	"fmt"
	"net/url"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerSink interface {
	String() string
	Sink(*url.URL) (zap.Sink, error)
}

// With creates a child logger and adds structured context to it
func With(name string) *zap.Logger {
	return logger.With(zap.String("parent", name))
}

// New return logger
func New() *zap.Logger {
	if logger == nil {
		panic("logger is nil")
	}
	return logger
}

// Set zap logger
func Set(l *zap.Logger) {
	logger = l
}

var logger *zap.Logger

// Init logger
func Init(level string, dev bool, fn ...LoggerSink) {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:       "msg",
		LevelKey:         "level",
		TimeKey:          "ts",
		NameKey:          "logger",
		CallerKey:        "caller",
		FunctionKey:      "",
		StacktraceKey:    "",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.LowercaseLevelEncoder,
		EncodeTime:       zapcore.ISO8601TimeEncoder,
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		EncodeName:       nil,
		ConsoleSeparator: "",
	}

	outputPaths := make([]string, 0, len(fn)+1)

	for _, v := range fn {
		if err := zap.RegisterSink(v.String(), v.Sink); err != nil {
			panic(err)
		}
		outputPaths = append(outputPaths, v.String()+":/tmp")
	}

	if dev {
		outputPaths = append(outputPaths, "stdout")
	}

	loggerConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapLevel(level)),
		Development:      dev,
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
		OutputPaths:      outputPaths,
		ErrorOutputPaths: []string{"stderr"},
	}

	var err error
	logger, err = loggerConfig.Build()
	if err != nil {
		panic(fmt.Sprintf("build zap logger from config error: %v", err))
	}
	zap.ReplaceGlobals(logger)
}

func zapLevel(s string) zapcore.Level {
	switch s {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	}
	return zap.FatalLevel
}

package zap

import (
	core "github.com/bradenrayhorn/ledger-core"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	l *zap.Logger
}

func CreateLogger(cfg *core.Config) core.Logger {
	zapConfig := zap.NewProductionConfig()
	zapConfig.Level = zap.NewAtomicLevelAt(zapcore.Level(cfg.LogLevel))
	zapConfig.Encoding = string(cfg.LogFormat)
	logger, err := zapConfig.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}

	return &ZapLogger{
		l: logger,
	}
}

func (z ZapLogger) Debug(msg string, args ...interface{}) {
	z.l.Sugar().Debugw(msg, args...)
}

func (z ZapLogger) Info(msg string, args ...interface{}) {
	z.l.Sugar().Infow(msg, args...)
}

func (z ZapLogger) Warn(msg string, args ...interface{}) {
	z.l.Sugar().Warnw(msg, args...)
}

func (z ZapLogger) Error(msg string, args ...interface{}) {
	z.l.Sugar().Errorw(msg, args...)
}

func (z ZapLogger) Fatal(msg string, args ...interface{}) {
	z.l.Sugar().Fatalw(msg, args...)
}

func (z ZapLogger) Panic(msg string, args ...interface{}) {
	z.l.Sugar().Panicw(msg, args...)
}

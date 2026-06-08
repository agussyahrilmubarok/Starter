package logger

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ctxKey struct{}

var (
	mu  sync.Mutex
	log *zap.Logger
)

func Init(filePath string, level zapcore.Level) error {
	mu.Lock()
	defer mu.Unlock()

	if err := os.MkdirAll(filepath.Dir(filePath), 0o755); err != nil {
		return fmt.Errorf("failed to create log directory: %w", err)
	}

	lvl := zap.NewAtomicLevelAt(level)

	consoleCfg := zap.NewDevelopmentEncoderConfig()
	consoleCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleCfg.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	consoleCfg.EncodeCaller = zapcore.ShortCallerEncoder

	fileCfg := zap.NewProductionEncoderConfig()
	fileCfg.TimeKey = "timestamp"
	fileCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	fileCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	fileCfg.EncodeCaller = zapcore.ShortCallerEncoder

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(consoleCfg), zapcore.AddSync(os.Stdout), lvl),
		zapcore.NewCore(zapcore.NewJSONEncoder(fileCfg), zapcore.AddSync(&lumberjack.Logger{
			Filename:   filePath,
			MaxSize:    5,
			MaxBackups: 10,
			MaxAge:     14,
			Compress:   true,
		}), lvl),
	)

	if log != nil {
		_ = log.Sync()
	}
	log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return nil
}

func Get() *zap.Logger {
	mu.Lock()
	defer mu.Unlock()
	if log != nil {
		return log
	}
	return zap.NewNop()
}

func Sync() {
	mu.Lock()
	defer mu.Unlock()
	if log != nil {
		_ = log.Sync()
	}
}

func Debug(msg string, fields ...zap.Field) { Get().Debug(msg, fields...) }
func Info(msg string, fields ...zap.Field)  { Get().Info(msg, fields...) }
func Warn(msg string, fields ...zap.Field)  { Get().Warn(msg, fields...) }
func Error(msg string, fields ...zap.Field) { Get().Error(msg, fields...) }
func Fatal(msg string, fields ...zap.Field) { Get().Fatal(msg, fields...) }

func FromCtx(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok {
		return l
	}
	return Get()
}

func WithCtx(ctx context.Context, l *zap.Logger) context.Context {
	if lp, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok && lp == l {
		return ctx
	}
	return context.WithValue(ctx, ctxKey{}, l)
}

// debug, info, warn, error, fatal
func ParseLevel(level string) zapcore.Level {
	var l zapcore.Level
	if err := l.UnmarshalText([]byte(level)); err != nil {
		return zapcore.InfoLevel
	}
	return l
}

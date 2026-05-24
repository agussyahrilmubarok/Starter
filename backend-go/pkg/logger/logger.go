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
	mu     sync.Mutex
	logger *zap.Logger
)

type Options struct {
	LogFilePath string
	LogLevel    zapcore.Level
	MaxSize     int // MB
	MaxBackups  int
	MaxAge      int // days
	Compress    bool
}

func DefaultOptions() Options {
	return Options{
		LogFilePath: "logs/app.log",
		LogLevel:    zapcore.DebugLevel,
		MaxSize:     5,
		MaxBackups:  10,
		MaxAge:      14,
		Compress:    true,
	}
}

// Init initializes the logger with the given options.
// Must be called once at application startup.
func Init(opts Options) error {
	mu.Lock()
	defer mu.Unlock()

	// log directory
	if err := os.MkdirAll(filepath.Dir(opts.LogFilePath), 0o755); err != nil {
		return fmt.Errorf("logger: failed to create log directory: %w", err)
	}

	// log level
	logLevel := zap.NewAtomicLevelAt(opts.LogLevel)

	// console encoder (human-readable + colored)
	consoleCfg := zap.NewDevelopmentEncoderConfig()
	consoleCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleCfg.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	consoleCfg.EncodeCaller = zapcore.ShortCallerEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(consoleCfg)

	// file encoder (JSON, no color)
	fileCfg := zap.NewProductionEncoderConfig()
	fileCfg.TimeKey = "timestamp"
	fileCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	fileCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	fileCfg.EncodeCaller = zapcore.ShortCallerEncoder
	fileEncoder := zapcore.NewJSONEncoder(fileCfg)

	// log rotation (lumberjack)
	rotation := zapcore.AddSync(&lumberjack.Logger{
		Filename:   opts.LogFilePath,
		MaxSize:    opts.MaxSize,
		MaxBackups: opts.MaxBackups,
		MaxAge:     opts.MaxAge,
		Compress:   opts.Compress,
	})

	// tee: write to both console and file
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), logLevel),
		zapcore.NewCore(fileEncoder, rotation, logLevel),
	)

	if logger != nil {
		_ = logger.Sync()
	}

	logger = zap.New(core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	return nil
}

// Get returns the global logger instance.
// Falls back to a no-op logger if Init has not been called.
func Get() *zap.Logger {
	mu.Lock()
	defer mu.Unlock()

	if logger != nil {
		return logger
	}
	return zap.NewNop()
}

// Sync flushes any buffered log entries.
// Should be deferred in main: defer logger.Sync()
func Sync() {
	mu.Lock()
	defer mu.Unlock()

	if logger != nil {
		_ = logger.Sync()
	}
}

// FromCtx returns the Logger associated with the ctx.
// Falls back to the global logger, or a no-op if not initialized.
func FromCtx(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok {
		return l
	}
	return Get()
}

// WithCtx returns a copy of ctx with the Logger attached.
func WithCtx(ctx context.Context, l *zap.Logger) context.Context {
	if lp, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok {
		if lp == l {
			return ctx
		}
	}
	return context.WithValue(ctx, ctxKey{}, l)
}

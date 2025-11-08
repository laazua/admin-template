package xlog

import (
	"log/slog"
	"os"

	"adtmp/pkg/config"
)

func Set() *slog.Logger {
	var level slog.Level
	switch config.Get().LogLevel {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}
	var handler slog.Handler
	opts := &slog.HandlerOptions{}

	if config.Get().LogFormat == "json" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}
	logger := slog.New(handler)
	// 设置默认 logger 的级别
	slog.SetLogLoggerLevel(level)

	return logger
}

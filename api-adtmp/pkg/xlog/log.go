package xlog

import (
	"log/slog"
	"strings"
)

type logWriter struct {
	logger *slog.Logger
}

// Write 实现 io.Writer 接口，使标准库 log 输出转发到 slog
// 日志级别设置示例: log.Printf("ERR: xxx %v", 12)
func (lw logWriter) Write(p []byte) (n int, err error) {
	msg := strings.TrimSpace(string(p))
	switch {
	case strings.HasPrefix(msg, "ERROR") || strings.HasPrefix(msg, "ERR"):
		lw.logger.Error(msg)
	case strings.HasPrefix(msg, "WARN"):
		lw.logger.Warn(msg)
	case strings.HasPrefix(msg, "DEBUG"):
		lw.logger.Debug(msg)
	default:
		lw.logger.Info(msg)
	}
	return len(p), nil
}

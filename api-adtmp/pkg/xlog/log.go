package xlog

import (
	"log/slog"
	"strings"
)

const (
	ErrFlag   = "ERR "
	WarnFlag  = "WARN "
	DebugFlag = "DEBUG "
)

type logWriter struct {
	logger *slog.Logger
}

// Write 实现 io.Writer 接口，使标准库 log 输出转发到 slog
// 日志级别设置示例: log.Printf("ERR xxx %v", 12)
// 日志级别通过前缀识别: 'ERR ', 'WARN ', 'DEBUG '
func (lw logWriter) Write(p []byte) (n int, err error) {
	msg := strings.TrimSpace(string(p))
	switch {
	case strings.HasPrefix(msg, ErrFlag):
		lw.logger.Error(strings.TrimPrefix(msg, ErrFlag))
	case strings.HasPrefix(msg, WarnFlag):
		lw.logger.Warn(strings.TrimPrefix(msg, WarnFlag))
	case strings.HasPrefix(msg, DebugFlag):
		lw.logger.Debug(strings.TrimPrefix(msg, DebugFlag))
	default:
		lw.logger.Info(msg)
	}
	return len(p), nil
}

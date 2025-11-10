package xlog

import (
	"context"
	"io"
	"log/slog"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"adtmp/pkg/config"
)

// 颜色常量
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

// 检查是否支持颜色输出
func supportsColor() bool {
	// 检查环境变量
	if os.Getenv("NO_COLOR") != "" {
		return false
	}

	// 检查是否是 TTY
	if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) == 0 {
		return false
	}

	// Windows 特殊处理
	if runtime.GOOS == "windows" {
		return os.Getenv("TERM") != "" || os.Getenv("WT_SESSION") != "" || os.Getenv("ConEmuANSI") == "ON"
	}

	return true
}

// 自定义文本处理器，避免颜色代码被转义
type colorTextHandler struct {
	opts    slog.HandlerOptions
	w       io.Writer
	groups  []string
	attrs   []slog.Attr
	noColor bool
}

func newColorTextHandler(w io.Writer, opts *slog.HandlerOptions, noColor bool) *colorTextHandler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}
	return &colorTextHandler{
		w:       w,
		opts:    *opts,
		noColor: noColor,
	}
}

func (h *colorTextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

func (h *colorTextHandler) WithGroup(name string) slog.Handler {
	return &colorTextHandler{
		w:       h.w,
		opts:    h.opts,
		groups:  append(h.groups, name),
		noColor: h.noColor,
	}
}

func (h *colorTextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &colorTextHandler{
		w:       h.w,
		opts:    h.opts,
		groups:  h.groups,
		attrs:   append(h.attrs, attrs...),
		noColor: h.noColor,
	}
}

func (h *colorTextHandler) Handle(ctx context.Context, r slog.Record) error {
	var buf []byte

	// 时间（固定宽度）
	timeStr := h.formatTime(r.Time)
	buf = append(buf, timeStr...)
	buf = append(buf, ' ')

	// 级别（固定宽度）
	levelStr := h.formatLevel(r.Level)
	buf = append(buf, levelStr...)
	buf = append(buf, ' ')

	// 源文件信息（如果启用）
	if h.opts.AddSource && r.PC != 0 {
		sourceStr := h.formatSource(r.PC)
		buf = append(buf, sourceStr...)
		buf = append(buf, ' ')
	}

	// 消息
	msgStr := h.formatMessage(r.Message)
	buf = append(buf, msgStr...)

	// 属性
	attrsCount := 0
	r.Attrs(func(attr slog.Attr) bool {
		if attrsCount == 0 {
			buf = append(buf, ' ')
		} else {
			buf = append(buf, ' ')
		}
		buf = append(buf, h.formatAttr(attr)...)
		attrsCount++
		return true
	})

	// 额外的属性
	for _, attr := range h.attrs {
		if attrsCount == 0 {
			buf = append(buf, ' ')
		} else {
			buf = append(buf, ' ')
		}
		buf = append(buf, h.formatAttr(attr)...)
		attrsCount++
	}

	buf = append(buf, '\n')
	_, err := h.w.Write(buf)
	return err
}

func (h *colorTextHandler) formatTime(t time.Time) []byte {
	timeStr := t.Format("2006-01-02 15:04:05.000")
	if h.noColor {
		return []byte(timeStr)
	}
	return []byte(colorWhite + timeStr + colorReset)
}

func (h *colorTextHandler) formatLevel(level slog.Level) []byte {
	var levelStr string
	var color string

	switch level {
	case slog.LevelDebug:
		levelStr = "DEBUG"
		color = colorCyan
	case slog.LevelInfo:
		levelStr = "INFO "
		color = colorGreen
	case slog.LevelWarn:
		levelStr = "WARN "
		color = colorYellow
	case slog.LevelError:
		levelStr = "ERROR"
		color = colorRed
	default:
		levelStr = level.String()
		if len(levelStr) < 5 {
			levelStr += strings.Repeat(" ", 5-len(levelStr))
		}
		color = colorWhite
	}

	if h.noColor {
		return []byte(levelStr)
	}
	return []byte(color + levelStr + colorReset)
}

func (h *colorTextHandler) formatMessage(msg string) []byte {
	if h.noColor {
		return []byte(msg)
	}
	return []byte(colorWhite + msg + colorReset)
}

func (h *colorTextHandler) formatSource(pc uintptr) []byte {
	fs := runtime.CallersFrames([]uintptr{pc})
	f, _ := fs.Next()
	if f.File == "" {
		return []byte{}
	}

	// 使用完整文件路径，但可以调整显示格式
	var source string
	if h.noColor {
		source = f.File + ":" + strconv.Itoa(f.Line)
	} else {
		source = colorGreen + f.File + ":" + strconv.Itoa(f.Line) + colorReset
	}

	return []byte(source)
}

func (h *colorTextHandler) formatAttr(attr slog.Attr) []byte {
	key := attr.Key
	value := attr.Value

	// 敏感信息过滤
	sensitiveKeys := []string{"password", "token", "secret", "authorization", "apikey"}
	for _, sensitive := range sensitiveKeys {
		if strings.Contains(strings.ToLower(key), sensitive) {
			value = slog.StringValue("***")
			break
		}
	}

	var buf []byte
	if h.noColor {
		buf = append(buf, key...)
		buf = append(buf, '=')
		buf = append(buf, h.formatValue(value)...)
	} else {
		buf = append(buf, colorWhite...)
		buf = append(buf, key...)
		buf = append(buf, '=')
		buf = append(buf, h.formatValue(value)...)
		buf = append(buf, colorReset...)
	}
	return buf
}

func (h *colorTextHandler) formatValue(v slog.Value) []byte {
	switch v.Kind() {
	case slog.KindString:
		return h.formatString(v.String())
	case slog.KindTime:
		return []byte(v.Time().Format("2006-01-02 15:04:05.000"))
	default:
		return []byte(v.String())
	}
}

func (h *colorTextHandler) formatString(s string) []byte {
	// 如果字符串需要引号，就加上引号
	if needsQuoting(s) {
		return []byte(strconv.Quote(s))
	}
	return []byte(s)
}

func needsQuoting(s string) bool {
	if s == "" {
		return true
	}
	for _, r := range s {
		if r < 0x20 || r == 0x7f || r == '"' || r == '=' {
			return true
		}
		if !utf8.ValidRune(r) {
			return true
		}
	}
	return false
}

func createHandlerOptions() *slog.HandlerOptions {
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

	return &slog.HandlerOptions{
		Level:     level,
		AddSource: config.Get().LogSource,
	}
}

func Set() {
	opts := createHandlerOptions()

	var handler slog.Handler

	if config.Get().LogFormat == "json" {
		// JSON 格式使用标准处理器
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		// 文本格式使用自定义处理器
		noColor := !supportsColor() || os.Getenv("NO_COLOR") != ""
		handler = newColorTextHandler(os.Stdout, opts, noColor)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
}

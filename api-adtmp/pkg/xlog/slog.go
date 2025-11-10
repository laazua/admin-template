package xlog

import (
	"context"
	"io"
	"log"
	"log/slog"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"adtmp/pkg/config"
)

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
	if os.Getenv("NO_COLOR") != "" {
		return false
	}
	if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) == 0 {
		return false
	}
	if runtime.GOOS == "windows" {
		return os.Getenv("TERM") != "" || os.Getenv("WT_SESSION") != "" || os.Getenv("ConEmuANSI") == "ON"
	}
	return true
}

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

	// TIME=...
	buf = append(buf, h.kv("Time", h.formatTime(r.Time))...)

	// LEVEL=...
	buf = append(buf, ' ')
	buf = append(buf, h.kv("Level", h.formatLevel(r.Level))...)

	// Source=...
	if h.opts.AddSource && r.PC != 0 {
		buf = append(buf, ' ')
		buf = append(buf, h.kv("Source", h.formatSource(r.PC))...)
	}

	// Message=...
	buf = append(buf, ' ')
	buf = append(buf, h.kv("Message", h.formatMessage(r.Message))...)

	// Attrs
	r.Attrs(func(attr slog.Attr) bool {
		buf = append(buf, ' ')
		buf = append(buf, h.formatAttr(attr)...)
		return true
	})
	for _, attr := range h.attrs {
		buf = append(buf, ' ')
		buf = append(buf, h.formatAttr(attr)...)
	}

	buf = append(buf, '\n')
	_, err := h.w.Write(buf)
	return err
}

func (h *colorTextHandler) kv(key string, val []byte) []byte {
	if h.noColor {
		return []byte(key + "=" + string(val))
	}
	return []byte(colorCyan + key + colorReset + "=" + string(val))
}

func (h *colorTextHandler) formatTime(t time.Time) []byte {
	s := t.Format("2006-01-02 15:04:05.000")
	if h.noColor {
		return []byte(s)
	}
	return []byte(colorWhite + s + colorReset)
}

func (h *colorTextHandler) formatLevel(level slog.Level) []byte {
	var s, c string
	switch level {
	case slog.LevelDebug:
		s, c = "DEBUG", colorCyan
	case slog.LevelInfo:
		s, c = "INFO", colorGreen
	case slog.LevelWarn:
		s, c = "WARN", colorYellow
	case slog.LevelError:
		s, c = "ERROR", colorRed
	default:
		s, c = level.String(), colorWhite
	}
	if h.noColor {
		return []byte(s)
	}
	return []byte(c + s + colorReset)
}

func (h *colorTextHandler) formatSource(pc uintptr) []byte {
	fs := runtime.CallersFrames([]uintptr{pc})
	f, _ := fs.Next()
	if f.File == "" {
		return []byte("unknown")
	}
	src := f.File + ":" + strconv.Itoa(f.Line)
	if h.noColor {
		return []byte(src)
	}
	return []byte(colorGreen + src + colorReset)
}

func (h *colorTextHandler) formatMessage(msg string) []byte {
	if h.noColor {
		return []byte(msg)
	}
	return []byte(colorWhite + msg + colorReset)
}

func (h *colorTextHandler) formatAttr(attr slog.Attr) []byte {
	key := attr.Key
	value := attr.Value

	sensitive := []string{"password", "token", "secret", "authorization", "apikey"}
	for _, s := range sensitive {
		if strings.Contains(strings.ToLower(key), s) {
			value = slog.StringValue("***")
			break
		}
	}
	if h.noColor {
		return []byte(key + "=" + string(h.formatValue(value)))
	}
	return []byte(colorCyan + key + colorReset + "=" + string(h.formatValue(value)))
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
	switch strings.ToLower(config.Get().LogLevel) {
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
	noColor := !supportsColor() || os.Getenv("NO_COLOR") != ""

	var handler slog.Handler
	if config.Get().LogFormat == "json" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = newColorTextHandler(os.Stdout, opts, noColor)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)

	// ---------- 标准库 log 适配 ----------
	// 1. 去掉标准库前缀与时间戳（因为 slog 会统一格式）
	log.SetFlags(0)

	// 2. 让标准库 log 输出经过 slog 处理
	log.SetOutput(logWriter{logger: logger})
}

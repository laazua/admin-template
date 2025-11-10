package cert

import (
	"crypto/tls"
	"log/slog"
	"os"

	"adtmp/pkg/config"
)

// TLSOptions 包含 TLS 配置选项
type TLSOptions struct {
	UseTLS    bool
	TLSConfig *tls.Config
	ServerCrt string
	ServerKey string
}

// SetupTLS 配置 TLS 选项，如果证书文件不存在则返回不使用 TLS
func SetupTLS() *TLSOptions {
	serverCrt := config.Get().ServerCrt
	serverKey := config.Get().ServerKey

	// 检查证书文件是否存在
	if serverCrt == "" || serverKey == "" {
		slog.Warn("未配置证书路径，将使用 HTTP 启动服务")
		return &TLSOptions{UseTLS: false}
	}

	if !fileExists(serverCrt) || !fileExists(serverKey) {
		slog.Warn("证书文件不存在，将使用 HTTP 启动服务",
			slog.String("ServerCrt", serverCrt),
			slog.String("ServerKey", serverKey),
		)
		return &TLSOptions{UseTLS: false}
	}

	// 配置 TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{
			tls.X25519, tls.CurveP256,
		},
		CipherSuites: []uint16{
			// HTTP/2 必需的密码套件
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,

			// 其他安全套件
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		},
	}

	// 加载证书
	cert, err := tls.LoadX509KeyPair(serverCrt, serverKey)
	if err != nil {
		slog.Error("加载证书失败，将使用 HTTP 启动服务", slog.String("Error", err.Error()))
		return &TLSOptions{UseTLS: false}
	}
	slog.Info("加载公钥证书文件成功", slog.String("Crt", serverCrt))
	slog.Info("加载私钥证书文件成功", slog.String("Key", serverKey))
	tlsConfig.Certificates = []tls.Certificate{cert}
	slog.Info("已成功加载证书,将使用 HTTPS 启动服务")

	return &TLSOptions{
		UseTLS:    true,
		TLSConfig: tlsConfig,
		ServerCrt: serverCrt,
		ServerKey: serverKey,
	}
}

// GetProtocol 返回协议名称
func (t *TLSOptions) GetProtocol() string {
	if t.UseTLS {
		return "HTTPS"
	}
	return "HTTP"
}

// fileExists 检查文件是否存在
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

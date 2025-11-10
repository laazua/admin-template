package app

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"adtmp/pkg/cert"
	"adtmp/pkg/config"
	"adtmp/pkg/db"
	"adtmp/pkg/di"
)

// Start 启动 HTTP 服务器
func Start() {
	// 获取服务实例
	mux := http.NewServeMux()
	slog.Info("程序开始启动...")

	// 依赖注入 && 路由注册
	dB, err := db.Get()
	if err != nil {
		slog.Error("数据库连接错误", slog.String("Error", err.Error()))
		os.Exit(-1)
	}

	for _, api := range di.Denps(dB) {
		api.Register(mux)
	}

	// 配置 TLS
	tlsOptions := cert.SetupTLS()

	// 实例化服务
	server := http.Server{
		Handler:      mux,
		Addr:         config.Get().Address,
		ReadTimeout:  config.Get().ReadTimeout,
		WriteTimeout: config.Get().WriteTimeout,
	}

	// 设置 TLS 配置
	if tlsOptions.UseTLS {
		server.TLSConfig = tlsOptions.TLSConfig
	}

	quit := make(chan os.Signal, 1)
	// 监听失败和退出信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		var err error
		if tlsOptions.UseTLS {
			err = server.ListenAndServeTLS("", "")
		} else {
			err = server.ListenAndServe()
		}

		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("启动服务失败", slog.String("Error", err.Error()))
			os.Exit(-2)
		}
	}()

	slog.Info("程序启动成功",
		slog.String("Address", config.Get().Address),
		slog.String("Protocol", tlsOptions.GetProtocol()),
	)

	<-quit
	slog.Info("程序关闭清理资源")
	db.Close()
}

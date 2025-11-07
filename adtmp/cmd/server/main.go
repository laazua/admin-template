package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"adtmp/pkg/config"
	"adtmp/pkg/db"
	"adtmp/pkg/di"
	"adtmp/pkg/xlog"
)

func main() {
	// 设置日志
	xlog.Set()
	// 初始化数据库
	dB, err := db.Get()
	if err != nil {
		slog.Error("数据库连接错误", slog.String("Error", err.Error()))
		os.Exit(-1)
	}
	// 获取服务实例
	mux := http.NewServeMux()
	slog.Info("程序开始启动...")
	// 依赖注入 && 路由注册
	for _, api := range di.Denps(dB) {
		api.Register(mux)
	}
	// 实例化服务
	server := http.Server{
		Handler:      mux,
		Addr:         config.Get().Address,
		ReadTimeout:  config.Get().ReadTimeout,
		WriteTimeout: config.Get().WriteTimeout,
	}

	quit := make(chan os.Signal, 1)
	// 监听失败和退出信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("启动服务失败", slog.String("Error", err.Error()))
			os.Exit(-2)
		}
	}()
	slog.Info("程序启动成功", slog.String("Address", config.Get().Address))
	<-quit
	slog.Info("程序关闭清理资源")
	db.Close()
}

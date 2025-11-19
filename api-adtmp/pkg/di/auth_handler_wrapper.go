package di

import (
	"log/slog"
	"net/http"

	"adtmp/pkg/internal/handler"
)

// 包装器结构体，为每个 handler 添加中间件能力
type authHandlerWithMw struct {
	hd      *handler.AuthHandler
	applier *middlewareApplier
}

func (h *authHandlerWithMw) Register(mux *http.ServeMux) {
	slog.Info("authHandler 注册路由并应用中间件 ...")
	mux.HandleFunc("POST /api/auth/login", h.hd.Login)
	mux.HandleFunc("GET /api/auth/userinfo", h.applier.WithAuth(h.hd.GetUserInfo))
}

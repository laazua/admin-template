package di

import (
	"adtmp/internal/handler"
	"log/slog"
	"net/http"
)

type userHandlerWithMw struct {
	hd      *handler.UserHandler
	applier *middlewareApplier
}

func (h *userHandlerWithMw) Register(mux *http.ServeMux) {
	slog.Info("userHandler 注册路由并应用中间件 ...")
	mux.HandleFunc("POST /api/user", h.applier.WithAuth(h.hd.Create))
	mux.HandleFunc("DELETE /api/user/{id}", h.applier.WithAuth(h.hd.Destroy))
	mux.HandleFunc("PUT /api/user/{id}", h.applier.WithAuth(h.hd.Update))
	mux.HandleFunc("GET /api/user/{id}", h.applier.WithAuth(h.hd.Retrieve))
	mux.HandleFunc("GET /api/user", h.applier.WithAuth(h.hd.List))
}

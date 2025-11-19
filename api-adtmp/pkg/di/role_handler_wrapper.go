package di

import (
	"log/slog"
	"net/http"

	"adtmp/pkg/internal/handler"
)

type roleHandlerWithMw struct {
	hd      *handler.RoleHandler
	applier *middlewareApplier
}

func (h *roleHandlerWithMw) Register(mux *http.ServeMux) {
	slog.Info("roleHandler 注册路由并应用中间件 ...")
	mux.HandleFunc("POST /api/role", h.applier.WithAuth(h.hd.Create))
	mux.HandleFunc("DELETE /api/role/{id}", h.applier.WithAuth(h.hd.Destroy))
	mux.HandleFunc("PUT /api/role/{id}", h.applier.WithAuth(h.hd.Update))
	mux.HandleFunc("GET /api/role/{id}", h.applier.WithAuth(h.hd.Retrieve))
	mux.HandleFunc("GET /api/role", h.applier.WithAuth(h.hd.List))
}

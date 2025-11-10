package handler

import (
	"log/slog"
	"net/http"

	"adtmp/internal/domain/entities/form"
	"adtmp/internal/service"
	"adtmp/pkg/api"
	"adtmp/pkg/middleware"
)

type authHandler struct {
	mw          middleware.Mw
	authService service.AuthService
}

func NewAuthHandler(authMw middleware.Mw, authService service.AuthService) *authHandler {
	return &authHandler{mw: authMw, authService: authService}
}

// 应用中间件
func (h *authHandler) withMws(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

// 路由注册
func (h *authHandler) Register(mux *http.ServeMux) {
	slog.Info("authHandler 注册路由并应用中间件 ...")
	mux.HandleFunc("POST /api/auth/login", h.login)
	mux.HandleFunc("GET /api/auth/{name}", h.withMws(
		h.getUserInfo,
		h.mw.Auth,
	))
}

// 登录认证接口
func (h *authHandler) login(w http.ResponseWriter, r *http.Request) {
	var user form.UserLogin
	err := api.BindJSON(r, &user)
	if err != nil {
		slog.Error("authHandler 请求参数错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusBadRequest, "authHandler 请求参数错误")
		return
	}
	token, err := h.authService.AuthUser(r.Context(), &user)
	if err != nil {
		slog.Error("authHandler 用户认证错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusInternalServerError, "authHandler 用户认证错误")
		return
	}
	api.Success(w, api.M{"token": token}, "登录成功")
}

// 用户路由信息接口
func (h *authHandler) getUserInfo(w http.ResponseWriter, r *http.Request) {}

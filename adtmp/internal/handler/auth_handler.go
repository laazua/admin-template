package handler

import (
	"log/slog"
	"net/http"

	"adtmp/internal/domain/entities/form"
	"adtmp/internal/service"
	"adtmp/pkg/api"
)

type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *authHandler {
	return &authHandler{authService: authService}
}

func (authHandler *authHandler) Register(mux *http.ServeMux) {
	slog.Info("注册认证路由...")
	mux.HandleFunc("POST /api/auth/login", authHandler.login)
	mux.HandleFunc("GET /api/auth/{name}", authHandler.getUserInfo)
}

func (authHandler *authHandler) login(w http.ResponseWriter, r *http.Request) {
	var user form.UserLogin
	err := api.BindJSON(r, &user)
	if err != nil {
		slog.Error("认证处理器请求参数错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusBadRequest, "认证处理器请求参数错误")
		return
	}
	token, err := authHandler.authService.AuthUser(r.Context(), &user)
	if err != nil {
		slog.Error("认证处理器用户认证错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusInternalServerError, "认证处理器用户认证错误")
		return
	}
	api.Success(w, api.M{"token": token}, "登录成功")
}

func (authHandler *authHandler) getUserInfo(w http.ResponseWriter, r *http.Request) {}

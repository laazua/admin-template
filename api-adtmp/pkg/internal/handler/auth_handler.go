package handler

import (
	"log"
	"net/http"

	"adtmp/pkg/api"
	"adtmp/pkg/internal/domain/entities/form"
	"adtmp/pkg/internal/service"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// 登录认证接口
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user form.UserLogin
	err := api.BindJSON(r, &user)
	if err != nil {
		// slog.Error("authHandler 请求参数错误", slog.String("Error", err.Error()))
		log.Printf("ERR authHandler 请求参数错误: %s", err.Error())
		api.Failure(w, http.StatusBadRequest, "authHandler 请求参数错误")
		return
	}
	log.Printf("用户 %v 尝试登录 ...", user.Email)
	token, err := h.authService.AuthUser(r.Context(), &user)
	if err != nil {
		// slog.Error("authHandler 用户认证错误", slog.String("Error", err.Error()))
		log.Printf("ERR authHandler 用户认证错误: %s", err.Error())
		api.Failure(w, http.StatusInternalServerError, "authHandler 用户认证错误")
		return
	}
	log.Printf("%s 登录成功", user.Email)
	api.Success(w, api.M{"token": token}, "登录成功")
}

// 用户路由信息接口
func (h *AuthHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	api.Success(w, api.M{"xxx": "250"}, "")
}

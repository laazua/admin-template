package handler

import (
	"log/slog"
	"net/http"

	"adtmp/internal/domain/entities/form"
	"adtmp/internal/service"
	"adtmp/pkg/api"
	"adtmp/pkg/middleware"
)

type userHandler struct {
	mw          middleware.Mw
	userService service.UserService
}

func NewUserHandler(mw middleware.Mw, userService service.UserService) *userHandler {
	return &userHandler{mw: mw, userService: userService}
}

func (h *userHandler) withMws(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

func (h *userHandler) Register(mux *http.ServeMux) {
	slog.Info("userHandler 注册路由并应用中间件 ...")
	mux.HandleFunc("POST /api/user", h.withMws(
		h.create,
		h.mw.Auth,
	))
	mux.HandleFunc("DELETE /api/user/{id}", h.withMws(
		h.destroy,
		h.mw.Auth,
	))
	mux.HandleFunc("PUT /api/user/{id}", h.withMws(
		h.update,
		h.mw.Auth,
	))
	mux.HandleFunc("GET /api/user/{id}", h.withMws(
		h.retrieve,
		h.mw.Auth,
	))
	mux.HandleFunc("GET /api/user", h.withMws(
		h.list,
		h.mw.Auth,
	))
}

func (h *userHandler) create(w http.ResponseWriter, r *http.Request) {
	var user form.UserCreate
	err := api.BindJSON(r, &user)
	if err != nil {
		slog.Error("userHandler 请求参数错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusBadRequest, "userHandler 请求参数错误")
		return
	}
	err = h.userService.CreateUser(r.Context(), &user)
	if err != nil {
		slog.Error("userHandler 新增用户错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusInternalServerError, "userHandler 新增用户错误")
		return
	}
	api.Success(w, nil, "新增用户成功")
}

func (h *userHandler) destroy(w http.ResponseWriter, r *http.Request) {
	userId, err := api.StrToUint(r.PathValue("id"))
	if err != nil {
		slog.Error("userHandler 路径参数错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusBadRequest, "userHandler 路径参数错误")
		return
	}
	if err := h.userService.DestroyUser(r.Context(), userId); err != nil {
		slog.Error("userHandler 删除用户错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusInternalServerError, "userHandler 删除用户错误")
		return
	}
	api.Success(w, nil, "删除用户成功")
}

func (h *userHandler) update(w http.ResponseWriter, r *http.Request) {
	userId, err := api.StrToUint(r.PathValue("id"))
	if err != nil {
		slog.Error("userHandler 路径参数错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusBadRequest, "userHandler 路径参数错误")
		return
	}
	var user form.UserUpdate
	if err := api.BindJSON(r, &user); err != nil {
		slog.Error("userHandler 请求参数错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusBadRequest, "userHandler 请求参数错误")
		return
	}
	err = h.userService.UpdateUser(r.Context(), userId, &user)
	if err != nil {
		slog.Error("userHandler 更新用户错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusInternalServerError, "userHandler 更新用户错误")
		return
	}
	api.Success(w, nil, "更新用户成功")
}

func (h *userHandler) retrieve(w http.ResponseWriter, r *http.Request) {
	userId, err := api.StrToUint(r.PathValue("id"))
	if err != nil {
		slog.Error("userHandler 路径参数类型转换错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusBadRequest, "userHandler 路径参数错误")
		return
	}
	userResp, err := h.userService.GetById(r.Context(), userId)
	if err != nil {
		slog.Error("userHandler 检索用户错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusInternalServerError, "userHandler 检索用户错误")
		return
	}
	api.Success(w, userResp, "检索用户成功")
}

func (h *userHandler) list(w http.ResponseWriter, r *http.Request) {
	limit := 0
	offset := 10
	h.userService.ListUser(r.Context(), limit, offset)
}

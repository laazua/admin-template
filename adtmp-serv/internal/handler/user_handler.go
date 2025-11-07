package handler

import (
	"log/slog"
	"net/http"

	"adtmp/internal/domain/entities/form"
	"adtmp/internal/service"
	"adtmp/pkg/api"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService: userService}
}

func (userHandler *userHandler) Register(mux *http.ServeMux) {
	slog.Info("注册用户路由...")
	mux.HandleFunc("POST /api/user", userHandler.create)
	mux.HandleFunc("DELETE /api/user/{id}", userHandler.destroy)
	mux.HandleFunc("PUT /api/user/{id}", userHandler.update)
	mux.HandleFunc("GET /api/user/{id}", userHandler.retrieve)
	mux.HandleFunc("GET /api/user", userHandler.list)
}

func (userHandler *userHandler) create(w http.ResponseWriter, r *http.Request) {
	var user form.UserCreate
	err := api.BindJSON(r, &user)
	if err != nil {
		slog.Error("用户处理器请求参数错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusBadRequest, "用户处理器请求参数错误")
		return
	}
	err = userHandler.userService.CreateUser(r.Context(), &user)
	if err != nil {
		slog.Error("用户处理器新增用户错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusInternalServerError, "用户处理器新增用户错误")
		return
	}
	api.Success(w, nil, "新增用户成功")
}

func (userHandler *userHandler) destroy(w http.ResponseWriter, r *http.Request) {
	userId, err := api.StrToUint(r.PathValue("id"))
	if err != nil {
		slog.Error("用户处理器路径参数错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusBadRequest, "用户处理器路径参数错误")
		return
	}
	if err := userHandler.userService.DestroyUser(r.Context(), userId); err != nil {
		slog.Error("用户处理器删除用户错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusInternalServerError, "用户处理器删除用户错误")
		return
	}
	api.Success(w, nil, "删除用户成功")
}

func (userHandler *userHandler) update(w http.ResponseWriter, r *http.Request) {
	userId, err := api.StrToUint(r.PathValue("id"))
	if err != nil {
		slog.Error("用户处理器路径参数错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusBadRequest, "用户处理器路径参数错误")
		return
	}
	var user form.UserUpdate
	if err := api.BindJSON(r, &user); err != nil {
		slog.Error("用户处理器请求参数错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusBadRequest, "用户处理器请求参数错误")
		return
	}
	err = userHandler.userService.UpdateUser(r.Context(), userId, &user)
	if err != nil {
		slog.Error("用户处理器更新用户错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusInternalServerError, "用户处理器更新用户错误")
		return
	}
	api.Success(w, nil, "更新用户成功")
}

func (userHandler *userHandler) retrieve(w http.ResponseWriter, r *http.Request) {
	userId, err := api.StrToUint(r.PathValue("id"))
	if err != nil {
		slog.Error("路径参数类型转换错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusBadRequest, "路径参数错误")
		return
	}
	userResp, err := userHandler.userService.GetById(r.Context(), userId)
	if err != nil {
		slog.Error("检索用户错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusInternalServerError, "检索用户错误")
		return
	}
	api.Success(w, userResp, "检索用户成功")
}

func (userHandler *userHandler) list(w http.ResponseWriter, r *http.Request) {
	limit := 0
	offset := 10
	userHandler.userService.ListUser(r.Context(), limit, offset)
}

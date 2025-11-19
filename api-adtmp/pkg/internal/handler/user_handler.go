package handler

import (
	"log"
	"net/http"

	"adtmp/pkg/api"
	"adtmp/pkg/internal/domain/entities/form"
	"adtmp/pkg/internal/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user form.UserCreate
	err := api.BindJSON(r, &user)
	if err != nil {
		// slog.Error("userHandler 请求参数错误", slog.String("Error", err.Error()))
		log.Printf("Err userHandler 请求参数错误: %s", err.Error())
		api.Failure(w, http.StatusBadRequest, "userHandler 请求参数错误")
		return
	}
	err = h.userService.CreateUser(r.Context(), &user)
	if err != nil {
		// slog.Error("userHandler 新增用户错误", slog.String("Error", err.Error()))
		log.Printf("Err userHandler 新增用户错误: %s", err.Error())
		api.Failure(w, http.StatusInternalServerError, "userHandler 新增用户错误")
		return
	}
	api.Success(w, nil, "新增用户成功")
}

func (h *UserHandler) Destroy(w http.ResponseWriter, r *http.Request) {
	userId, err := api.StrToUint(r.PathValue("id"))
	if err != nil {
		// slog.Error("userHandler 路径参数错误", slog.String("Error", err.Error()))
		log.Printf("Err userHandler 路径参数错误: %s", err.Error())
		api.Failure(w, http.StatusBadRequest, "userHandler 路径参数错误")
		return
	}
	if err := h.userService.DestroyUser(r.Context(), userId); err != nil {
		// slog.Error("userHandler 删除用户错误", slog.String("Error", err.Error()))
		log.Printf("Err userHandler 删除用户错误: %s", err.Error())
		api.Failure(w, http.StatusInternalServerError, "userHandler 删除用户错误")
		return
	}
	api.Success(w, nil, "删除用户成功")
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	userId, err := api.StrToUint(r.PathValue("id"))
	if err != nil {
		// slog.Error("userHandler 路径参数错误", slog.String("Error", err.Error()))
		log.Printf("Err userHandler 路径参数错误: %s", err.Error())
		api.Failure(w, http.StatusBadRequest, "userHandler 路径参数错误")
		return
	}
	var user form.UserUpdate
	if err := api.BindJSON(r, &user); err != nil {
		// slog.Error("userHandler 请求参数错误", slog.String("Error", err.Error()))
		log.Printf("Err userHandler 请求参数错误: %s", err.Error())
		api.Failure(w, http.StatusBadRequest, "userHandler 请求参数错误")
		return
	}
	err = h.userService.UpdateUser(r.Context(), userId, &user)
	if err != nil {
		// slog.Error("userHandler 更新用户错误", slog.String("Error", err.Error()))
		log.Printf("Err userHandler 更新用户错误: %s", err.Error())
		api.Failure(w, http.StatusInternalServerError, "userHandler 更新用户错误")
		return
	}
	api.Success(w, nil, "更新用户成功")
}

func (h *UserHandler) Retrieve(w http.ResponseWriter, r *http.Request) {
	userId, err := api.StrToUint(r.PathValue("id"))
	if err != nil {
		// slog.Error("userHandler 路径参数类型转换错误", slog.String("Error", err.Error()))
		log.Printf("Err userHandler 路径参数类型转换错误: %s", err.Error())
		api.Failure(w, http.StatusBadRequest, "userHandler 路径参数错误")
		return
	}
	userResp, err := h.userService.GetById(r.Context(), userId)
	if err != nil {
		// slog.Error("userHandler 检索用户错误", slog.String("Error", err.Error()))
		log.Printf("Err userHandler 检索用户错误: %s", err.Error())
		api.Failure(w, http.StatusInternalServerError, "userHandler 检索用户错误")
		return
	}
	api.Success(w, userResp, "检索用户成功")
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	limit := 0
	offset := 10
	h.userService.ListUser(r.Context(), limit, offset)
}

package handler

import (
	"log"
	"net/http"

	"adtmp/pkg/api"
	"adtmp/pkg/internal/domain/entities/form"
	"adtmp/pkg/internal/service"
)

type RoleHandler struct {
	name string

	roleService service.RoleService
}

func NewRoleHandler(roleService service.RoleService) *RoleHandler {
	return &RoleHandler{roleService: roleService, name: "roleHandler"}
}

func (h *RoleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var role form.RoleCreate
	if err := api.BindJSON(r, &role); err != nil {
		// slog.Error("roleHandler 请求参数错误", slog.String("Error", err.Error()))
		log.Printf("Err roleHandler 请求参数错误: %s", err.Error())
		api.Failure(w, http.StatusBadRequest, "roleHandler 请求参数错误")
		return
	}
	if err := h.roleService.CreateRole(r.Context(), &role); err != nil {
		// slog.Error("roleHandler 新增角色错误", slog.String("Error", err.Error()))
		log.Printf("ERR roleHandler 新增角色错误: %s", err.Error())
		api.Failure(w, http.StatusInternalServerError, "roleHandler 新增角色错误")
		return
	}
	api.Success(w, nil, "创建角色成功")
}

func (h *RoleHandler) Destroy(w http.ResponseWriter, r *http.Request) {
	roleId, err := api.StrToUint(r.PathValue("id"))
	if err != nil {
		// slog.Error("roleHandler 路径参数错误", slog.String("Error", err.Error()))
		log.Printf("ERR roleHandler 路径参数错误: %s", err.Error())
		api.Failure(w, http.StatusBadRequest, "roleHandler 路径参数错误")
		return
	}
	if err := h.roleService.DestroyRole(r.Context(), roleId); err != nil {
		// slog.Error("roleHandler 删除角色错误", slog.String("Error", err.Error()))
		log.Printf("ERR roleHandler 删除角色错误: %s", err.Error())
		api.Failure(w, http.StatusInternalServerError, "roleHandler 删除角色错误")
		return
	}
	api.Success(w, nil, "删除角色成功")
}

func (h *RoleHandler) Update(w http.ResponseWriter, r *http.Request) {
	roleId, err := api.StrToUint(r.PathValue("id"))
	if err != nil {
		// slog.Error("roleHandler 路径参数错误", slog.String("Error", err.Error()))
		log.Printf("ERR roleHandler 路径参数错误: %s", err.Error())
		api.Failure(w, http.StatusBadRequest, "roleHandler 路径参数错误")
		return
	}
	var role form.RoleUpdate
	if err := api.BindJSON(r, &role); err != nil {
		// slog.Error("roleHandler 请求参数错误", slog.String("Error", err.Error()))
		log.Printf("ERR roleHandler 请求参数错误: %s", err.Error())
		api.Failure(w, http.StatusBadRequest, "roleHandler 请求参数错误")
		return
	}
	if err := h.roleService.UpdateRole(r.Context(), roleId, &role); err != nil {
		// slog.Error("roleHandler 更新角色错误", slog.String("Error", err.Error()))
		log.Printf("ERR roleHandler 更新角色错误: %s", err.Error())
		api.Failure(w, http.StatusInternalServerError, "roleHandler 更新角色错误")
		return
	}
	api.Success(w, nil, "更新角色成功")
}

func (h *RoleHandler) Retrieve(w http.ResponseWriter, r *http.Request) {
	roleId, err := api.StrToUint(r.PathValue("id"))
	if err != nil {
		// slog.Error("roleHandler 路径参数错误", slog.String("Error", err.Error()))
		log.Printf("ERR roleHandler 路径参数错误: %s", err.Error())
		api.Failure(w, http.StatusBadRequest, "roleHandler 路径参数错误")
		return
	}
	roleResp, err := h.roleService.GetById(r.Context(), roleId)
	if err != nil {
		// slog.Error("roleHandler 检索角色错误", slog.String("Error", err.Error()))
		log.Printf("ERR roleHandler 检索角色错误: %s", err.Error())
		api.Failure(w, http.StatusInternalServerError, "roleHandler 检索角色错误")
		return
	}
	api.Success(w, roleResp, "检索角色成功")
}

func (h *RoleHandler) List(w http.ResponseWriter, r *http.Request) {
	limit := 10
	offset := 10
	h.roleService.ListRole(r.Context(), limit, offset)
}

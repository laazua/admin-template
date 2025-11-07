package handler

import (
	"log/slog"
	"net/http"

	"adtmp/internal/domain/entities/form"
	"adtmp/internal/service"
	"adtmp/pkg/api"
)

type roleHandler struct {
	roleService service.RoleService
}

func NewRoleHandler(roleService service.RoleService) *roleHandler {
	return &roleHandler{roleService: roleService}
}

func (roleHandler *roleHandler) Register(mux *http.ServeMux) {
	slog.Info("注册角色路由...")
	mux.HandleFunc("POST /api/role", roleHandler.create)
	mux.HandleFunc("DELETE /api/role/{id}", roleHandler.destroy)
	mux.HandleFunc("PUT /api/role/{id}", roleHandler.update)
	mux.HandleFunc("GET /api/role/{id}", roleHandler.retrieve)
	mux.HandleFunc("GET /api/role", roleHandler.list)
}

func (roleHandler *roleHandler) create(w http.ResponseWriter, r *http.Request) {
	var role form.RoleCreate
	if err := api.BindJSON(r, &role); err != nil {
		slog.Error("角色处理器请求参数错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusBadRequest, "角色处理器请求参数错误")
		return
	}
	if err := roleHandler.roleService.CreateRole(r.Context(), &role); err != nil {
		slog.Error("角色处理器新增角色错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusInternalServerError, "角色处理器新增角色错误")
		return
	}
	api.Success(w, nil, "创建角色成功")
}

func (roleHandler *roleHandler) destroy(w http.ResponseWriter, r *http.Request) {
	roleId, err := api.StrToUint(r.PathValue("id"))
	if err != nil {
		slog.Error("角色处理器路径参数错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusBadRequest, "角色处理器路径参数错误")
		return
	}
	if err := roleHandler.roleService.DestroyRole(r.Context(), roleId); err != nil {
		slog.Error("角色处理器删除角色错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusInternalServerError, "角色处理器删除角色错误")
		return
	}
	api.Success(w, nil, "删除角色成功")
}

func (roleHandler *roleHandler) update(w http.ResponseWriter, r *http.Request) {
	roleId, err := api.StrToUint(r.PathValue("id"))
	if err != nil {
		slog.Error("角色处理器路径参数错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusBadRequest, "角色处理器路径参数错误")
		return
	}
	var role form.RoleUpdate
	if err := api.BindJSON(r, &role); err != nil {
		slog.Error("角色处理器请求参数错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusBadRequest, "角色处理器请求参数错误")
		return
	}
	if err := roleHandler.roleService.UpdateRole(r.Context(), roleId, &role); err != nil {
		slog.Error("角色处理器更新角色错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusInternalServerError, "角色处理器更新角色错误")
		return
	}
	api.Success(w, nil, "更新角色成功")
}

func (roleHandler *roleHandler) retrieve(w http.ResponseWriter, r *http.Request) {
	roleId, err := api.StrToUint(r.PathValue("id"))
	if err != nil {
		slog.Error("角色处理器路径参数错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusBadRequest, "角色处理器路径参数错误")
		return
	}
	roleResp, err := roleHandler.roleService.GetById(r.Context(), roleId)
	if err != nil {
		slog.Error("角色处理器检索角色错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusInternalServerError, "角色处理器检索角色错误")
		return
	}
	api.Success(w, roleResp, "检索角色成功")
}

func (roleHandler *roleHandler) list(w http.ResponseWriter, r *http.Request) {
	limit := 10
	offset := 10
	roleHandler.roleService.ListRole(r.Context(), limit, offset)
}

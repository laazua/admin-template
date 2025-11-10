package handler

import (
	"log/slog"
	"net/http"

	"adtmp/internal/domain/entities/form"
	"adtmp/internal/service"
	"adtmp/pkg/api"
	"adtmp/pkg/middleware"
)

type roleHandler struct {
	mw          middleware.Mw
	roleService service.RoleService
}

func NewRoleHandler(mw middleware.Mw, roleService service.RoleService) *roleHandler {
	return &roleHandler{mw: mw, roleService: roleService}
}

func (h *roleHandler) withMws(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

func (h *roleHandler) Register(mux *http.ServeMux) {
	slog.Info("roleHandler 注册路由并应用中间件 ...")
	mux.HandleFunc("POST /api/role", h.withMws(
		h.create,
		h.mw.Auth,
	))
	mux.HandleFunc("DELETE /api/role/{id}", h.withMws(
		h.destroy,
		h.mw.Auth,
	))
	mux.HandleFunc("PUT /api/role/{id}", h.withMws(
		h.update,
		h.mw.Auth,
	))
	mux.HandleFunc("GET /api/role/{id}", h.withMws(
		h.retrieve,
		h.mw.Auth,
	))
	mux.HandleFunc("GET /api/role", h.withMws(
		h.list,
		h.mw.Auth,
	))
}

func (h *roleHandler) create(w http.ResponseWriter, r *http.Request) {
	var role form.RoleCreate
	if err := api.BindJSON(r, &role); err != nil {
		slog.Error("roleHandler 请求参数错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusBadRequest, "roleHandler 请求参数错误")
		return
	}
	if err := h.roleService.CreateRole(r.Context(), &role); err != nil {
		slog.Error("roleHandler 新增角色错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusInternalServerError, "roleHandler 新增角色错误")
		return
	}
	api.Success(w, nil, "创建角色成功")
}

func (h *roleHandler) destroy(w http.ResponseWriter, r *http.Request) {
	roleId, err := api.StrToUint(r.PathValue("id"))
	if err != nil {
		slog.Error("roleHandler 路径参数错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusBadRequest, "roleHandler 路径参数错误")
		return
	}
	if err := h.roleService.DestroyRole(r.Context(), roleId); err != nil {
		slog.Error("roleHandler 删除角色错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusInternalServerError, "roleHandler 删除角色错误")
		return
	}
	api.Success(w, nil, "删除角色成功")
}

func (h *roleHandler) update(w http.ResponseWriter, r *http.Request) {
	roleId, err := api.StrToUint(r.PathValue("id"))
	if err != nil {
		slog.Error("roleHandler 路径参数错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusBadRequest, "roleHandler 路径参数错误")
		return
	}
	var role form.RoleUpdate
	if err := api.BindJSON(r, &role); err != nil {
		slog.Error("roleHandler 请求参数错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusBadRequest, "roleHandler 请求参数错误")
		return
	}
	if err := h.roleService.UpdateRole(r.Context(), roleId, &role); err != nil {
		slog.Error("roleHandler 更新角色错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusInternalServerError, "roleHandler 更新角色错误")
		return
	}
	api.Success(w, nil, "更新角色成功")
}

func (h *roleHandler) retrieve(w http.ResponseWriter, r *http.Request) {
	roleId, err := api.StrToUint(r.PathValue("id"))
	if err != nil {
		slog.Error("roleHandler 路径参数错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusBadRequest, "roleHandler 路径参数错误")
		return
	}
	roleResp, err := h.roleService.GetById(r.Context(), roleId)
	if err != nil {
		slog.Error("roleHandler 检索角色错误", slog.String("Error", err.Error()))
		api.Failure(w, http.StatusInternalServerError, "roleHandler 检索角色错误")
		return
	}
	api.Success(w, roleResp, "检索角色成功")
}

func (h *roleHandler) list(w http.ResponseWriter, r *http.Request) {
	limit := 10
	offset := 10
	h.roleService.ListRole(r.Context(), limit, offset)
}

package form

import "adtmp/pkg/internal/domain/entities"

type RoleCreate struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Routes      []entities.Route `json:"routes"`
}

type RoleUpdate struct {
	RoleCreate
}

// --------------------------------------------------------

type RoleForm interface {
	GetName() string
	GetDesc() string
	GetRoutes() []entities.Route
}

// 实现表单方法
func (f *RoleCreate) GetName() string             { return f.Name }
func (f *RoleCreate) GetDesc() string             { return f.Description }
func (f *RoleCreate) GetRoutes() []entities.Route { return f.Routes }
func (f *RoleUpdate) GetName() string             { return f.Name }
func (f *RoleUpdate) GetDesc() string             { return f.Description }
func (f *RoleUpdate) GetRoutes() []entities.Route { return f.Routes }

// 转换函数
func ToRoleDb[T RoleForm](t T) *entities.Role {
	if any(t) == nil {
		return nil
	}

	form := any(t).(RoleForm)
	return &entities.Role{
		Name:        form.GetName(),
		Description: form.GetDesc(),
		Routes:      form.GetRoutes(),
	}
}

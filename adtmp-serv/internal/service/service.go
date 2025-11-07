package service

import (
	"context"

	"adtmp/internal/domain/entities/dto"
	"adtmp/internal/domain/entities/form"
)

// - 认证接口
type AuthService interface {
	AuthUser(ctx context.Context, user *form.UserLogin) (string, error)
	GetUser(ctx context.Context, name string) // 返回值待确认
}

// - 用户接口
type UserService interface {
	CreateUser(ctx context.Context, user *form.UserCreate) error
	DestroyUser(ctx context.Context, id uint) error
	UpdateUser(ctx context.Context, id uint, user *form.UserUpdate) error
	GetById(ctx context.Context, id uint) (*dto.UserResponse, error)
	ListUser(ctx context.Context, limit, offset int) ([]*dto.UserResponse, error)
}

// - 角色接口
type RoleService interface {
	CreateRole(ctx context.Context, role *form.RoleCreate) error
	DestroyRole(ctx context.Context, id uint) error
	UpdateRole(ctx context.Context, id uint, role *form.RoleUpdate) error
	GetById(ctx context.Context, id uint) (*dto.RoleResponse, error)
	ListRole(ctx context.Context, limit, offset int) ([]*dto.RoleResponse, error)
}

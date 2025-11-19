package service

import (
	"context"

	"adtmp/pkg/internal/domain/entities"
	"adtmp/pkg/internal/domain/entities/dto"
	"adtmp/pkg/internal/domain/entities/form"
	"adtmp/pkg/internal/domain/entities/mapper"
	"adtmp/pkg/internal/domain/repositories"
)

type roleService struct {
	roleRepo repositories.RoleRepo
}

func NewRoleService(roleRepo repositories.RoleRepo) *roleService {
	return &roleService{roleRepo: roleRepo}
}

func (roleService *roleService) CreateRole(ctx context.Context, role *form.RoleCreate) error {
	return roleService.roleRepo.Create(ctx, form.ToRoleDb(role))
}

func (roleService *roleService) DestroyRole(ctx context.Context, id uint) error {
	return roleService.roleRepo.Destroy(ctx, id)
}

func (roleService *roleService) UpdateRole(ctx context.Context, id uint, role *form.RoleUpdate) error {
	return roleService.roleRepo.Update(ctx, id, form.ToRoleDb(role))
}

func (roleService *roleService) GetById(ctx context.Context, id uint) (*dto.RoleResponse, error) {
	roleRepo, err := roleService.roleRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.ToDtoResp(&roleRepo).(*dto.RoleResponse), nil
}

func (roleService *roleService) ListRole(ctx context.Context, limit, offset int) ([]*dto.RoleResponse, error) {
	rolesRepo, err := roleService.roleRepo.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	roleDtos := mapper.ToDtoListResp(rolesRepo, func(e entities.Role) *dto.RoleResponse {
		return &dto.RoleResponse{
			ID: e.ID, Name: e.Name, Description: e.Description,
		}
	})
	return roleDtos, nil
}

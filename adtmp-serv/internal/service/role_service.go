package service

import (
	"context"

	"adtmp/internal/domain/entities/dto"
	"adtmp/internal/domain/entities/form"
	"adtmp/internal/domain/entities/mapper"
	"adtmp/internal/domain/repositories"
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

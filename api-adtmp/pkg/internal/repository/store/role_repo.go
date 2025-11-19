package store

import (
	"adtmp/pkg/internal/domain/entities"
	"adtmp/pkg/internal/domain/repositories"

	"gorm.io/gorm"
)

type roleRepo struct {
	db *gorm.DB
	*baseRepo[entities.Role]
}

func NewRoleRepo(db *gorm.DB) repositories.RoleRepo {
	return &roleRepo{db: db, baseRepo: newBaseRepo[entities.Role](db)}
}

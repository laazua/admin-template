// 泛型实现

package store

import (
	"adtmp/internal/domain/entities"
	"adtmp/internal/domain/repositories"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
	*baseRepo[entities.User]
}

func NewUserRepo(db *gorm.DB) repositories.UserRepo {
	return &userRepo{db: db, baseRepo: newBaseRepo[entities.User](db)}
}

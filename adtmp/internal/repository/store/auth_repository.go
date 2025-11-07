// 实现 internal/repositories 中的接口
package store

import (
	"context"

	"adtmp/internal/domain/entities"
	"adtmp/internal/domain/entities/form"
	"adtmp/internal/domain/repositories"

	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) repositories.AuthRepository {
	return &authRepository{db: db}
}

func (authRepository *authRepository) Auth(ctx context.Context, req *form.UserLogin) (entities.User, error) {
	return gorm.G[entities.User](authRepository.db).Where("email = ?", req.Email).First(ctx)
}

func (authRepository *authRepository) GetUserInfo(ctx context.Context, name string) {}

package store

import (
	"context"

	"adtmp/pkg/internal/domain/entities"
	"adtmp/pkg/internal/domain/entities/form"
	"adtmp/pkg/internal/domain/entities/mapper"
	"adtmp/pkg/internal/domain/repositories"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &userRepository{db: db}
}

func (userRepository *userRepository) Create(ctx context.Context, user *form.UserCreate) (*entities.User, error) {
	dbUser := mapper.ToUserDb(user)
	err := gorm.G[entities.User](userRepository.db).Create(ctx, dbUser)
	if err != nil {
		return nil, err
	}
	return dbUser, nil
}
func (userRepository *userRepository) Destroy(ctx context.Context, id uint) error {
	_, err := gorm.G[entities.User](userRepository.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}
func (userRepository *userRepository) Update(ctx context.Context, id uint, user *form.UserUpdate) error {
	_, err := gorm.G[entities.User](userRepository.db).Where("id = ?", id).Updates(ctx, entities.User{Name: user.Name, Email: user.Email, Password: user.Password, Phone: user.Phone, Avatar: user.Avatar, Roles: user.Roles})
	if err != nil {
		return err
	}
	return nil
}
func (userRepository *userRepository) GetById(ctx context.Context, id uint) (*entities.User, error) {
	repoUser, err := gorm.G[entities.User](userRepository.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}
	return &repoUser, nil
}

func (userRepository *userRepository) List(ctx context.Context, limit, offset int) ([]*entities.User, error) {

	return nil, nil
}

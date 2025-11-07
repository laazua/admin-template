package repositories

import (
	"context"

	"adtmp/internal/domain/entities"
	"adtmp/internal/domain/entities/form"
)

type UserRepository interface {
	Create(ctx context.Context, user *form.UserCreate) (*entities.User, error)
	Destroy(ctx context.Context, id uint) error
	Update(ctx context.Context, id uint, user *form.UserUpdate) error
	GetById(ctx context.Context, id uint) (*entities.User, error)
	List(ctx context.Context, limit, offset int) ([]*entities.User, error)
}

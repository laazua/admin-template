package repositories

import (
	"context"

	"adtmp/pkg/internal/domain/entities"
)

type Entity interface {
	entities.User | entities.Role | entities.Route
}

type Repo[E Entity] interface {
	Create(ctx context.Context, c *E) error
	Destroy(ctx context.Context, id uint) error
	Update(ctx context.Context, id uint, u *E) error
	GetById(ctx context.Context, id uint) (E, error)
	List(ctx context.Context, limit, offset int) ([]E, error)
}

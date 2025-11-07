package repositories

import (
	"context"

	"adtmp/internal/domain/entities"
	"adtmp/internal/domain/entities/form"
)

type AuthRepository interface {
	Auth(ctx context.Context, user *form.UserLogin) (entities.User, error)
	GetUserInfo(ctx context.Context, name string) // 返回值待确认
}

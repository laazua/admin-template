package repositories

import (
	"adtmp/pkg/internal/domain/entities"
)

type UserRepo interface {
	Repo[entities.User]
}

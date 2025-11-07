package repositories

import (
	"adtmp/internal/domain/entities"
)

type UserRepo interface {
	Repo[entities.User]
}

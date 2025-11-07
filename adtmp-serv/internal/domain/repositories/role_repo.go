package repositories

import (
	"adtmp/internal/domain/entities"
)

type RoleRepo interface {
	Repo[entities.Role]
}

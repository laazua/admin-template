package repositories

import (
	"adtmp/pkg/internal/domain/entities"
)

type RoleRepo interface {
	Repo[entities.Role]
}

package mapper

import (
	"adtmp/internal/domain/entities"
	"adtmp/internal/domain/entities/dto"
)

func toRoleResponse(r *entities.Role) *dto.RoleResponse {
	if r == nil {
		return nil
	}
	return &dto.RoleResponse{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}

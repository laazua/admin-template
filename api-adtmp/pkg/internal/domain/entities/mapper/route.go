package mapper

import (
	"adtmp/pkg/internal/domain/entities"
	"adtmp/pkg/internal/domain/entities/dto"
)

func toRouteResponse(r *entities.Route) *dto.RouteResponse {
	if r == nil {
		return nil
	}

	return &dto.RouteResponse{
		Id:   r.ID,
		Name: r.Name,
	}
}

package mapper

import (
	"adtmp/internal/domain/entities"
	"adtmp/internal/domain/entities/dto"
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

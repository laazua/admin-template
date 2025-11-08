package mapper

import (
	"adtmp/internal/domain/entities"
)

// -----------------------------------------------------------------------

type Entity interface {
	entities.User | entities.Role | entities.Route
}

func ToDtoResp[E Entity](e *E) any {
	if e == nil {
		return nil
	}

	switch v := any(e).(type) {
	case *entities.User:
		return toUserResponse(v)
	case *entities.Role:
		return toRoleResponse(v)
	case *entities.Route:
		return toRouteResponse(v)
	default:
		return nil
	}
}

func ToDtoListResp[E any, D any](es []E, mapFunc func(E) *D) []*D {
	if len(es) == 0 {
		return []*D{}
	}
	res := make([]*D, len(es))
	for i, e := range es {
		res[i] = mapFunc(e)
	}
	return res
}

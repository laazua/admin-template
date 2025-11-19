package mapper

import (
	"adtmp/pkg/internal/domain/entities"
	"adtmp/pkg/internal/domain/entities/dto"
	"adtmp/pkg/internal/domain/entities/form"
)

// ToUserResponse 将User实体转换为UserResponse DTO
func toUserResponse(u *entities.User) *dto.UserResponse {
	if u == nil {
		return nil
	}

	return &dto.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// ToUserResponseSlice 将User实体切片转换为UserResponse切片
func ToUserResponseSlice(users []entities.User) []*dto.UserResponse {
	if users == nil {
		return nil
	}

	responses := make([]*dto.UserResponse, len(users))
	for i, user := range users {
		responses[i] = toUserResponse(&user)
	}
	return responses
}

func ToUserDb(u *form.UserCreate) *entities.User {
	if u == nil {
		return nil
	}
	return &entities.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Avatar:   u.Avatar,
		Roles:    u.Roles,
	}
}

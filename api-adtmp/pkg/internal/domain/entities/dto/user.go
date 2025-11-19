package dto

import "time"

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserListResponse 用户列表响应
type UserListResponse struct {
	Users  []*UserResponse `json:"users"`
	Total  uint            `json:"total"`
	Limit  int             `json:"limit"`
	Offset int             `json:"offset"`
}

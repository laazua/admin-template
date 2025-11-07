package dto

import "time"

type RoleResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RoleListResponse struct {
	Roles  []*RoleResponse `json:"roles"`
	Total  uint            `json:"total"`
	Limit  int             `json:"limit"`
	Offset int             `json:"offset"`
}

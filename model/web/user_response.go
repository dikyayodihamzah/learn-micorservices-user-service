package web

import (
	"time"
)

type UserRoleResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserResponse struct {
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	Username  string           `json:"username"`
	Email     string           `json:"email"`
	Phone     string           `json:"phone"`
	Role      UserRoleResponse `json:"role"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

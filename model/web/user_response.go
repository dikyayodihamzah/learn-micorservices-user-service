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

// func NewUserResponse(user domain.User) UserResponse {
// 	return UserResponse{
// 		ID:       user.ID,
// 		Name:     user.Name,
// 		Username: user.Username,
// 		Email:    user.Email,
// 		Phone:    user.Phone,
// 		Role: UserRoleResponse{
// 			ID:   user.RoleID,
// 			Name: user.RoleName,
// 		},
// 		CreatedAt: user.CreatedAt,
// 		UpdatedAt: user.UpdatedAt,
// 	}
// }

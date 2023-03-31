package helper

import (
	"gitlab.com/learn-micorservices/user-service/model/domain"
	"gitlab.com/learn-micorservices/user-service/model/web"
)

// User Responses
func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
		Role: web.UserRoleResponse{
			ID:   user.Role.ID,
			Name: user.Role.Name,
		},
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

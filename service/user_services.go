package service

import (
	"context"
	"time"

	"github.com/go-playground/validator"
	"gitlab.com/learn-micorservices/user-service/exception"
	"gitlab.com/learn-micorservices/user-service/helper"
	"gitlab.com/learn-micorservices/user-service/model/domain"
	"gitlab.com/learn-micorservices/user-service/model/web"
	"gitlab.com/learn-micorservices/user-service/repository"
)

type UserService interface {
	CreateUser(c context.Context, claims helper.JWTClaims, request web.CreateUserRequest) (web.UserResponse, error)
}

type userService struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &userService{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (service *userService) CreateUser(c context.Context, claims helper.JWTClaims, request web.CreateUserRequest) (web.UserResponse, error) {
	if err := service.Validate.Struct(request); err != nil {
		return web.UserResponse{}, exception.ErrBadRequest(err.Error())
	}

	user := domain.User{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Phone:    request.Phone,
		Role: domain.Role{
			ID: request.RoleID,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user.GenerateID()
	user.SetPassword(request.Password)

	err := service.UserRepository.CreateUser(c, user)
	if err != nil {
		return web.UserResponse{}, err
	}

	user, err = service.UserRepository.GetUserByID(c, user.ID)
	return helper.ToUserResponse(user), nil
}

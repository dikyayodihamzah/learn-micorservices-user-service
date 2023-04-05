package service

import (
	"context"
	"log"
	"time"

	"github.com/go-playground/validator"
	"gitlab.com/learn-micorservices/user-service/exception"
	"gitlab.com/learn-micorservices/user-service/helper"
	"gitlab.com/learn-micorservices/user-service/model/domain"
	"gitlab.com/learn-micorservices/user-service/model/web"
	"gitlab.com/learn-micorservices/user-service/repository"
)

type UserService interface {
	CreateUser(c context.Context, request web.CreateUserRequest) (web.UserResponse, error)
	GetAllUser(c context.Context) ([]web.UserResponse, error)
	GetUserByID(c context.Context, id string) (web.UserResponse, error)
	UpdateUser(c context.Context, id string, request web.UpdateUserRequest) (web.UserResponse, error)
	DeleteUser(c context.Context, id string) error
}

type userService struct {
	UserRepository repository.UserRepository
	RoleRepository repository.RoleRepository
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, roleRepository repository.RoleRepository, validate *validator.Validate) UserService {
	return &userService{
		UserRepository: userRepository,
		RoleRepository: roleRepository,
		Validate:       validate,
	}
}

func (service *userService) CreateUser(c context.Context, request web.CreateUserRequest) (web.UserResponse, error) {
	if err := service.Validate.Struct(request); err != nil {
		return web.UserResponse{}, exception.ErrBadRequest(err.Error())
	}

	if userByUsername, _ := service.UserRepository.GetUsersByQuery(c, "username", request.Username); userByUsername.ID != "" {
		return web.UserResponse{}, exception.ErrBadRequest("username already registered")
	}

	if userByEmail, _ := service.UserRepository.GetUsersByQuery(c, "email", request.Email); userByEmail.ID != "" {
		return web.UserResponse{}, exception.ErrBadRequest("email already registered")
	}

	if request.Phone != "" {
		if !helper.IsNumeric(request.Phone) {
			return web.UserResponse{}, exception.ErrBadRequest("Phone should numeric")
		}

		if len([]rune(request.Phone)) < 10 || len([]rune(request.Phone)) > 13 {
			return web.UserResponse{}, exception.ErrBadRequest("Phone should 10-13 digit")
		}

		if userByPhone, _ := service.UserRepository.GetUsersByQuery(c, "phone", request.Phone); userByPhone.ID != "" {
			return web.UserResponse{}, exception.ErrBadRequest("phone already registered")
		}
	}

	if role := service.RoleRepository.GetRoleByID(c, request.RoleID); role.ID == "" {
		return web.UserResponse{}, exception.ErrBadRequest("role not found")
	}

	user := domain.User{
		Name:      request.Name,
		Username:  request.Username,
		Email:     request.Email,
		Password:  request.Password,
		Phone:     request.Phone,
		RoleID:    request.RoleID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user.GenerateID()
	user.SetPassword(request.Password)

	err := service.UserRepository.CreateUser(c, user)
	if err != nil {
		return web.UserResponse{}, err
	}

	log.Println(user.ID)
	user, err = service.UserRepository.GetUserByID(c, user.ID)
	if err != nil {
		return web.UserResponse{}, err
	}

	// KAFKA

	return helper.ToUserResponse(user), nil
}

func (service *userService) GetAllUser(c context.Context) ([]web.UserResponse, error) {
	users, err := service.UserRepository.GetAllUser(c)
	if err != nil {
		return []web.UserResponse{}, err
	}
	// log.Println(users)
	return helper.ToAllUserResponses(users), nil
	// return users, nil
}

func (service *userService) GetUserByID(c context.Context, id string) (web.UserResponse, error) {
	user, err := service.UserRepository.GetUserByID(c, id)
	if err != nil {
		return web.UserResponse{}, err
	}

	if id == "" {
		return web.UserResponse{}, exception.ErrNotFound("user not found")
	}
	return helper.ToUserResponse(user), nil
}

func (service *userService) UpdateUser(c context.Context, id string, request web.UpdateUserRequest) (web.UserResponse, error) {
	if err := service.Validate.Struct(request); err != nil {
		return web.UserResponse{}, exception.ErrBadRequest(err.Error())
	}

	user, err := service.UserRepository.GetUserByID(c, id)
	if err != nil {
		return web.UserResponse{}, exception.ErrNotFound(err.Error())
	}

	if request.Email != "" {
		if userByEmail, _ := service.UserRepository.GetUsersByQuery(c, "email", request.Email); userByEmail.ID != "" && userByEmail.ID != id {
			exception.ErrBadRequest("email already registered")
		}
		user.Email = request.Email
	}

	if request.Username != "" {
		if userByUsername, _ := service.UserRepository.GetUsersByQuery(c, "username", request.Username); userByUsername.ID != "" && userByUsername.ID != id {
			exception.ErrBadRequest("username already registered")
		}
		user.Username = request.Username
	}

	if request.Phone != "" {
		if !helper.IsNumeric(request.Phone) {
			panic(exception.ErrBadRequest("Phone should numeric"))
		}

		if len([]rune(request.Phone)) < 10 || len([]rune(request.Phone)) > 13 {
			panic(exception.ErrBadRequest("Phone should 10-13 digit"))
		}

		if userByPhone, _ := service.UserRepository.GetUsersByQuery(c, "phone", request.Phone); userByPhone.ID != "" && userByPhone.ID != id {
			exception.ErrBadRequest("phone already registered")
		}
		user.Phone = request.Phone
	}

	if request.RoleID != "" {
		if role := service.RoleRepository.GetRoleByID(c, request.RoleID); role.ID == "" {
			exception.ErrBadRequest("role not found")
		}
		user.RoleID = request.RoleID
	}

	user.UpdatedAt = time.Now()

	if err := service.UserRepository.UpdateUser(c, user); err != nil {
		return web.UserResponse{}, exception.ErrInternalServer(err.Error())
	}

	// KAFAK

	user, _ = service.UserRepository.GetUserByID(c, id)
	return helper.ToUserResponse(user), nil
}

func (service *userService) DeleteUser(c context.Context, id string) error {
	user, err := service.UserRepository.GetUserByID(c, id)
	if err != nil || user.ID == "" {
		return exception.ErrNotFound("user not found")
	}

	err = service.UserRepository.DeleteUser(c, id)
	if err != nil {
		return err
	}

	// KAFKA

	return nil
}

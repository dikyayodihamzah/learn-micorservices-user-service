package service

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"gitlab.com/learn-micorservices/user-service/model/domain"
	"gitlab.com/learn-micorservices/user-service/repository"
)

type KafkaUserConsumerService interface {
	Insert(message []byte) error
	Update(message []byte) error
	Delete(message []byte) error
}

type kafkaUserConsumerService struct {
	userRepository repository.UserRepository
}

func NewKafkaUserConsumerService(userRepository repository.UserRepository) KafkaUserConsumerService {
	return &kafkaUserConsumerService{
		userRepository: userRepository,
	}
}

func (service *kafkaUserConsumerService) Insert(message []byte) error {
	var userDTO map[string]interface{}
	
	if err := json.Unmarshal(message, &userDTO); err != nil {
		log.Println("error kafka insert user consumer:", err.Error())
	}
	
	id := userDTO["id"]
	name := userDTO["name"]
	username := userDTO["username"]
	email := userDTO["email"]
	password := userDTO["password"]
	phone := userDTO["phone"]
	role_id := userDTO["role_id"]
	createdat := userDTO["created_at"]
	updatedat := userDTO["updated_at"]
	created_at, _ := time.Parse(time.RFC3339, createdat.(string))
	update_at, _ := time.Parse(time.RFC3339, updatedat.(string))
	
	user := domain.User{
		ID : id.(string),
		Name : name.(string),
		Username : username.(string),
		Email : email.(string),
		Password : password.(string),
		Phone : phone.(string),
		RoleID : role_id.(string),
		CreatedAt : created_at,
		UpdatedAt : update_at,
	}

	if err := service.userRepository.CreateUser(context.Background(), user); err != nil {
		log.Println("error kafka insert user consumer:", err.Error())
		return err
	}

	return nil
}

func (service *kafkaUserConsumerService) Update(message []byte) error {
	var userDTO map[string]interface{}

	if err := json.Unmarshal(message, &userDTO); err != nil {
		log.Println("error kafka update user consumer:", err.Error())
	}
	
	id := userDTO["id"]
	name := userDTO["name"]
	username := userDTO["username"]
	email := userDTO["email"]
	password := userDTO["password"]
	phone := userDTO["phone"]
	role_id := userDTO["role_id"]
	createdat := userDTO["created_at"]
	updatedat := userDTO["updated_at"]
	created_at, _ := time.Parse(time.RFC3339, createdat.(string))
	update_at, _ := time.Parse(time.RFC3339, updatedat.(string))
	
	user := domain.User{
		ID : id.(string),
		Name : name.(string),
		Username : username.(string),
		Email : email.(string),
		Password : password.(string),
		Phone : phone.(string),
		RoleID : role_id.(string),
		CreatedAt : created_at,
		UpdatedAt : update_at,
	}

	if err := service.userRepository.UpdateUser(context.Background(), user); err != nil {
		log.Println("error kafka update user consumer:", err.Error())
		return err
	}

	return nil
}

func (service *kafkaUserConsumerService) Delete(message []byte) error {
	var userDTO map[string]interface{}

	if err := json.Unmarshal(message, &userDTO); err != nil {
		log.Println("error kafka update user consumer:", err.Error())
	}

	user := domain.User{
		ID: userDTO["id"].(string),
	}
	
	if err := service.userRepository.DeleteUser(context.Background(), user.ID); err != nil {
		log.Println("error kafka update user consumer:", err.Error())
		return err
	}

	return nil
}
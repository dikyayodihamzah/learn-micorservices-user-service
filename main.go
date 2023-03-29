package main

import (
	"log"
	"time"

	"github.com/go-playground/validator"
	"gitlab.com/learn-micorservices/user-service/config"
)

func controller() {
	time.Local = time.UTC
	db := config.NewDB
	validate := validator.New()

	// userRepository :
}

func main() {
	go controller()
	log.Println("test")
}

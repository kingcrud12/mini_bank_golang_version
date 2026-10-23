package main

import (
	"learn_go_with_api/internal/adapter/security"
	"learn_go_with_api/internal/adapter/web/in/controllers"
	"learn_go_with_api/internal/adapter/web/out/persistance"
	"learn_go_with_api/internal/adapter/web/out/persistance/config"
	"learn_go_with_api/internal/application/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.Connect()
	if err != nil {
		log.Fatal(err)
	}

	customerRepository := persistance.NewCustomerAdapter(db)
	passwordHasher := security.NewBcryptPasswordHasher()
	registerService := service.NewRegisterService(customerRepository, passwordHasher)
	authController := controllers.NewAuthController(registerService)

	router := gin.Default()
	router.POST("/register", authController.RegisterController)

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}

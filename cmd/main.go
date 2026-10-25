package main

import (
	"learn_go_with_api/internal/adapter/security"
	webconfig "learn_go_with_api/internal/adapter/web/config"
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

	jwtConfig := webconfig.LoadJWTConfig()

	customerRepository := persistance.NewCustomerAdapter(db)
	passwordHasher := security.NewBcryptPasswordHasher()
	tokenGenerator := security.NewJWTTokenGenerator(jwtConfig)

	registerService := service.NewRegisterService(customerRepository, passwordHasher)
	loginService := service.NewLoginService(customerRepository, passwordHasher, tokenGenerator)

	authController := controllers.NewAuthController(registerService, loginService, jwtConfig)

	router := gin.Default()
	router.POST("/register", authController.RegisterController)
	router.POST("/login", authController.LoginController)

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}

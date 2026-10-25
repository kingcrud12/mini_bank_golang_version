package controllers

import (
	"errors"
	"learn_go_with_api/internal/adapter/web/config"
	"learn_go_with_api/internal/adapter/web/in/dto"
	"learn_go_with_api/internal/application/port/in"
	"learn_go_with_api/internal/domain/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthcController struct {
	register  in.CustomerRegistered
	login     in.CustomerLogin
	jwtConfig config.JWTConfig
}

func NewAuthController(register in.CustomerRegistered, login in.CustomerLogin, jwtConfig config.JWTConfig) *AuthcController {
	return &AuthcController{register: register, login: login, jwtConfig: jwtConfig}

}

func (c *AuthcController) RegisterController(ctx *gin.Context) {
	var userInput dto.CustomerRegigestRequestDTO

	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	_, err := c.register.Register(userInput)

	if err != nil {
		if errors.Is(err, model.ErrEmailAlreadyUsed) {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, model.ErrInvalidEmailFormat) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CustomerRegisterResponseDTO{
		FirstName: userInput.FirstName,
		LastName:  userInput.LastName,
		Email:     userInput.Email,
	})

}

func (c *AuthcController) LoginController(ctx *gin.Context) {
	var credentials dto.CustomerLoginRequestDTO

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.login.Login(credentials)
	if err != nil {
		if errors.Is(err, model.ErrInvalidCredentials) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(
		c.jwtConfig.CookieName,
		result.Token,
		int(c.jwtConfig.Expiration.Seconds()),
		"/",
		"",
		c.jwtConfig.Secure,
		true,
	)

	ctx.JSON(http.StatusOK, dto.CustomerLoginResponseDTO{
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Email:     result.Email,
	})
}

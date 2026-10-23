package controllers

import (
	"errors"
	"learn_go_with_api/internal/adapter/web/in/dto"
	"learn_go_with_api/internal/application/port/in"
	"learn_go_with_api/internal/domain/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthcController struct {
	register in.CustomerRegistered
}

func NewAuthController(register in.CustomerRegistered) *AuthcController {
	return &AuthcController{register: register}

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

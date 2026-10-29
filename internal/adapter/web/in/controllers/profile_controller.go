package controllers

import (
	"errors"
	"net/http"

	"learn_go_with_api/internal/adapter/web/in/dto"
	"learn_go_with_api/internal/adapter/web/in/middleware"
	"learn_go_with_api/internal/application/port/in"
	"learn_go_with_api/internal/domain/model"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	getProfile in.CustomerGetProfile
}

func NewProfileController(getProfile in.CustomerGetProfile) *ProfileController {
	return &ProfileController{getProfile: getProfile}
}

func (c *ProfileController) GetProfileController(ctx *gin.Context) {
	customerID := ctx.GetString(middleware.CustomerIDContextKey)

	profile, err := c.getProfile.GetProfile(customerID)
	if err != nil {
		if errors.Is(err, model.ErrCustomerNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CustomerProfileResponseDTO{
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		Email:     profile.Email,
	})
}

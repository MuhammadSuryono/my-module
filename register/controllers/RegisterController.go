package controllers

import (
	"github.com/MuhammadSuryono1997/framework-okta/base/database"
	"github.com/MuhammadSuryono1997/framework-okta/register/models"
	"github.com/MuhammadSuryono1997/framework-okta/register/services"
	"github.com/MuhammadSuryono1997/framework-okta/utils"
	"github.com/gin-gonic/gin"
)

type RegisterController interface {
	RegisterUser(c *gin.Context) string
}

type registerController struct {
	registerService services.RegisterService
}

func RegisterHandler(registerService services.RegisterService) RegisterController {
	return &registerController{
		registerService: registerService,
	}
}

func (controller *registerController) RegisterUser(c *gin.Context) string {
	var credential *models.TMerchant

	if err := c.ShouldBindJSON(&credential); err != nil {
		return "Error input"
	}

	isUserRegistered := controller.registerService.RegisterUser(credential)
	if isUserRegistered {
		return "Number is registered"
	}

	// generateToken := service.JWTAuthService().GenerateToken(credential)
	database.GetDb().Select(&credential)
	return credential.NoHp

}

type RegisterControllerStatic interface {
	RegisterStatic(c *gin.Context) string
}

type registerControllerStatic struct {
	registerService services.RegisterServiceStatic
}

func RegisterHandlerStatic(registerService services.RegisterServiceStatic) RegisterControllerStatic {
	return &registerControllerStatic{
		registerService: registerService,
	}
}

func (controller *registerControllerStatic) RegisterStatic(ctx *gin.Context) string {
	var credential *models.TMerchant
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "Error input"
	}
	isUserAuthenticated := controller.registerService.RegisterStatic(credential.DeviceId, credential.NoHp)
	if isUserAuthenticated {
		return "Number is registered"
	}
	return utils.MaskedNumber(credential.NoHp)
}

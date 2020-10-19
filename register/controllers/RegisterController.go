package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/MuhammadSuryono1997/framework-okta/base/database"
	db "github.com/MuhammadSuryono1997/framework-okta/base/database"
	"github.com/MuhammadSuryono1997/framework-okta/register/models"
	"github.com/MuhammadSuryono1997/framework-okta/register/services"
	"github.com/MuhammadSuryono1997/framework-okta/utils"
	"github.com/gin-gonic/gin"
)

const URL_OTP = "http://localhost:5005/request-otp"

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
	var merchant []models.TMerchant

	if err := c.ShouldBindJSON(&credential); err != nil {
		return "Error input"
	}

	err := db.GetDb().Where("phone_number = ?", credential.PhoneNumber).First(&merchant)
	if err.RowsAffected > 0 {
		return "Number is registered"
	}

	fmt.Println(merchant)
	fmt.Println("Request OTP ....")
	database.GetDb().Create(&credential)
	RequestOTP(credential.PhoneNumber)

	return credential.PhoneNumber

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
	isUserAuthenticated := controller.registerService.RegisterStatic(credential.DeviceId, credential.PhoneNumber)
	if isUserAuthenticated {
		return "Number is registered"
	}
	return utils.MaskedNumber(credential.PhoneNumber)
}

func RequestOTP(nohp string) (string, error) {

	jsonReq, err := json.Marshal(map[string]interface{}{"phone_number": nohp})
	resp, err := http.NewRequest("POST", URL_OTP, bytes.NewBuffer(jsonReq))
	client := &http.Client{}
	req, err := client.Do(resp)

	if err != nil {
		fmt.Println(string(utils.ColorYellow()), err)
		return "", err
	}
	body, _ := ioutil.ReadAll(req.Body)
	fmt.Println(string(utils.ColorCyan()), string(body))

	return "Success request", nil

}

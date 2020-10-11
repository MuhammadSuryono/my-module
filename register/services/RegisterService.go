package services

import (
	"fmt"

	"github.com/MuhammadSuryono1997/framework-okta/config"
	"github.com/MuhammadSuryono1997/framework-okta/register/models"
)

type RegisterService interface {
	RegisterUser(credential *models.TMerchant) bool
}

type RegisterServiceStatic interface {
	RegisterStatic(devid string, nohp string) bool
}

type registerInformation struct {
	device_id string
	no_hp     string
}

func StaticRegisterService() RegisterServiceStatic {
	return &registerInformation{
		device_id: "123456789",
		no_hp:     "0895355698652",
	}
}

func RegisterUser(credential *models.TMerchant) bool {
	var merchant []models.TMerchant

	err := config.GetDb().Where("no_hp = ? AND device_id = ?", credential.NoHp, credential.DeviceId).First(&merchant)
	if err != nil {
		return false
	}

	fmt.Println(merchant)

	return true
}

func (info *registerInformation) RegisterStatic(devid string, nohp string) bool {
	StaticRegisterService()
	return info.device_id == devid && info.no_hp == nohp
}
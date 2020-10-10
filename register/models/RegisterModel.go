package models

//Login credential
type TMerchant struct {
	DeviceId string `form:"device_id"`
	NoHp     string `form:"no_hp"`
}

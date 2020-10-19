package models

//Login credential
type TMerchant struct {
	DeviceId    string `json:"device_id"`
	PhoneNumber string `json:"phone_number"`
}

type TMerchantSecret struct {
	Secret       string
	RandomString string
	Timestamp    string
}

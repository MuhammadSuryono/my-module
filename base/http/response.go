package httpresponse

type ErrorCode struct {
	Code    int
	Message string
}

//0000: No Error
//1001: User not found
//1002: User no permissions
type BaseResponse struct {
	//Success
	IsSuccess bool `json:"is_success"`

	//Error
	Error ErrorCode   `json:"error"`
	Data  interface{} `json:"data"`
}

func (err ErrorCode) AsInvalidResponse() BaseResponse {

	return BaseResponse{
		IsSuccess: false,
		Error:     err,
		Data:      nil,
	}
}
func (err ErrorCode) AsValidResponse(data interface{}) BaseResponse {

	return BaseResponse{
		IsSuccess: false,
		Error:     err,
		Data:      data,
	}
}

var NOT_FOUND = ErrorCode{
	Code:    0001,
	Message: "",
}

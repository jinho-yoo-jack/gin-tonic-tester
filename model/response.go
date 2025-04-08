package model

type ApiResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  *apiError   `json:"error,omitempty"`
}

type apiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Success(data interface{}) *ApiResponse {
	return &ApiResponse{
		Status: "success",
		Data:   data,
	}
}

func Fail(code int, message string) *ApiResponse {
	return &ApiResponse{
		Status: "fail",
		Error: &apiError{
			Code:    code,
			Message: message,
		},
	}
}

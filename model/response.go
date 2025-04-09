package model

import "fmt"

type ApiResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  *apiError   `json:"error,omitempty"`
}

type apiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"error"`
}

func (e *apiError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s : %v", e.Message, e.Err)
	}
	return e.Message
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

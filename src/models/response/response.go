package response

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

type ErrorBody struct {
	Code    int             `json:"code,omitempty"`
	Message string          `json:"message,omitempty"`
	Error   []*ErrorMessage `json:"errors"`
}

func mapToErrorOutput(e *fiber.Error) *ErrorMessage {
	return &ErrorMessage{
		Message: e.Message,
	}
}

func HttpResponse(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func HttpErrorResponse(errorObj []*fiber.Error) *ErrorBody {
	var errorSlice []*ErrorMessage

	for i := 0; i < len(errorObj); i++ {
		errorSlice = append(errorSlice, mapToErrorOutput(errorObj[i]))
	}

	return &ErrorBody{
		Code:    http.StatusBadRequest,
		Message: "Unprocessable Entity",
		Error:   errorSlice,
	}
}

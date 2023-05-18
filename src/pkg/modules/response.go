package modules

import (
	"net/http"
)

type Response struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Total   int         `json:"total,omitempty"`
}

func HttpResponse(response ...Response) *Response {
	result := &Response{}

	for _, resp := range response {

		result.Code = resp.Code
		result.Message = resp.Message
		result.Data = resp.Data
		result.Total = resp.Total

		if resp.Code == 0 {
			result.Code = http.StatusOK
		}

		if resp.Message == "" {
			result.Message = "data has been received"
		}
	}

	return result
}

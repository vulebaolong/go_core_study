package response

import "encoding/json"

type AppSuccess struct {
	Message string
	Data    any
}

func NewAppSuccess(data any, message string) string {
	if message == "" {
		message = "OK"
	}

	response := &AppSuccess{
		Message: message,
		Data:    data,
	}

	b, _ := json.MarshalIndent(response, "", "  ")

	return string(b)
}

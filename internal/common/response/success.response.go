package response

type AppSuccess struct {
	Status  string
	Message string
	Data    any
}

func NewAppSuccess(data any, message string) *AppSuccess {
	if message == "" {
		message = "OK"
	}

	return &AppSuccess{
		Status:  "Success",
		Message: message,
		Data:    data,
	}
}

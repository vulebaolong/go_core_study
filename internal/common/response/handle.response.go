package response

import (
	"encoding/json"
)

func HandleResponse(data any, err error, message string) string {
	if err != nil {
		if appError, ok := err.(*AppError); ok {
			b, _ := json.MarshalIndent(appError, "", "  ")
			return string(b)
		} else {
			resErr := NewAppError(err.Error())
			b, _ := json.MarshalIndent(resErr, "", "  ")
			return string(b)
		}
	} else {
		response := NewAppSuccess(data, message)
		b, _ := json.MarshalIndent(response, "", "  ")
		return string(b)

	}
}

package helper

import (
	"net/http"

	"github.com/ariefsn/upwork/models"
)

func ResponseJson(w http.ResponseWriter, data models.ResponseModel, statusCode ...int) {
	_statusCode := http.StatusOK

	if len(statusCode) > 0 {
		_statusCode = statusCode[0]
	}

	w.WriteHeader(_statusCode)

	w.Write(ToBytes(data))
}

func ResponseJsonSuccess(w http.ResponseWriter, data interface{}, statusCode ...int) {
	ResponseJson(w, models.ResponseModel{
		Success: true,
		Data:    data,
	}, statusCode...)
}

func ResponseJsonError(w http.ResponseWriter, message string, statusCode ...int) {
	ResponseJson(w, models.ResponseModel{
		Success: false,
		Message: message,
	}, statusCode...)
}

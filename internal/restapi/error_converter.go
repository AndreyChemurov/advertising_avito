package restapi

import (
	"encoding/json"
	"strconv"
)

// errorType оборачивает данные (http статус код + сообщение) об ошибке
// в json и возвращает его вызывающему хэндлеру
func errorType(status int, message string) (response []byte) {
	err := map[string]map[string]string{
		"error": {
			"status_code":    strconv.Itoa(status),
			"status_message": message,
		},
	}

	response, _ = json.Marshal(err)
	return response
}

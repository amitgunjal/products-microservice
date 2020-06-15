package utils

import (
	"encoding/json"
	"net/http"
	"products/constants"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{constants.Status: status, constants.Message: message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add(constants.ContentType, constants.ApplicationJson)
	_ = json.NewEncoder(w).Encode(data)
}

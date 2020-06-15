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

// respondJSON makes the response with payload as json format
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, map[string]string{"error": message})
}

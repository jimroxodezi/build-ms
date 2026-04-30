package handlers

import (
	"encoding/json"
	"net/http"
)

func WriteError(w http.ResponseWriter, statusCode int, errMessage string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	msg := map[string]string{"error": errMessage}
	json.NewEncoder(w).Encode(msg)
}

package utils

import (
	"encoding/json"
	"net/http"
)

func Success(w http.ResponseWriter, status int, data any) {
	writeJSON(w, status, map[string]any{
		"success": true,
		"data":    data,
	})
}

func Error(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]any{
		"success": false,
		"error":   msg,
	})
}

func ReadJSON(r *http.Request, dst any) error {
	return json.NewDecoder(r.Body).Decode(dst)
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

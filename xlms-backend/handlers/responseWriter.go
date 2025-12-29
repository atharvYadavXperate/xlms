package handlers

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, code int, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": message,
		"data":    data,
	})
}

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mostafasaad-m/task-api/internal/middleware"
)

// test passed
func Me(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]any{
		"user_id": userID,
		"message": "authenticated",
	})
}

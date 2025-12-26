package handler

import (
	"encoding/json"
	"net/http"
)

func (h *HandlerUser) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"id":   "1",
		"name": "John Doe",
	})
}

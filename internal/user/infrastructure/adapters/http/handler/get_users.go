package handler

import (
	"encoding/json"
	"net/http"
)

func (h *HandlerUser) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.GetUserUseCase.GetUsers(r.Context())

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	response := []*Response{}

	for _, u := range users {
		response = append(response, &Response{
			ID:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			UserName:  u.UserName,
		})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

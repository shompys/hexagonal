package handler

import (
	"encoding/json"
	"net/http"

	handlerDTO "github.com/shompys/hexagonal/internal/user/infrastructure/adapters/http/handler/dto"
)

func (h *HandlerUser) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.GetUserUseCase.GetUsers()

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response := []*handlerDTO.Response{}

	for _, u := range users {
		response = append(response, &handlerDTO.Response{
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

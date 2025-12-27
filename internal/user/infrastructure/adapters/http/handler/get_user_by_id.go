package handler

import (
	"encoding/json"
	"net/http"

	handlerDTO "github.com/shompys/hexagonal/internal/user/infrastructure/adapters/http/handler/dto"
)

func (h *HandlerUser) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	w.Header().Set("Content-Type", "application/json")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "id is required"})
		return
	}

	userDTO, err := h.GetUserUseCase.GetUserByID(r.Context(), id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&handlerDTO.Response{
		ID:        userDTO.ID,
		FirstName: userDTO.FirstName,
		LastName:  userDTO.LastName,
		Email:     userDTO.Email,
		UserName:  userDTO.UserName,
	})
}

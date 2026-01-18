package handler

import (
	"encoding/json"
	"net/http"
)

func (h *HandlerUser) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	w.Header().Set("Content-Type", "application/json")

	userDTO, err := h.GetUserUseCase.GetUserByID(r.Context(), id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Response{
		ID:        userDTO.ID,
		FirstName: userDTO.FirstName,
		LastName:  userDTO.LastName,
		Email:     userDTO.Email,
		UserName:  userDTO.UserName,
	})
}

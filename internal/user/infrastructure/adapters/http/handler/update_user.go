package handler

import (
	"encoding/json"
	"net/http"

	domainDTO "github.com/shompys/hexagonal/internal/user/domain/dto"
)

func (h *HandlerUser) UpdateUser(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	var user UpdateRequest
	json.NewDecoder(r.Body).Decode(&user)

	userUpdated, err := h.GetUserUseCase.UpdateUser(r.Context(), id, &domainDTO.UserUpdateInput{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		UserName:  user.UserName,
		Password:  user.Password,
	})
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(Response{
		ID:        userUpdated.ID,
		FirstName: userUpdated.FirstName,
		LastName:  userUpdated.LastName,
		Email:     userUpdated.Email,
		UserName:  userUpdated.UserName,
	})
}

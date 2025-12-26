package handler

import (
	"encoding/json"
	"net/http"

	domainDTO "github.com/shompys/hexagonal/internal/user/domain/dto"
	handlerDTO "github.com/shompys/hexagonal/internal/user/infrastructure/adapters/http/handler/dto"
)

func (h *HandlerUser) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user handlerDTO.CreateRequest
	json.NewDecoder(r.Body).Decode(&user)

	userCreated, err := h.GetUserUseCase.CreateUser(&domainDTO.UserCreateInput{
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

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(handlerDTO.Response{
		ID:        userCreated.ID,
		FirstName: userCreated.FirstName,
		LastName:  userCreated.LastName,
		Email:     userCreated.Email,
		UserName:  userCreated.UserName,
	})

}

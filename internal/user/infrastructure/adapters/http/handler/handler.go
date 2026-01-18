package handler

import (
	"github.com/shompys/hexagonal/internal/user/domain/ports"
)

type CreateRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	UserName  string `json:"userName"`
	Password  string `json:"password"`
}

type UpdateRequest struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Email     *string `json:"email"`
	UserName  *string `json:"userName"`
	Password  *string `json:"password"`
}

type Response struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	UserName  string `json:"userName"`
}

type HandlerUser struct {
	GetUserUseCase ports.UserUseCases
}

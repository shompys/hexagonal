package handler

import (
	"github.com/shompys/hexagonal/internal/user/domain/ports"
)

type HandlerUser struct {
	GetUserUseCase ports.UserUseCases
}

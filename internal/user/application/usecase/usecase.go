package usecase

import (
	"github.com/shompys/hexagonal/internal/user/domain"
	"github.com/shompys/hexagonal/internal/user/domain/ports"
)

type UserUseCase struct {
	UserRepository ports.UserRepository
	PasswordHasher domain.PasswordHasher
}

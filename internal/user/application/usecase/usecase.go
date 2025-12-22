package usecase

import "github.com/shompys/hexagonal/internal/user/domain/ports"

type UserUseCase struct {
	userRepository ports.UserRepository
}

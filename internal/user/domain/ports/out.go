package ports

import (
	"github.com/shompys/hexagonal/internal/user/domain"
)

type UserRepository interface {
	Create(userEntity *domain.User) (*domain.User, error)
	GetUserByID(id string) (*domain.User, error)
	GetUsers() ([]*domain.User, error)
	UpdateUser(id string, userEntity *domain.User) (*domain.User, error)
	DeleteUser(id string) error
}

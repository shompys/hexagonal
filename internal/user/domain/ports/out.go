package ports

import (
	"context"

	"github.com/shompys/hexagonal/internal/user/domain"
	"github.com/shompys/hexagonal/internal/user/domain/dto"
)

// es el como se conecta el repository al useCase
type UserRepository interface {
	Create(ctx context.Context, userEntity *domain.User) (*domain.User, error)
	GetUserByID(ctx context.Context, id domain.UserIDVO) (*domain.User, error)
	GetUsers(ctx context.Context, filters dto.Filters) ([]*domain.User, error)
	UpdateUser(ctx context.Context, id domain.UserIDVO, userEntity *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id domain.UserIDVO) error
	DeleteSoftUser(ctx context.Context, id domain.UserIDVO, userEntity *domain.User) error
}

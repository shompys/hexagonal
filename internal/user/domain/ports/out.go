package ports

import (
	"context"

	"github.com/shompys/hexagonal/internal/user/domain"
)

type UserRepository interface {
	Create(ctx context.Context, userEntity *domain.User) (*domain.User, error)
	GetUserByID(ctx context.Context, id domain.UserIDVO) (*domain.User, error)
	GetUsers(ctx context.Context) ([]*domain.User, error)
	UpdateUser(ctx context.Context, id domain.UserIDVO, userEntity *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id domain.UserIDVO) error
}

package ports

import (
	"context"

	"github.com/shompys/hexagonal/internal/user/domain/dto"
)

// Contrato para el handler, es el como se conecta el useCase al handler
type UserUseCases interface {
	CreateUser(ctx context.Context, userDTO *dto.UserCreateInput) (*dto.UserOutput, error)
	GetUserByID(ctx context.Context, id string) (*dto.UserOutput, error)
	GetUsers(ctx context.Context, filters dto.Filters) ([]*dto.UserOutput, error)
	UpdateUser(ctx context.Context, id string, userDTO *dto.UserUpdateInput) (*dto.UserOutput, error)
	DeleteUser(ctx context.Context, id string) error
	DeleteSoftUser(ctx context.Context, id string) error
}

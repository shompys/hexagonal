package ports

import (
	"context"

	"github.com/shompys/hexagonal/internal/user/domain/dto"
)

type UserUseCases interface {
	CreateUser(ctx context.Context, userDTO *dto.UserCreateInput) (*dto.UserOutput, error)
	GetUserByID(ctx context.Context, id string) (*dto.UserOutput, error)
	GetUsers(ctx context.Context) ([]*dto.UserOutput, error)
	UpdateUser(ctx context.Context, id string, userDTO *dto.UserUpdateInput) (*dto.UserOutput, error)
	DeleteUser(ctx context.Context, id string) error
}

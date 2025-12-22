package ports

import "github.com/shompys/hexagonal/internal/user/domain/dto"

type UserUseCases interface {
	CreateUser(userDTO *dto.UserCreateInput) (*dto.UserOutput, error)
	GetUserByID(id string) (*dto.UserOutput, error)
	GetUsers() ([]*dto.UserOutput, error)
	UpdateUser(userDTO *dto.UserUpdateInput) (*dto.UserOutput, error)
	DeleteUser(id string) error
}

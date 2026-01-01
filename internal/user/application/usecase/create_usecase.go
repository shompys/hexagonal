package usecase

import (
	"context"

	"github.com/shompys/hexagonal/internal/user/domain"
	"github.com/shompys/hexagonal/internal/user/domain/dto"
)

func (uc *UserUseCase) CreateUser(ctx context.Context, userDTO *dto.UserCreateInput) (*dto.UserOutput, error) {

	password, err := domain.NewUserPassword(userDTO.Password, uc.PasswordHasher)
	if err != nil {
		return nil, err //TODO: wrap error
	}

	user, err := domain.NewUser(userDTO.FirstName, userDTO.LastName, userDTO.Email, userDTO.UserName, password)
	if err != nil {
		return nil, err //TODO: wrap error
	}

	userEntity, err := uc.UserRepository.Create(ctx, user)
	if err != nil {
		return nil, err //TODO: wrap error
	}

	return &dto.UserOutput{
		ID:        userEntity.ID(),
		FirstName: userEntity.FirstName(),
		LastName:  userEntity.LastName(),
		Email:     userEntity.Email(),
		UserName:  userEntity.UserName(),
	}, nil
}

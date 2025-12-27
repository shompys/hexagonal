package usecase

import (
	"context"

	"github.com/shompys/hexagonal/internal/user/domain/dto"
)

func (uc *UserUseCase) GetUserByID(ctx context.Context, id string) (*dto.UserOutput, error) {

	user, err := uc.UserRepository.GetUserByID(ctx, id)

	if err != nil {
		return nil, err //TODO: wrap error
	}
	return &dto.UserOutput{
		ID:        user.ID(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
		Email:     user.Email(),
		UserName:  user.UserName(),
	}, nil
}

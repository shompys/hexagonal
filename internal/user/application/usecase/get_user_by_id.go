package usecase

import (
	"context"
	"fmt"

	"github.com/shompys/hexagonal/internal/user/domain"
	"github.com/shompys/hexagonal/internal/user/domain/dto"
)

func (uc *UserUseCase) GetUserByID(ctx context.Context, id string) (*dto.UserOutput, error) {

	idVO, err := domain.NewUserID(id)

	if err != nil {
		return nil, err
	}

	user, err := uc.UserRepository.GetUserByID(ctx, idVO)

	if err != nil {
		return nil, err //TODO: wrap error
	}

	if user.Status() != domain.StatusActive {
		return nil, fmt.Errorf("user not found")
	}

	return &dto.UserOutput{
		ID:        user.ID(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
		Email:     user.Email(),
		UserName:  user.UserName(),
	}, nil
}

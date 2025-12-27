package usecase

import (
	"context"

	"github.com/shompys/hexagonal/internal/user/domain/dto"
)

func (uc *UserUseCase) GetUserByID(ctx context.Context, id string) (*dto.UserOutput, error) {
	return &dto.UserOutput{
		ID:        "",
		FirstName: "",
		LastName:  "",
		Email:     "",
		UserName:  "",
	}, nil
}

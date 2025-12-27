package usecase

import (
	"context"

	"github.com/shompys/hexagonal/internal/user/domain/dto"
)

func (uc *UserUseCase) UpdateUser(ctx context.Context, userDTO *dto.UserUpdateInput) (*dto.UserOutput, error) {
	return &dto.UserOutput{
		ID:        "",
		FirstName: "",
		LastName:  "",
		Email:     "",
		UserName:  "",
	}, nil
}

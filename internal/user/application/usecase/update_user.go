package usecase

import "github.com/shompys/hexagonal/internal/user/domain/dto"

func (uc *UserUseCase) UpdateUser(userDTO *dto.UserUpdateInput) (*dto.UserOutput, error) {
	return &dto.UserOutput{
		ID:        "",
		FirstName: "",
		LastName:  "",
		Email:     "",
		UserName:  "",
	}, nil
}

package usecase

import "github.com/shompys/hexagonal/internal/user/domain/dto"

func (uc *UserUseCase) GetUserByID(id string) (*dto.UserOutput, error) {
	return &dto.UserOutput{
		ID:        "",
		FirstName: "",
		LastName:  "",
		Email:     "",
		UserName:  "",
	}, nil
}

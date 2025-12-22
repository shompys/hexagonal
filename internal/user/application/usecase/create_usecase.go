package usecase

import (
	"github.com/shompys/hexagonal/internal/user/domain"
	"github.com/shompys/hexagonal/internal/user/domain/dto"
)

func (uc *UserUseCase) CreateUser(userDTO *dto.UserCreateInput) (*dto.UserOutput, error) {

	userEntity, err := uc.userRepository.Create(&domain.User{
		ID:           "",
		FirstName:    userDTO.FirstName,
		LastName:     userDTO.LastName,
		Email:        userDTO.Email,
		UserName:     userDTO.UserName,
		PasswordHash: userDTO.Password,
	})
	if err != nil {
		return nil, err //TODO: wrap error
	}

	return &dto.UserOutput{
		ID:        userEntity.ID,
		FirstName: userEntity.FirstName,
		LastName:  userEntity.LastName,
		Email:     userEntity.Email,
		UserName:  userEntity.UserName,
		Password:  userEntity.PasswordHash,
	}, nil
}

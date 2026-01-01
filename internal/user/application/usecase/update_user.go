package usecase

import (
	"context"

	"github.com/shompys/hexagonal/internal/user/domain"
	"github.com/shompys/hexagonal/internal/user/domain/dto"
)

func (uc *UserUseCase) UpdateUser(ctx context.Context, id string, userDTO *dto.UserUpdateInput) (*dto.UserOutput, error) {

	idVO, err := domain.NewUserID(id)

	if err != nil {
		return nil, err
	}

	user, err := uc.UserRepository.GetUserByID(ctx, idVO)

	if err != nil {
		return nil, err
	}

	if userDTO.FirstName != nil && *userDTO.FirstName != user.FirstName() {
		if err := user.SetFirstName(*userDTO.FirstName); err != nil {
			return nil, err
		}
	}
	if userDTO.LastName != nil && *userDTO.LastName != user.LastName() {
		if err := user.SetLastName(*userDTO.LastName); err != nil {
			return nil, err
		}
	}
	if userDTO.Email != nil && *userDTO.Email != user.Email() {
		if err := user.SetEmail(*userDTO.Email); err != nil {
			return nil, err
		}
	}
	if userDTO.UserName != nil && *userDTO.UserName != user.UserName() {
		if err := user.SetUserName(*userDTO.UserName); err != nil {
			return nil, err
		}
	}
	//este bloques define si no viene el password y si viene siempre genera uno nuevo no se comparan.
	if userDTO.Password != nil {
		passwordVO, err := domain.NewUserPassword(*userDTO.Password, uc.PasswordHasher)
		if err != nil {
			return nil, err
		}
		user.SetPasswordHash(passwordVO)
	}

	userUpdated, err := uc.UserRepository.UpdateUser(ctx, idVO, user)

	if err != nil {
		return nil, err
	}

	return &dto.UserOutput{
		ID:        userUpdated.ID(),
		FirstName: userUpdated.FirstName(),
		LastName:  userUpdated.LastName(),
		Email:     userUpdated.Email(),
		UserName:  userUpdated.UserName(),
	}, nil
}

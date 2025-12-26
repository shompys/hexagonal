package usecase

import "github.com/shompys/hexagonal/internal/user/domain/dto"

func (uc *UserUseCase) GetUsers() ([]*dto.UserOutput, error) {
	users, err := uc.UserRepository.GetUsers()

	if err != nil {
		return nil, err //TODO: wrap error
	}

	response := []*dto.UserOutput{}

	for _, u := range users {
		response = append(response, &dto.UserOutput{
			ID:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			UserName:  u.UserName,
		})
	}

	return response, nil
}

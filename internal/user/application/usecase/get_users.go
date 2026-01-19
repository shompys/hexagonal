package usecase

import (
	"context"

	"github.com/shompys/hexagonal/internal/user/domain"
	"github.com/shompys/hexagonal/internal/user/domain/dto"
)

func (uc *UserUseCase) GetUsers(ctx context.Context, filters dto.Filters) ([]*dto.UserOutput, error) {

	//siempre es active porque aún no existe un administrador que pueda filtrar
	activeStatus := domain.StatusActive
	filters.Status = &activeStatus

	users, err := uc.UserRepository.GetUsers(ctx, filters)

	if err != nil {
		return nil, err //TODO: wrap error
	}

	response := []*dto.UserOutput{}

	for _, u := range users {
		response = append(response, &dto.UserOutput{
			ID:        u.ID(),
			FirstName: u.FirstName(),
			LastName:  u.LastName(),
			Email:     u.Email(),
			UserName:  u.UserName(),
		})
	}

	return response, nil
}

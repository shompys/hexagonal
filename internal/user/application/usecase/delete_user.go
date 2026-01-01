package usecase

import (
	"context"

	"github.com/shompys/hexagonal/internal/user/domain"
)

func (uc *UserUseCase) DeleteUser(ctx context.Context, id string) error {

	idVO, err := domain.NewUserID(id)
	if err != nil {
		return err
	}
	if err := uc.UserRepository.DeleteUser(ctx, idVO); err != nil {
		return err
	}
	return nil
}

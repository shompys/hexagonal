package usecase

import (
	"context"

	"github.com/shompys/hexagonal/internal/user/domain"
)

func (uc *UserUseCase) DeleteSoftUser(ctx context.Context, id string) error {

	idVO, err := domain.NewUserID(id)
	if err != nil {
		return err
	}
	user, err := uc.UserRepository.GetUserByID(ctx, idVO)

	if err != nil {
		return err
	}

	if err := user.Deactivate(); err != nil {
		// si esto falla listo retornamos se cancela el cambio, porque ya se encuentra desactivado
		return err
	}

	if err := uc.UserRepository.DeleteSoftUser(ctx, idVO, user); err != nil {
		return err
	}
	return nil
}

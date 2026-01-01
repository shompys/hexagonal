package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/shompys/hexagonal/internal/user/domain"
)

type MemoryUserRepository struct {
	users []*domain.User
}

func (r *MemoryUserRepository) Create(ctx context.Context, userEntity *domain.User) (*domain.User, error) {
	newID := fmt.Sprintf("%d", len(r.users)+1)

	id, err := domain.NewUserID(newID)

	if err != nil {
		return nil, err
	}

	userEntity.SetID(id)

	r.users = append(r.users, userEntity)
	return userEntity, nil
}

func (r *MemoryUserRepository) GetUserByID(ctx context.Context, id domain.UserIDVO) (*domain.User, error) {

	for _, user := range r.users {

		if user.ID() == id.Value() {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *MemoryUserRepository) GetUsers(ctx context.Context) ([]*domain.User, error) {
	result := []*domain.User{}

	for i := range r.users {
		result = append(result, r.users[i])
	}
	return result, nil
}

func (r *MemoryUserRepository) UpdateUser(ctx context.Context, id domain.UserIDVO, userEntity *domain.User) (*domain.User, error) {
	for i, user := range r.users {
		if user.ID() == id.Value() {
			r.users[i] = userEntity
			return userEntity, nil
		}
	}
	return nil, errors.New("user not found")
}
func (r *MemoryUserRepository) DeleteUser(ctx context.Context, id domain.UserIDVO) error {
	for i, user := range r.users {
		if user.ID() == id.Value() {
			//[123, 4, 56] aca esta pasando que para eliminar el elemento los saltea
			// [123, 56]
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

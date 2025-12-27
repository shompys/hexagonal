package repository

import (
	"context"
	"errors"

	"github.com/shompys/hexagonal/internal/user/domain"
)

type MemoryUserRepository struct {
	users []domain.User
}

func (r *MemoryUserRepository) Create(ctx context.Context, userEntity *domain.User) (*domain.User, error) {
	r.users = append(r.users, *userEntity)
	return userEntity, nil
}
func (r *MemoryUserRepository) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	for _, user := range r.users {

		if user.ID() == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}
func (r *MemoryUserRepository) GetUsers(ctx context.Context) ([]*domain.User, error) {
	result := []*domain.User{}

	for i := range r.users {
		result = append(result, &r.users[i])
	}
	return result, nil
}

func (r *MemoryUserRepository) UpdateUser(ctx context.Context, id string, userEntity *domain.User) (*domain.User, error) {
	for i, user := range r.users {
		if user.ID() == id {
			r.users[i] = *userEntity
			return userEntity, nil
		}
	}
	return nil, errors.New("user not found")
}
func (r *MemoryUserRepository) DeleteUser(ctx context.Context, id string) error {
	for i, user := range r.users {
		if user.ID() == id {
			//[123, 4, 56] aca esta pasando que para eliminar el elemento los saltea
			// [123, 56]
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

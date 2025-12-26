package repository

import (
	"errors"

	"github.com/shompys/hexagonal/internal/user/domain"
)

type MemoryUserRepository struct {
	users []domain.User
}

func (r *MemoryUserRepository) Create(userEntity *domain.User) (*domain.User, error) {
	r.users = append(r.users, *userEntity)
	return userEntity, nil
}
func (r *MemoryUserRepository) GetUserByID(id string) (*domain.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}
func (r *MemoryUserRepository) GetUsers() ([]*domain.User, error) {
	result := []*domain.User{}

	for i := range r.users {
		result = append(result, &r.users[i])
	}
	return result, nil
}

func (r *MemoryUserRepository) UpdateUser(id string, userEntity *domain.User) (*domain.User, error) {
	for i, user := range r.users {
		if user.ID == id {
			r.users[i] = *userEntity
			return userEntity, nil
		}
	}
	return nil, errors.New("user not found")
}
func (r *MemoryUserRepository) DeleteUser(id string) error {
	for i, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

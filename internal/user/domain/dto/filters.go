package dto

import "github.com/shompys/hexagonal/internal/user/domain"

type Filters struct {
	Status *domain.UserStatus
}

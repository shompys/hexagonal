package bootstrap

import (
	"github.com/shompys/hexagonal/internal/user/application/usecase"
	"github.com/shompys/hexagonal/internal/user/infrastructure/repository"
	"github.com/shompys/hexagonal/pkg/hash"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// UserDependencies contiene todas las dependencias para la feature de usuarios
type UserDependencies struct {
	Repository *repository.MongoUserRepository
	// Repository *repository.MemoryUserRepository
	PasswordHasher hash.PasswordHasher
	UseCase        *usecase.UserUseCase
}

// InitializeUser inicializa todas las dependencias de la feature de usuarios
func InitializeUser(db *mongo.Database) *UserDependencies {
	// Infraestructura
	userRepo := repository.NewMongoUserRepository(db)
	// userRepo := &repository.MemoryUserRepository{}
	passwordHasher := hash.PasswordHasher{}

	// Use cases
	userUC := &usecase.UserUseCase{
		UserRepository: userRepo,
		PasswordHasher: passwordHasher,
	}

	return &UserDependencies{
		Repository:     userRepo,
		PasswordHasher: passwordHasher,
		UseCase:        userUC,
	}
}

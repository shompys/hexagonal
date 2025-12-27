package main

import (
	"fmt"

	"log"
	"net/http"

	"github.com/shompys/hexagonal/internal/user/application/usecase"
	httpAdapter "github.com/shompys/hexagonal/internal/user/infrastructure/adapters/http"
	"github.com/shompys/hexagonal/internal/user/infrastructure/adapters/http/handler"
	"github.com/shompys/hexagonal/internal/user/infrastructure/repository"
	"github.com/shompys/hexagonal/pkg/hash"
)

func main() {

	userRepo := &repository.MemoryUserRepository{}

	passwordHasher := hash.PasswordHasher{}

	userUC := &usecase.UserUseCase{
		UserRepository: userRepo,
		PasswordHasher: passwordHasher,
	}
	userHandler := &handler.HandlerUser{
		GetUserUseCase: userUC,
	}

	router := httpAdapter.NewRouter(userHandler)

	router.RegisterRoutes()

	root := http.NewServeMux()
	root.Handle("/v1/", http.StripPrefix("/v1", router.Handler()))

	fmt.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", root); err != nil {
		log.Fatal("Failed to start server:", err)
	}

}

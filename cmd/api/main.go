package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shompys/hexagonal/internal/bootstrap"
	httpAdapter "github.com/shompys/hexagonal/internal/user/infrastructure/adapters/http"
	"github.com/shompys/hexagonal/internal/user/infrastructure/adapters/http/handler"
)

func main() {
	// Inicializar dependencias de usuario
	userDeps := bootstrap.InitializeUser()

	// Crear handler HTTP
	userHandler := &handler.HandlerUser{
		GetUserUseCase: userDeps.UseCase,
	}

	// Configurar router
	router := httpAdapter.NewRouter(userHandler)
	router.RegisterRoutes()

	// Configurar servidor
	root := http.NewServeMux()
	root.Handle("/v1/", http.StripPrefix("/v1", router.Handler()))

	// Iniciar servidor
	fmt.Println("🚀 Server started on port 8080")
	fmt.Println("📡 API available at http://localhost:8080/v1")
	if err := http.ListenAndServe(":8080", root); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

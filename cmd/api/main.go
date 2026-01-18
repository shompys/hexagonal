package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/shompys/hexagonal/internal/bootstrap"
	"github.com/shompys/hexagonal/internal/storage"
	httpAdapter "github.com/shompys/hexagonal/internal/user/infrastructure/adapters/http"
	"github.com/shompys/hexagonal/internal/user/infrastructure/adapters/http/handler"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	cfg := storage.Config{
		User:     os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
		Host:     os.Getenv("MONGO_HOST"),
		Port:     os.Getenv("MONGO_PORT"),
		Database: os.Getenv("MONGO_DB_NAME"),
	}

	db, client, err := storage.NewClient(context.Background(), cfg)

	if err != nil {
		log.Fatal(err) //log se ocupa de ejecutar el erro.Error()
	}

	defer client.Disconnect(context.Background())

	// Inicializar dependencias de usuario
	userDeps := bootstrap.InitializeUser(db)

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

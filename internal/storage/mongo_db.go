package storage

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func NewClient(ctx context.Context, cfg Config) (*mongo.Database, *mongo.Client, error) {
	// 1. Construir URI
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin",
		cfg.User, cfg.Password, cfg.Host, cfg.Port)
	// 2. Configurar opciones (con timeout de conexión)
	clientOptions := options.Client().
		ApplyURI(uri).
		SetConnectTimeout(10 * time.Second)
	// 3. Conectar
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		//%w envuelve el error y permite seguir usando errors.Is() o errors.As()
		return nil, nil, fmt.Errorf("failed to connect to mongo: %w", err)
	}
	// 4. Verificar conexión (importante para fallar rápido en el arranque)
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := client.Ping(pingCtx, nil); err != nil {
		return nil, nil, fmt.Errorf("failed to ping mongo: %w", err)
	}
	return client.Database(cfg.Database), client, nil
}

.PHONY: dev run build test clean help coverage-view

# Desarrollo con hot-reload
dev:
	@echo "🚀 Iniciando servidor en modo desarrollo..."
	@pnpm dlx nodemon -e go -x "go run cmd/main.go" --signal SIGTERM --ignore '**/*_test.go'

# Ejecutar sin hot-reload
run:
	@echo "▶️  Ejecutando aplicación..."
	@go run cmd/main.go

# Compilar binario
build:
	@echo "🔨 Compilando binario..."
	@go build -o bin/app cmd/main.go
	@echo "✅ Binario creado en bin/app"

# Ejecutar tests
test:
	@echo "🧪 Ejecutando tests..."
	@go test ./... -v

# Ejecutar tests con coverage
test-coverage:
	@echo "📊 Ejecutando tests con cobertura..."
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@echo "✅ Reporte generado en coverage.html"

# Ver reporte de cobertura en el navegador
coverage-view:
	@if [ ! -f coverage.html ]; then \
		echo "❌ No existe coverage.html. Ejecuta 'make test-coverage' primero."; \
		exit 1; \
	fi
	@echo "🌐 Abriendo reporte de cobertura en http://localhost:3000/coverage.html ..."
	@pnpm dlx http-server -p 3000 -o /coverage.html

# Limpiar archivos generados
clean:
	@echo "🧹 Limpiando archivos generados..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html
	@echo "✅ Limpieza completada"

# Formatear código
fmt:
	@echo "💅 Formateando código..."
	@go fmt ./...

# Verificar código
lint:
	@echo "🔍 Verificando código..."
	@go vet ./...

# Descargar dependencias
deps:
	@echo "📦 Descargando dependencias..."
	@go mod download
	@go mod tidy

# Mostrar ayuda
help:
	@echo "Comandos disponibles:"
	@echo "  make dev           - Iniciar en modo desarrollo con hot-reload"
	@echo "  make run           - Ejecutar aplicación sin hot-reload"
	@echo "  make build         - Compilar binario"
	@echo "  make test          - Ejecutar tests"
	@echo "  make test-coverage - Ejecutar tests con reporte de cobertura"
	@echo "  make coverage-view - Abrir reporte de cobertura en el navegador"
	@echo "  make clean         - Limpiar archivos generados"
	@echo "  make fmt           - Formatear código"
	@echo "  make lint          - Verificar código"
	@echo "  make deps          - Descargar y limpiar dependencias"

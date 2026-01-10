.PHONY: dev run build test clean help coverage-view build-without-debug build-win build-win-without-debug tui

# Desarrollo con hot-reload
dev:
	@echo "🚀 Iniciando servidor en modo desarrollo..."
	@pnpm dlx nodemon -e go -x "go run cmd/api/main.go" --signal SIGTERM --ignore '**/*_test.go'

# Ejecutar sin hot-reload
run:
	@echo "▶️  Ejecutando aplicación..."
	@go run cmd/api/main.go

# Compilar binario
build:
	@echo "🔨 Compilando binario..."
	@go build -o bin/app cmd/api/main.go
	@go build -o bin/tui cmd/tui/main.go
	@echo "✅ Binario creado en bin/app"

build-without-debug:
	@echo "🔨 Compilando binario..."
	@go build -ldflags="-s -w" -o bin/app cmd/api/main.go
	@go build -ldflags="-s -w" -o bin/tui cmd/tui/main.go
	@echo "✅ Binario creado en bin/app"

build-win:
	@echo "🪟 Compilando binarios para Windows (64 bits)..."
	@mkdir -p bin/windows
	@GOOS=windows GOARCH=amd64 go build -o bin/windows/api.exe cmd/api/main.go
	@GOOS=windows GOARCH=amd64 go build -o bin/windows/tui.exe cmd/tui/main.go
	@echo "✅ Binarios de Windows creados en bin/windows/"

build-win-without-debug:
	@echo "🪟 Compilando binarios para Windows (64 bits)..."
	@mkdir -p bin/windows
	@GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o bin/windows/api.exe cmd/api/main.go
	@GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o bin/windows/tui.exe cmd/tui/main.go
	@echo "✅ Binarios de Windows creados en bin/windows/"

tui:
	@echo "🖥️  Abriendo interfaz TUI..."
	@go run cmd/tui/main.go

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
	@echo "Uso: make [comando]"
	@echo ""
	@echo "Comandos disponibles:"
	@echo "  EJECUCIÓN"
	@echo "    dev                          Iniciar en modo desarrollo (hot-reload)"
	@echo "    run                          Ejecutar API sin hot-reload"
	@echo "    tui                          Ejecutar interfaz de terminal interactiva"
	@echo ""
	@echo "  COMPILACIÓN (LINUX)"
	@echo "    build                        Compilar binarios (con debug)"
	@echo "    build-without-debug          Compilar binarios (optimizados/livianos)"
	@echo ""
	@echo "  COMPILACIÓN (WINDOWS)"
	@echo "    build-win                    Compilar .exe para Windows (con debug)"
	@echo "    build-win-without-debug      Compilar .exe para Windows (optimizado)"
	@echo ""
	@echo "  TEST Y CALIDAD"
	@echo "    test                         Ejecutar todos los tests"
	@echo "    test-coverage                Ejecutar tests y generar reporte de cobertura"
	@echo "    coverage-view                Ver reporte de cobertura en navegador (puerto 3000)"
	@echo "    fmt                          Formatear código fuente (go fmt)"
	@echo "    lint                         Verificar estática del código (go vet)"
	@echo ""
	@echo "  MANTENIMIENTO"
	@echo "    deps                         Sincronizar dependencias (download/tidy)"
	@echo "    clean                        Borrar binarios y archivos temporales"


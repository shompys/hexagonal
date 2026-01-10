package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/shompys/hexagonal/internal/bootstrap"
	"github.com/shompys/hexagonal/internal/user/infrastructure/adapters/tui"
)

func main() {
	bootstrap.InitializeUser()
	// Inicializar dependencias de usuario
	Model := tui.InitialModel()

	p := tea.NewProgram(Model, tea.WithAltScreen())
	_, err := p.Run()

	if err != nil {
		log.Fatal(err)
	}

	// Crear CLI
	// userCLI := cli.NewUserCLI(userDeps.UseCase)

	// Contexto
	// ctx := context.Background()

}

package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/shompys/hexagonal/internal/user/domain/ports"
)

type StateTUI struct {
	currentView currentView
	inputValue  textinput.Model
	UserUseCase ports.UserUseCases
}

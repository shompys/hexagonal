package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/shompys/hexagonal/internal/user/domain/ports"
)

func InitialModel(
	UserUseCase ports.UserUseCases,
) tea.Model {
	inputValue := textinput.New()
	inputValue.Focus()
	return StateTUI{
		viewWelcome,
		inputValue,
		UserUseCase,
	}
}

func (m StateTUI) Init() tea.Cmd {
	return textinput.Blink
}

func (m StateTUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//fmt.Printf("Tipo de mensaje o evento: %T\n", msg)
	var cmd tea.Cmd

	if v, ok := msg.(tea.KeyMsg); ok {
		key := v.String()
		if key == "ctrl+c" || key == "esc" {
			return m, tea.Quit
		}
		if key == "enter" {
			if m.currentView == viewWelcome {
				m.currentView = viewForm
			}
		}
	}
	m.inputValue, cmd = m.inputValue.Update(msg)
	return m, cmd
}

func (m StateTUI) View() string {

	switch m.currentView {
	case viewWelcome:
		return m.viewWelcome()
	case viewForm:
		return fmt.Sprintf("Escribí cualquier cosa:\n%s\n (esc para salir)", m.inputValue.View())
	case viewUserTable:
		return m.viewUsersTable(ctx)
	default:
		return "default"
	}
}

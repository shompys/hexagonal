package tui

import "github.com/charmbracelet/bubbles/textinput"

type StateTUI struct {
	currentView currentView
	inputValue  textinput.Model
}

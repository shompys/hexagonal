package tui

type currentView uint8

const (
	viewWelcome currentView = iota
	viewForm
)

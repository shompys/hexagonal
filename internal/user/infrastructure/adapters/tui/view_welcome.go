package tui

import "github.com/charmbracelet/lipgloss"

// TODO: llevar los colores a parte
var (
	purple = lipgloss.Color("99")
	cyan   = lipgloss.Color("86")
	gray   = lipgloss.Color("240")
	white  = lipgloss.Color("255")

	// Estilo para el banner principal
	bannerStyle = lipgloss.NewStyle().
			Foreground(purple).
			PaddingLeft(1)
	// Estilo para la info de sesión
	infoStyle = lipgloss.NewStyle().
			Foreground(cyan).Bold(true)
	// Estilo para el separador
	separatorStyle = lipgloss.NewStyle().
			Foreground(white)
	separatorStyle2 = lipgloss.NewStyle().
			Foreground(gray)
)

func (m StateTUI) viewWelcome() string {
	bannerRaw := `
  ____  _                                          
 / ___|| |__   ___  _ __ ___  _ __  _   _ ___      
 \___ \| '_ \ / _ \| '_ ' _ \| '_ \| | | / __|     
  ___) | | | | (_) | | | | | | |_) | |_| \__ \     
 |____/|_| |_|\___/|_| |_| |_| .__/ \__, |___/     
                             |_|    |___/          
	`
	banner := bannerStyle.Render(bannerRaw)
	// 2. Un subtítulo o info de la sesión
	info := infoStyle.Render("Hola flaquito")
	separatorWhite := separatorStyle.Render("---------------------------------------------------")
	separatorGray := separatorStyle2.Render("---------------------------------------------------")
	nl := "\n"

	return separatorWhite + nl + separatorGray + banner + nl + separatorGray + nl + separatorWhite + nl + info + nl
}

package styles

import "github.com/charmbracelet/lipgloss"

var (
	StyleMenuTitle        = lipgloss.NewStyle().PaddingLeft(4).PaddingBottom(1).Foreground(lipgloss.Color("#3333FF"))
	StyleMenuItem         = lipgloss.NewStyle().PaddingLeft(4)
	StyleMenuItemSelected = lipgloss.NewStyle().PaddingLeft(3).Foreground(lipgloss.Color("#0000FF")) // different padding, because we add [ to front
)

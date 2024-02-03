package styles

import "github.com/charmbracelet/lipgloss"

var (
	StyleMenuTitle         = lipgloss.NewStyle().PaddingLeft(4).PaddingBottom(1).Foreground(lipgloss.Color("#3333FF"))
	StyleMenuItem          = lipgloss.NewStyle().PaddingLeft(4)
	StyleMenuItemSelected  = lipgloss.NewStyle().PaddingLeft(3).Foreground(lipgloss.Color("#0000FF")) // different padding, because we add [ to front
	StyleGameGraphics      = lipgloss.NewStyle().PaddingLeft(4).PaddingRight(4).Foreground(lipgloss.Color("#FF0000")).BorderStyle(lipgloss.RoundedBorder()).MarginLeft(4).MarginRight(4)
	StyleGameUnusedLetter  = lipgloss.NewStyle().Background(lipgloss.Color("#d0c0bd")).Foreground(lipgloss.Color("#000000")).Width(3).Align(lipgloss.Center)
	StyleGameWrongLetter   = lipgloss.NewStyle().Background(lipgloss.Color("#FF0000")).Foreground(lipgloss.Color("#000000")).Width(3).Align(lipgloss.Center)
	StyleGameCorrectLetter = lipgloss.NewStyle().Background(lipgloss.Color("#006600")).Foreground(lipgloss.Color("#000000")).Width(3).Align(lipgloss.Center)
	StyleGameGameOver      = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).MarginLeft(4)
	StyleGameGameWin       = lipgloss.NewStyle().Foreground(lipgloss.Color("#006600")).MarginLeft(4)
	StyleGameQuessLetter   = lipgloss.NewStyle().Background(lipgloss.Color("#f5efb5")).Foreground(lipgloss.Color("#000000")).Width(3).Align(lipgloss.Center)
	StyleGameReplay        = lipgloss.NewStyle().MarginLeft(4)
)

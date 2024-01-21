package model

import tea "github.com/charmbracelet/bubbletea"

type Substate interface {
	Init()
	Update(msg tea.Msg) StateExitCode
	View() string
}

type StateExitCode int

const (
	Nothing StateExitCode = iota
	Quit
	Menu
	Game
	Statistics
	Settings
	About
)

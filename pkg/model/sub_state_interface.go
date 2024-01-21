package model

import tea "github.com/charmbracelet/bubbletea"

type Substate interface {
	Init()
	Update(msg tea.Msg) StateExitCode
	View() string
}

type StateExitCode int

const (
	Nothing    StateExitCode = 0
	Quit       StateExitCode = 1
	Menu       StateExitCode = 2
	Game       StateExitCode = 3
	Statistics StateExitCode = 4
	Settings   StateExitCode = 5
	About      StateExitCode = 6
)

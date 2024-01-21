package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

type MainState struct {
	Text  string
	state Substate
}

func NewMainState() *MainState {
	return &MainState{Text: "Main :)", state: &MenuState{Text: "Menu :O"}}
}

func (m MainState) Init() tea.Cmd {
	return tea.SetWindowTitle("Hangman")
}

func (m MainState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	stateExitCode := m.state.Update(msg)
	if stateExitCode == Quit {
		return m, tea.Quit
	} else if stateExitCode == Game {
		m.state = &GameState{Text: "Game ^^"}
	} else if stateExitCode == Menu {
		m.state = &MenuState{Text: "Menu again :O"}
	}
	return m, nil
}

func (m MainState) View() string {
	text := m.Text
	text += "\n"
	text += m.state.View()
	return text
}

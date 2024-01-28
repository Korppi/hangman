package model

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainState struct {
	Text  string
	state Substate
	count int
}

func NewMainState() *MainState {
	return &MainState{Text: "Main :)", state: &MenuState{Text: "Menu :O"}}
}

type TickMsg time.Time

// Send a message every second.
func tickEvery() tea.Cmd {
	return tea.Every(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m MainState) Init() tea.Cmd {
	return tea.Batch(tea.SetWindowTitle("Hangman"), tickEvery())
}

func (m MainState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case TickMsg:
		m.count++
		return m, tickEvery()
	}
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
	text := m.Text + fmt.Sprint(m.count)
	text += "\n"
	text += m.state.View()
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("228")).
		BorderBackground(lipgloss.Color("63"))
	return style.Render(text)
}

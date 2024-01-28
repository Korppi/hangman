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
	test  int
	flag  bool
}

func NewMainState() *MainState {
	return &MainState{Text: "Main :)", state: &MenuState{Text: "Menu :O"}}
}

type TickMsg time.Time

type FlagMsg int

func ticking() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(time.Millisecond * 250)
		return FlagMsg(1)
	}
}

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
	case FlagMsg:
		m.test--
		if m.test <= 0 {
			m.flag = false
			return m, nil
		}
		return m, ticking()
	case tea.KeyMsg:
		m.test = 10
		m.flag = true
		return m, ticking()
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
	text := m.Text + fmt.Sprint(m.count) + " " + fmt.Sprint(m.test) + " " + fmt.Sprint(m.flag)
	text += "\n"
	text += m.state.View()
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("228")).
		BorderBackground(lipgloss.Color("63"))
	return style.Render(text)
}

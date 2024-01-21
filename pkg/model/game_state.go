package model

import tea "github.com/charmbracelet/bubbletea"

type GameState struct {
	Text string
}

func (m *GameState) Init() {
	// TODO: nothing?
}

func (m *GameState) Update(msg tea.Msg) StateExitCode {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return Quit
		case tea.KeySpace:
			return Menu
		}

	}
	return Nothing
}

func (m *GameState) View() (text string) {
	return m.Text
}

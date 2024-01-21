package model

import tea "github.com/charmbracelet/bubbletea"

type MenuState struct {
	Text string
}

func (m *MenuState) Init() {
	// TODO: nothing?
}

func (m *MenuState) Update(msg tea.Msg) StateExitCode {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return Quit // TODO: at this iteration, lets use just ints. Later on, lets change to enums and if it looks better, then lets change to use msg structs
		case tea.KeyEnter:
			return Game
		}

	}
	return Nothing
}

func (m *MenuState) View() (text string) {
	return m.Text
}

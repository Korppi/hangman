package model

import tea "github.com/charmbracelet/bubbletea"

type MainModel struct {
	text string
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c":
			return m, tea.Quit

		case "k":
			m.text += "k"
		}
	}
	return m, nil
}

func (m MainModel) View() string {
	return m.text
}

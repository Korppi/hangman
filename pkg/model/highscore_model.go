package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

type HighscoreModel struct {
	text string
}

func NewHighscoreModel() *HighscoreModel {
	return &HighscoreModel{
		text: "Nothing yet!",
	}
}

func (m HighscoreModel) Init() tea.Cmd {
	return nil
}

func (m HighscoreModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q", "b":
			return NewMenuModelWithStartingIndex(1), nil
		}

	}
	return m, nil
}

func (m HighscoreModel) View() string {
	return m.text
}

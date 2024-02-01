package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

type MenuModel struct {
	selectionIndex int
	selections     []string
}

func NewMenuModel() *MenuModel {
	return NewMenuModelWithStartingIndex(0)
}

func NewMenuModelWithStartingIndex(startingIndex int) *MenuModel {
	return &MenuModel{
		selectionIndex: startingIndex,
		selections:     []string{"Start game", "Highscore", "Quit"},
	}
}

func (m MenuModel) Init() tea.Cmd {
	return nil
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "crtl+c":
			return m, tea.Quit
		case "up":
			m.selectionIndex = (m.selectionIndex - 1 + len(m.selections)) % len(m.selections)
		case "down":
			m.selectionIndex = (m.selectionIndex + 1) % len(m.selections)
		case "enter":
			switch m.selectionIndex {
			case 0:
				return m, nil
			case 1:
				return NewHighscoreModel(), nil
			case 2:
				return m, tea.Quit
			}
		}

	}
	return m, nil
}

func (m MenuModel) View() string {
	text := "Hangman"
	text += "\n\n"
	for i, v := range m.selections {
		if i == m.selectionIndex {
			text += "[" + v + "]"
		} else {
			text += v
		}
		text += "\n"
	}
	return text
}

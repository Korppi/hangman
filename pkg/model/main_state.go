package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

type MainState struct {
	Text string
}

func NewMainState() *MainState {
	return &MainState{Text: "Main :)"}
}

func (m MainState) Init() tea.Cmd {
	return nil
}

func (m MainState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m MainState) View() string {
	return m.Text
}

package model

import (
	"github.com/Korppi/hangmancli/pkg/styles"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
	Quit  key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Enter}, // first column
		{k.Quit},                // second column
	}
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "move down"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "confirm selection"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q/esc/ctrl+c", "quit"),
	),
}

type MenuModel struct {
	selectionIndex int
	selections     []string
	keys           keyMap
	help           help.Model
}

func NewMenuModel() *MenuModel {
	return NewMenuModelWithStartingIndex(0)
}

func NewMenuModelWithStartingIndex(startingIndex int) *MenuModel {
	model := &MenuModel{
		selectionIndex: startingIndex,
		selections:     []string{"Start game", "Highscore", "Credits", "Quit"},
		keys:           keys,
		help:           help.New(),
	}
	model.help.ShowAll = true
	return model
}

func (m MenuModel) Init() tea.Cmd {
	return tea.SetWindowTitle("Hangman - Menu")
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Up):
			m.selectionIndex = (m.selectionIndex - 1 + len(m.selections)) % len(m.selections)
		case key.Matches(msg, m.keys.Down):
			m.selectionIndex = (m.selectionIndex + 1) % len(m.selections)
		case key.Matches(msg, m.keys.Enter):
			switch m.selectionIndex {
			case 0:
				return m, nil
			case 1:
				return NewHighscoreModel(), nil
			case 2:
				return m, nil
			case 3:
				return m, tea.Quit
			}
		}

	}
	return m, nil
}

func (m MenuModel) View() string {
	text := styles.StyleMenuTitle.Render(`
 __   __  _______  __    _  _______  __   __  _______  __    _ 
|  | |  ||   _   ||  |  | ||       ||  |_|  ||   _   ||  |  | |
|  |_|  ||  |_|  ||   |_| ||    ___||       ||  |_|  ||   |_| |
|       ||       ||       ||   | __ |       ||       ||       |
|       ||       ||  _    ||   ||  ||       ||       ||  _    |
|   _   ||   _   || | |   ||   |_| || ||_|| ||   _   || | |   |
|__| |__||__| |__||_|  |__||_______||_|   |_||__| |__||_|  |__|
	`) // Title text created with this cool tool: https://patorjk.com/software/taag/
	text += "\n"
	for i, v := range m.selections {
		if i == m.selectionIndex {
			text += styles.StyleMenuItemSelected.Render("[" + v + "]")
		} else {
			text += styles.StyleMenuItem.Render(v)
		}
		text += "\n"
	}
	text += "\n" + m.help.View(m.keys)
	return text
}

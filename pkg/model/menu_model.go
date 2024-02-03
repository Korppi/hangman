package model

import (
	"github.com/Korppi/hangmancli/pkg/graphics"
	"github.com/Korppi/hangmancli/pkg/styles"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type menuKeyMap struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
	Quit  key.Binding
	Help  key.Binding
}

func (k menuKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help}
}

func (k menuKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Enter}, // first column
		{k.Help, k.Quit},        // second column
	}
}

var menuKeys = menuKeyMap{
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
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
	Help: key.NewBinding(
		key.WithKeys("h"),
		key.WithHelp("h", "toggle help"),
	),
}

type MenuModel struct {
	selectionIndex int
	selections     []string
	keys           menuKeyMap
	help           help.Model
}

func NewMenuModel() *MenuModel {
	return NewMenuModelWithConfigurations(0, false)
}

// TODO: replace this and other similar functions with function, that takes optional options struct as parameter...
func NewMenuModelWithConfigurations(startingIndex int, helpEnabled bool) *MenuModel {
	model := &MenuModel{
		selectionIndex: startingIndex,
		selections:     []string{"Start game", "Statistics", "Credits", "Quit"},
		keys:           menuKeys,
		help:           help.New(),
	}
	model.help.ShowAll = helpEnabled
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
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Enter):
			switch m.selectionIndex {
			case 0:
				return NewGameModel(m.help.ShowAll), nil
			case 1:
				return NewStatisticsModel(m.help.ShowAll), nil
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
	text := styles.StyleMenuTitle.Render(graphics.Title)
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

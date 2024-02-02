package model

import (
	"github.com/Korppi/hangmancli/pkg/graphics"
	"github.com/Korppi/hangmancli/pkg/styles"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type statisticsKeyMap struct {
	Enter key.Binding
	Quit  key.Binding
	Help  key.Binding
}

func (k statisticsKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help}
}

func (k statisticsKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Enter, k.Help, k.Quit},
	}
}

var statisticsKeys = statisticsKeyMap{
	Enter: key.NewBinding(
		key.WithKeys("enter", "esc"),
		key.WithHelp("enter/esc", "back to menu"),
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

type StatisticsModel struct {
	text string
	keys statisticsKeyMap
	help help.Model
}

func NewStatisticsModel(helpEnabled bool) *StatisticsModel {
	model := &StatisticsModel{
		text: "Nothing yet!",
		keys: statisticsKeys,
		help: help.New(),
	}
	model.help.ShowAll = helpEnabled
	return model
}

func (m StatisticsModel) Init() tea.Cmd {
	return tea.SetWindowTitle("Hangman - Statistics")
}

func (m StatisticsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Enter):
			return NewMenuModelWithConfigurations(1, m.help.ShowAll), nil
		}

	}
	return m, nil
}

func (m StatisticsModel) View() string {
	text := styles.StyleMenuTitle.Render(graphics.Title)
	// TODO: remove hardcoded example and create actual table with actual data from REST api...
	// TODO: filtering? most solved words, hardest words,
	hardcodedExample := "\nWord\t\t\tSuccess rate\n"
	hardcodedExample += "------------------------------------\n"
	hardcodedExample += "example\t\t\t85%\n"
	hardcodedExample += "genius\t\t\t84%\n"
	hardcodedExample += "animation\t\t80%\n"
	hardcodedExample += "professor\t\t45%\n"
	hardcodedExample += "agricultural\t\t31%"
	text += hardcodedExample
	text += "\n\n" + m.help.View(m.keys)
	return text
}

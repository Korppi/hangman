package model

import (
	"math/rand"
	"strings"

	"github.com/Korppi/hangmancli/pkg/graphics"
	"github.com/Korppi/hangmancli/pkg/styles"
	"github.com/Korppi/hangmancli/pkg/utils"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const MIN_WORD_LENGTH = 7
const LETTERS = "abcdefghijklmnopqrstuvwxyz"

var keyboard = [][]string{
	{"Q", "W", "E", "R", "T", "Y", "U", "I", "O", "P"},
	{"A", "S", "D", "F", "G", "H", "J", "K", "L"},
	{"Z", "X", "C", "V", "B", "N", "M"},
}

var Words = []string{}

type gameKeyMap struct {
	Keys key.Binding
	Quit key.Binding
}

func (k gameKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Keys, k.Quit}
}

func (k gameKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Keys, k.Quit},
	}
}

var gameKeys = gameKeyMap{
	Keys: key.NewBinding(
		// first i create []string array using Split. 3 dots (...) unpacks this slice to individual values
		key.WithKeys(strings.Split(LETTERS, "")...),
		key.WithHelp("a-z", "quess letter"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
}

type GameModel struct {
	Word            string
	Quess           string
	maxQuesses      int
	WrongQuessCount int
	UsedLetters     []string
	keys            gameKeyMap
	help            help.Model
}

func NewGameModel(helpEnabled bool) *GameModel {
	word := Words[rand.Intn(len(Words))]
	model := &GameModel{
		Word:            word,
		Quess:           strings.Repeat("_", len(word)),
		maxQuesses:      5,
		WrongQuessCount: 0,
		UsedLetters:     []string{},
		keys:            gameKeys,
		help:            help.New(),
	}
	model.help.ShowAll = helpEnabled
	return model
}

func (m *GameModel) MakeGuess(guess string) {
	if !utils.Contains(m.UsedLetters, guess) {
		correctQuess := false
		m.UsedLetters = append(m.UsedLetters, guess)
		for i, v := range m.Word {
			if guess == string(v) {
				m.ReplaceStringAtIndex(guess, i)
				correctQuess = true
			}
		}
		if !correctQuess {
			m.WrongQuessCount++
		}
	}
}

func (m *GameModel) ReplaceStringAtIndex(s string, i int) {
	m.Quess = m.Quess[:i] + s + m.Quess[i+1:]
}

func (m GameModel) Init() tea.Cmd {
	return tea.SetWindowTitle("Hangman - Game")
}

func (m GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Keys):
			runeString := strings.ToLower(string(msg.Runes))
			if m.maxQuesses > m.WrongQuessCount && strings.Contains(m.Quess, "_") {
				m.MakeGuess(runeString)
			}
		}

	}
	return m, nil
}

func (m GameModel) View() string {
	text := styles.StyleMenuTitle.Render(graphics.Title) + "\n"
	hangmanGraphics := styles.StyleGameGraphics.Render(graphics.HangmanGraphics[m.WrongQuessCount])
	keyboardGraphics := ""
	for i, row := range keyboard {
		for _, letter := range row {
			stringLetter := string(letter)
			if strings.Contains(m.Quess, strings.ToLower(stringLetter)) {
				keyboardGraphics += styles.StyleGameCorrectLetter.Render(stringLetter)
			} else if strings.Contains(strings.Join(m.UsedLetters, ""), strings.ToLower(stringLetter)) {
				keyboardGraphics += styles.StyleGameWrongLetter.Render(stringLetter)
			} else {
				keyboardGraphics += styles.StyleGameUnusedLetter.Render(stringLetter)
			}

		}
		keyboardGraphics += "\n" + strings.Repeat(" ", (i+1)*2) // dirty... but result looks good enough for me
	}
	text += lipgloss.JoinHorizontal(lipgloss.Center, hangmanGraphics, keyboardGraphics)
	text += "\n" + m.Quess + "\n"
	text += "\n\n" + m.help.View(m.keys)
	return text
}

package model

import (
	"math/rand"
	"strings"

	"github.com/Korppi/hangmancli/pkg/utils"
	tea "github.com/charmbracelet/bubbletea"
)

const MIN_WORD_LENGTH = 7

var Words = []string{}
var HangmanGraphics = []string{
	`
	_________
	|       |
	|       O
	|
	|
	|
	|
	----------
	`,
	`
	_________
	|       |
	|       O
	|       |
	|       |
	|      
	|
	----------
	`,
	`
	_________
	|       |
	|       O
	|      /|
	|       |
	|      
	|
	----------
	`,
	`
	_________
	|       |
	|       O
	|      /|\
	|       |
	|      
	|
	----------
	`,
	`
	_________
	|       |
	|       O
	|      /|\
	|       |
	|      / 
	|
	----------
	`,
	`
	_________
	|       |
	|       O
	|      /|\
	|       |
	|      / \
	|
	----------
	`,
}

type Hangman struct {
	Word            string
	Quess           string
	maxQuesses      int
	WrongQuessCount int
	UsedLetters     []string
}

func NewHangmanGame() *Hangman {
	word := Words[rand.Intn(len(Words))]
	return &Hangman{
		Word:            word,
		Quess:           strings.Repeat("_", len(word)),
		maxQuesses:      5,
		WrongQuessCount: 0,
		UsedLetters:     []string{},
	}
}

func (h *Hangman) NewGame() {
	h.Word = Words[rand.Intn(len(Words))]
	h.Quess = strings.Repeat("_", len(h.Word))
	h.WrongQuessCount = 0
	h.UsedLetters = []string{}
}

func (h *Hangman) MakeGuess(guess string) {
	if !utils.Contains(h.UsedLetters, guess) {
		correctQuess := false
		h.UsedLetters = append(h.UsedLetters, guess)
		for i, v := range h.Word {
			if guess == string(v) {
				h.ReplaceStringAtIndex(guess, i)
				correctQuess = true
			}
		}
		if !correctQuess {
			h.WrongQuessCount++
		}
	}
}

func (h *Hangman) ReplaceStringAtIndex(s string, i int) {
	h.Quess = h.Quess[:i] + s + h.Quess[i+1:]
}

func (h Hangman) Init() tea.Cmd {
	return nil
}

func (h Hangman) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyRunes:
			runeString := strings.ToLower(string(msg.Runes))
			if h.maxQuesses > h.WrongQuessCount && strings.Contains(h.Quess, "_") {
				switch runeString {
				case "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "ä", "ö":
					h.MakeGuess(runeString)
				}
			} else {
				switch runeString {
				case "n":
					h.NewGame()
				case "q":
					return h, tea.Quit
				}
			}

		case tea.KeyCtrlC:
			return h, tea.Quit
		}
	}
	return h, nil
}

func (h Hangman) View() string {
	text := HangmanGraphics[h.WrongQuessCount]
	text += "\n" + h.Quess
	text += "\nUsed letters: " + strings.Join(h.UsedLetters, ",")
	if h.maxQuesses <= h.WrongQuessCount {
		text += "\nGame lost! Word was " + h.Word + "\nPress 'n' for new game, 'q' to quit."
	} else if h.maxQuesses > h.WrongQuessCount && !strings.Contains(h.Quess, "_") {
		text += "\nYou won! Press 'n' for new game, 'q' to quit."
	}
	return text
}

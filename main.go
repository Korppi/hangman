package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Korppi/hangmancli/pkg/model"
	tea "github.com/charmbracelet/bubbletea"
)

//go:embed words_alpha.txt
var content embed.FS

func main() {
	file, err := content.Open("words_alpha.txt") //os.Open("words_alpha.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.ToLower(scanner.Text())
		if len(text) >= model.MIN_WORD_LENGTH {
			model.Words = append(model.Words, text)
		}
	}
	file.Close()
	p := tea.NewProgram(model.NewMenuModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

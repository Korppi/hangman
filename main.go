package main

import (
	"fmt"
	"os"

	"github.com/Korppi/hangmancli/pkg/model"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(model.NewHangmanGame())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

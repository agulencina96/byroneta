package main

import (
	"fmt"

	"github.com/agulencina96/byroneta/views"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(views.InitIndexModel())
	_, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
	}
}

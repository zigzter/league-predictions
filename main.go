package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zigzter/league-predictions/models"
)

func main() {
	f, err := tea.LogToFile("./debug.log", "debug")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	m := models.InitialRootModel()
	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

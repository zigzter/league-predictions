package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zigzter/league-predictions/models"
	"github.com/zigzter/league-predictions/utils"
)

func main() {
	configPath := utils.SetupPath()
	utils.InitConfig()
	f, err := tea.LogToFile(configPath+"/debug.log", "debug")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	m := models.InitRootModel()
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

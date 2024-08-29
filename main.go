package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zigzter/league-predictions/models"
	"github.com/zigzter/league-predictions/utils"
)

func main() {
	var dump *os.File
	if _, ok := os.LookupEnv("DEBUG"); ok {
		var err error
		dump, err = os.OpenFile("messages.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
		log.Println("ALJDLKJFLKSJ")
		if err != nil {
			os.Exit(1)
		}
	}
	configPath := utils.SetupPath()
	utils.InitConfig()
	f, err := tea.LogToFile(configPath+"/debug.log", "debug")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	m := models.InitRootModel(dump)
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

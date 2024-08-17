package main

import (
	"log"
	"os"
	"os/signal"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	f, err := tea.LogToFile("./debug.log", "debug")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
}

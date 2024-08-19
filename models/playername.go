package models

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zigzter/league-predictions/utils"
)

type PlayerNameModel struct {
	textinput     textinput.Model
	nameIsMissing bool
}

func InitPlayerNameModel() PlayerNameModel {
	ti := textinput.New()
	ti.Placeholder = "Player name"
	ti.Focus()
	return PlayerNameModel{
		textinput: ti,
	}
}

func (m PlayerNameModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m PlayerNameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			utils.SaveConfig(utils.PlayerNameKey, m.textinput.Value())
			return ChangeView(m, choosePredView)
		}
	}
	var textinputCmd tea.Cmd
	m.textinput, textinputCmd = m.textinput.Update(msg)
	return m, textinputCmd
}

func (m PlayerNameModel) View() string {
	var b strings.Builder
	b.WriteString("Enter the streamer's Riot name\n")
	b.WriteString(m.textinput.View() + "\n")
	return b.String()
}

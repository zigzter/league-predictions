package models

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type AuthKeyModel struct {
	textinput textinput.Model
}

func InitAuthKeyModel() AuthKeyModel {
	ti := textinput.New()
	ti.Placeholder = "API key"
	ti.Focus()
	return AuthKeyModel{
		textinput: ti,
	}
}

func (m AuthKeyModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m AuthKeyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var textinputCmd tea.Cmd
	m.textinput, textinputCmd = m.textinput.Update(msg)
	return m, textinputCmd
}

func (m AuthKeyModel) View() string {
	var b strings.Builder
	b.WriteString("Enter API auth key here\n")
	b.WriteString(m.textinput.View() + "\n")
	return b.String()
}

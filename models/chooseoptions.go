package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

type ChooseOptionsModel struct {
	prediction string
}

func InitChooseOptionsModel() ChooseOptionsModel {
	return ChooseOptionsModel{
		prediction: "",
	}
}

func (m ChooseOptionsModel) Init() tea.Cmd {
	return nil
}

func (m ChooseOptionsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ChangeViewMsg:
		prediction := msg.state.(string)
		m.prediction = prediction
	}
	return m, nil
}

func (m ChooseOptionsModel) View() string {
	return "You picked:" + m.prediction
}

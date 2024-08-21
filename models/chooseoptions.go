package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/viper"
	"github.com/zigzter/league-predictions/utils"
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
	switch msg.(type) {
	case tea.WindowSizeMsg:
		prediction := viper.GetString(utils.PredictionKey)
		m.prediction = prediction
	}
	return m, nil
}

func (m ChooseOptionsModel) View() string {
	return "You picked:" + m.prediction
}

package models

import tea "github.com/charmbracelet/bubbletea"

type choiceState int

const (
	winLoss choiceState = iota
	kills
	deaths
)

type ChoosePredModel struct {
	choice choiceState
}

func InitChoosePredModel() ChoosePredModel {
	return ChoosePredModel{}
}

func (m ChoosePredModel) Init() tea.Cmd {
	return nil
}

func (m ChoosePredModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m ChoosePredModel) View() string {
	return "Choose pred model"
}

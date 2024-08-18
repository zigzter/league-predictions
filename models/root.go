package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

type AppState int

type ChangeStateMsg struct {
	NewState AppState
}

const (
	PlayerNameState AppState = iota
	ChoosePredState
	ChooseOptionsState
	WaitingState
	InProgressState
)

type RootModel struct {
	State AppState
}

func InitialRootModel() RootModel {
	return RootModel{}
}

func (m RootModel) Init() tea.Cmd {
	return nil
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m RootModel) View() string {
	return "Root model"
}

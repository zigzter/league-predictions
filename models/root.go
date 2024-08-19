package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

type sessionState int

type ChangeStateMsg struct {
	NewState sessionState
}

const (
	playerNameView sessionState = iota
	choosePredView
	chooseOptionsView
	waitingView
	inProgressView
)

type RootModel struct {
	state      sessionState
	playerName tea.Model
}

func InitRootModel() RootModel {
	playerNameModel := InitPlayerNameModel()
	return RootModel{
		playerName: playerNameModel,
	}
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

	var cmd tea.Cmd
	switch m.state {
	case playerNameView:
		newPlayerName, newCmd := m.playerName.Update(msg)
		playerModel, ok := newPlayerName.(PlayerNameModel)
		if !ok {
			panic("could not perform assertion on playerName model")
		}
		m.playerName = playerModel
		cmd = newCmd
	}
	return m, cmd
}

func (m RootModel) View() string {
	switch m.state {
	case playerNameView:
		return m.playerName.View()
	default:
		return ""
	}
}

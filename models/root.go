package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zigzter/league-predictions/utils"
)

type sessionState int

type ChangeStateMsg struct {
	NewState sessionState
}

const (
	playerNameView sessionState = iota
	authKeyView
	choosePredView
	chooseOptionsView
	waitingView
	inProgressView
)

type RootModel struct {
	state      sessionState
	playerName tea.Model
	authKey    tea.Model
}

func InitRootModel() RootModel {
	playerNameModel := InitPlayerNameModel()
	authKeyModel := InitAuthKeyModel()
	m := RootModel{
		playerName: playerNameModel,
		authKey:    authKeyModel,
	}
	isAuthKeyMissing := utils.IsAuthKeyMissing()
	if isAuthKeyMissing {
		m.state = authKeyView
	}
	return m
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
	case authKeyView:
		newAuthKey, newCmd := m.authKey.Update(msg)
		authKeyModel, ok := newAuthKey.(AuthKeyModel)
		if !ok {
			panic("could not perform assertion on authKey model")
		}
		m.authKey = authKeyModel
		cmd = newCmd
	}
	return m, cmd
}

func (m RootModel) View() string {
	switch m.state {
	case playerNameView:
		return m.playerName.View()
	case authKeyView:
		return m.authKey.View()
	default:
		return ""
	}
}

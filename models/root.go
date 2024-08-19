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
	authKeyView
	choosePredView
	chooseOptionsView
	waitingView
	inProgressView
)

type RootModel struct {
	state  sessionState
	models map[sessionState]tea.Model
}

func InitRootModel() RootModel {
	playerNameModel := InitPlayerNameModel()
	authKeyModel := InitAuthKeyModel()
	choosePredModel := InitChoosePredModel()
	m := RootModel{
		models: map[sessionState]tea.Model{
			playerNameView: playerNameModel,
			authKeyView:    authKeyModel,
			choosePredView: choosePredModel,
		},
	}
	return m
}

func (m RootModel) Init() tea.Cmd {
	return nil
}

// PropagateUpdate shuttles the message to the correct model's update method
func (m RootModel) PropagateUpdate(msg tea.Msg) tea.Cmd {
	targetModel := m.models[m.state]
	newModel, newCmd := targetModel.Update(msg)
	model, ok := newModel.(tea.Model)
	if !ok {
		panic("could not perform assertion on model")
	}
	m.models[m.state] = model
	return newCmd
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ChangeStateMsg:
		m.state = msg.NewState
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	cmd := m.PropagateUpdate(msg)
	return m, cmd
}

func (m RootModel) View() string {
	return m.models[m.state].View()
}

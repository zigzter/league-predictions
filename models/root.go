package models

import (
	"os"

	"github.com/charmbracelet/bubbles/cursor"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/davecgh/go-spew/spew"
	"github.com/zigzter/league-predictions/utils"
)

type sessionState int

type ChangeViewMsg struct {
	newView sessionState
	state   any
}

const (
	configView sessionState = iota
	choosePredView
	chooseOptionsView
	waitingView
	inProgressView
)

type RootModel struct {
	state  sessionState
	models map[sessionState]tea.Model
	dump   *os.File
}

func InitRootModel(dump *os.File) RootModel {
	configModel := InitConfigModel()
	choosePredModel := InitChoosePredModel()
	chooseOptionsModel := InitChooseOptionsModel()
	isConfigEntryRequired := utils.IsConfigEntryRequired()
	state := sessionState(1)
	if isConfigEntryRequired {
		state = sessionState(0)
	}
	m := RootModel{
		dump: dump,
		models: map[sessionState]tea.Model{
			configView:        configModel,
			choosePredView:    choosePredModel,
			chooseOptionsView: chooseOptionsModel,
		},
		state: state,
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
	if m.dump != nil {
		if _, ok := msg.(cursor.BlinkMsg); !ok {
			spew.Fdump(m.dump, msg)
		}
	}
	switch msg := msg.(type) {
	case ChangeViewMsg:
		m.state = msg.newView
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

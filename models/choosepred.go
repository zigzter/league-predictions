package models

import (
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zigzter/league-predictions/utils"
)

type choiceState int

const (
	winLoss choiceState = iota
	kills
	deaths
)

type ChoosePredModel struct {
	choice choiceState
	list   list.Model
}

type item struct {
	title, description string
}

func (i item) Title() string {
	return i.title
}

func (i item) Description() string {
	return i.description
}

func (i item) FilterValue() string {
	return i.title
}

var items = []list.Item{
	item{title: "win/loss", description: "Whether the streamer wins or loses the game"},
	item{title: "kills", description: "How many kills the streamer gets this game"},
	item{title: "deaths", description: "How many deaths the streamer gets this game"},
}

func InitChoosePredModel() ChoosePredModel {
	list := list.New(items, list.NewDefaultDelegate(), 20, 4)
	list.Title = "Choose a prediction"
	list.SetShowPagination(false)
	list.SetShowFilter(false)
	return ChoosePredModel{
		list: list,
	}
}

func (m ChoosePredModel) Init() tea.Cmd {
	return nil
}

func (m ChoosePredModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			choice, ok := m.list.SelectedItem().(item)
			if !ok {
				log.Fatalln("Item assertion error")
			}
			utils.SaveConfig(utils.PredictionKey, choice.Title())
			return ChangeView(m, chooseOptionsView, choice.Title())
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ChoosePredModel) View() string {
	return m.list.View()
}

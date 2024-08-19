package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/spf13/viper"
	"github.com/zigzter/league-predictions/utils"
)

type ConfigModel struct {
	form *huh.Form
}

func InitConfigModel() ConfigModel {
	m := ConfigModel{}
	riotAPIKey := viper.GetString(utils.RiotAPIKey)
	playerName := viper.GetString(utils.PlayerNameKey)
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("player-name").
				Title("Player's name").
				Value(&playerName),
			huh.NewInput().
				Key("riot-api-key").
				Title("Riot API Key").
				Value(&riotAPIKey),
			huh.NewConfirm().
				Key("save").
				Title("Save Settings").
				Validate(func(v bool) error {
					m.SaveConfig(v)
					return nil
				}).
				Affirmative("Save").
				Negative("Cancel"),
		),
	)
	m.form = form
	return m
}

func (m ConfigModel) SaveConfig(shouldSave bool) {
	if shouldSave {
		utils.SaveConfig(utils.RiotAPIKey, m.form.GetString("riot-api-key"))
		utils.SaveConfig(utils.PlayerNameKey, m.form.GetString("player-name"))
		ChangeView(m, choosePredView)
	}
}

func (m ConfigModel) Init() tea.Cmd {
	return nil
}

func (m ConfigModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}
	if m.form.State == huh.StateCompleted {
		return ChangeView(m, choosePredView)
	}
	return m, tea.Batch(cmds...)
}

func (m ConfigModel) View() string {
	return m.form.View()
}

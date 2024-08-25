package models

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/spf13/viper"
	"github.com/zigzter/league-predictions/twitch"
	"github.com/zigzter/league-predictions/utils"
)

func listenForExternalMsgs(externalMsgs chan tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return <-externalMsgs
	}
}

type ConfigModel struct {
	form     *huh.Form
	authMsgs chan tea.Msg
	status   string
}

func InitConfigModel() ConfigModel {
	m := ConfigModel{
		authMsgs: make(chan tea.Msg, 10),
		status:   "nothing yet",
	}
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
				Key("twitch-auth").
				Title("Start Twitch Auth").
				Affirmative("Yes").
				Negative("No").
				Validate(func(start bool) error {
					if start {
						m.InitTwitchAuth()
					}
					return nil
				}),
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

func (m ConfigModel) InitTwitchAuth() tea.Msg {
	ready := make(chan struct{}, 1)
	cmds := []tea.Cmd{
		twitch.StartLocalServer(ready, m.authMsgs),
		listenForExternalMsgs(m.authMsgs),
	}
	return tea.Batch(cmds...)
}

func (m ConfigModel) SaveConfig(shouldSave bool) {
	if shouldSave {
		utils.SaveConfig(utils.RiotAPIKey, m.form.GetString("riot-api-key"))
		utils.SaveConfig(utils.PlayerNameKey, m.form.GetString("player-name"))
		ChangeView(m, choosePredView, nil)
	}
}

func (m ConfigModel) Init() tea.Cmd {
	return listenForExternalMsgs(m.authMsgs)
}

func (m ConfigModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg.(type) {
	case twitch.ServerStartMsg:
		m.status = "Start server"
	case twitch.ServerStartedMsg:
		m.status = "Server started"
		cmds = append(
			cmds,
			twitch.PromptTwitchAuth(),
			listenForExternalMsgs(m.authMsgs),
		)
	case twitch.AuthOpenMsg:
		m.status = "Auth open"
	case twitch.AuthOpenedMsg:
		m.status = "Auth opened"
	}
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}
	if cmd := listenForExternalMsgs(m.authMsgs); cmd != nil {
		cmds = append(cmds, cmd)
	}
	if m.form.State == huh.StateCompleted {
		return ChangeView(m, choosePredView, nil)
	}
	return m, tea.Batch(cmds...)
}

func (m ConfigModel) View() string {
	var b strings.Builder
	b.WriteString(m.form.View() + "\n")
	b.WriteString(m.status)
	return b.String()
}

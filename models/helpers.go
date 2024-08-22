package models

import tea "github.com/charmbracelet/bubbletea"

// Changes which model the root TUI model is displaying. Used in the Update method.
// Takes in an optional state value to pass any required state to the new view.
func ChangeView(model tea.Model, newView sessionState, state any) (tea.Model, tea.Cmd) {
	return model, tea.Cmd(func() tea.Msg {
		return ChangeViewMsg{newView: newView, state: state}
	})
}

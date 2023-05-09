package views

import (
	tea "github.com/charmbracelet/bubbletea"
)

type ReportsModel struct {
}

func InitReportsModel() ReportsModel {
	return ReportsModel{}
}

func (s ReportsModel) Init() tea.Cmd {
	return nil
}

func (s ReportsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return s, nil
}

func (s ReportsModel) View() string {
	return "pepe peepe pepe"
}

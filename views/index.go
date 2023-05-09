package views

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Shared states between all views
var (
	Token string = ""
)

type IndexModel struct {
	selected     int
	sessionModel tea.Model
	reportsModel tea.Model
}

const (
	SESSION = iota
	REPORTS
)

func (m IndexModel) Init() tea.Cmd {
	return nil
}

func (m IndexModel) View() string {
	if Token == "" {
		return m.sessionModel.View()
	} else {
		switch m.selected {
		case SESSION:
			return m.sessionModel.View()
		case REPORTS:
			return m.reportsModel.View()
		}
		return "ERROR"
	}
}

func (m IndexModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "tab":
			if Token != "" {
				if m.selected == SESSION {
					m.selected = REPORTS
				} else {
					m.selected = SESSION
				}
			}
		}

	}

	switch m.selected {
	case SESSION:
		var cmd tea.Cmd
		m.sessionModel, cmd = m.sessionModel.Update(msg)
		return m, cmd
	case REPORTS:
		var cmd tea.Cmd
		m.reportsModel, cmd = m.reportsModel.Update(msg)
		return m, cmd
	}

	return m, nil
}

func InitIndexModel() IndexModel {
	return IndexModel{
		sessionModel: InitSessionModel(),
		reportsModel: InitReportsModel(),
		selected:     SESSION,
	}
}

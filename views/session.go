package views

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type SessionModel struct {
	username textinput.Model
	password textinput.Model
}
type accessToken struct {
	AccessToken string
}

func InitSessionModel() SessionModel {
	username := textinput.New()
	username.Placeholder = "Username"
	username.Focus()
	username.CharLimit = 20
	username.Width = 20

	password := textinput.New()
	password.Placeholder = "Password"
	password.CharLimit = 20
	password.Width = 20
	password.EchoMode = textinput.EchoPassword

	return SessionModel{
		username: username,
		password: password,
	}
}

func (s SessionModel) Init() tea.Cmd {
	return nil
}

func (s SessionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return s, tea.Quit
		case "enter":
			if s.username.Value() != "" {
				s.username.Blur()
				s.password.Focus()
			}
			if s.password.Value() != "" {
				s.password.Blur()
			}
			if s.username.Value() != "" && s.password.Value() != "" {
				body := fmt.Sprintf(`{"username": "%s", "password": "%s"}`, s.username.Value(), s.password.Value())
				jsonStr := []byte(body)
				res, err := http.Post("http://localhost/api/auth/login", "application/json", bytes.NewBuffer(jsonStr)) //TODO: Change to real url
				if err != nil {
					panic(err)
				}
				accessToken := accessToken{}
				err = json.NewDecoder(res.Body).Decode(&accessToken)
				res.Body.Close()
				if err != nil {
					panic(err)
				}
				if accessToken.AccessToken != "" {
					Token = accessToken.AccessToken
				} else {
					s.username.Reset()
					s.password.Reset()
					s.username.Focus()
				}

			}
		}
	}
	s.username, _ = s.username.Update(msg)
	s.password, _ = s.password.Update(msg)
	return s, nil
}

func (s SessionModel) View() string {
	if Token == "" {
		return fmt.Sprintf("Not Nini Login\n\nIntroduzca su nombre de usuario: %s \nIntroduzca su contraseña: %s", s.username.View(), s.password.View())
	} else {
		return "LOGGED"
	}
}

package main

import (
	"fmt"
	Pages "poker-cli/pages"

	tea "github.com/charmbracelet/bubbletea"
)

type Page int

const (
	Menu Page = iota
	Game
)

type RootModel struct {
	page Page
	Menu Pages.MenuModel
	Game Pages.GameModel
}

func (m RootModel) Init() tea.Cmd { return nil }

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			if m.page == Menu {
				m.page = Game
			} else {
				m.page = Menu
			}
		case "q":
			return m, tea.Quit
		}
	}

	switch m.page {
	case Menu:
		newMenu, cmd := m.Menu.Update(msg)
		m.Menu = newMenu
		return m, cmd
	case Game:
		newGame, cmd := m.Game.Update(msg)
		m.Game = newGame
		return m, cmd
	}
	return m, nil
}

func (m RootModel) View() string {
	switch m.page {
	case Menu:
		return m.Menu.View()
	case Game:
		return m.Game.View()
	}
	return ""
}

func main() {
	root := RootModel{
		page: Menu,
		Menu: Pages.InitialMenuModel(),
		Game: Pages.InitialGameModel(),
	}

	p := tea.NewProgram(root)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}

package main

import (
	"fmt"
	"image/color"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/common-nighthawk/go-figure"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/muesli/gamut"
)

type model struct {
	cursor  int
	choices []string // 0: none, 1: pressed'
	pressed int
	width   int
	height  int
}

var (
	buttonStyle = lipgloss.NewStyle().
			Padding(1, 3).
			Foreground(lipgloss.Color(""))

	buttonStyleFocused = buttonStyle.
				Padding(0, 2).
				Border(lipgloss.RoundedBorder()).
				Foreground(lipgloss.Color("#08D9D6")).
				BorderForeground(lipgloss.Color("#08D9D6"))
)

func initialModel() model {
	return model{
		cursor:  0,
		choices: []string{"(p) Play", "(q) Quit"},
		pressed: -1,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "left", "h":
			if m.cursor > 0 {
				m.cursor--
			}
		case "right", "l":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		}
	}
	return m, nil
}

func padFigureOutput(raw string) string {
	lines := strings.Split(raw, "\n")
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	blends := gamut.Blends(lipgloss.Color("#FFB236"), lipgloss.Color("#FF4C67"), maxWidth)

	for i, line := range lines {
		lines[i] = lipgloss.
			NewStyle().
			Width(maxWidth).
			Render(gradient(lipgloss.NewStyle(), line, blends))
	}
	return strings.Join(lines, "\n")
}

func gradient(base lipgloss.Style, s string, colors []color.Color) string {
	var str string
	for i, ss := range s {
		color, _ := colorful.MakeColor(colors[i%len(colors)])
		str = str + base.Foreground(lipgloss.Color(color.Hex())).Render(string(ss))
	}
	return str
}

func (m model) View() string {
	if m.pressed != -1 {
		return fmt.Sprintf("You selected: %s\n", m.choices[m.pressed])
	}

	var buttons []string

	for i, choice := range m.choices {
		if i == m.cursor {
			buttons = append(buttons, buttonStyleFocused.Render(choice))
		} else {
			buttons = append(buttons, buttonStyle.Render(choice))
		}
		buttons = append(buttons, "   ")
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, buttons...)

	rendered := lipgloss.NewStyle().Render("\n\n\n" + row + "\n\nUse ← → and Enter:")

	fig := figure.NewFigure("Poker-CLI", "doom", false).String()
	title := padFigureOutput(fig)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, title+"\n\n"+rendered)
}

func main() {

	f, err := tea.LogToFile("debug.log", "debug")

	if err != nil {
		log.Fatalf("err: %d", err)
	}

	defer f.Close()
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

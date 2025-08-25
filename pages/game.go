package Pages

import (
	"fmt"
	"math/rand"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type GameModel struct {
	playerHand []int
	dealerHand []int
	deck       []int
	width      int
	height     int
}

func InitialGameModel() GameModel {
	m := GameModel{playerHand: []int{}, dealerHand: []int{}, deck: make([]int, 52)}

	for i := range m.deck {
		m.deck[i] = i
	}

	drawCard(&m.deck, &m.dealerHand, 1)

	drawCard(&m.deck, &m.playerHand, 2)

	return m
}

func (m GameModel) Update(msg tea.Msg) (GameModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}

func (m GameModel) Init() tea.Cmd {
	return nil
}

func (m GameModel) View() string {
	rendered := make([]string, len(m.playerHand))

	for i, c := range m.playerHand {
		rendered[i] = cardToStr(c)
	}
	// return fmt.Sprintln("Player: ", m.playerHand) + fmt.Sprintln("Dealer: ", m.dealerHand) + fmt.Sprintln("Deck", m.deck)
	stack := lipgloss.JoinHorizontal(
		lipgloss.Left,
		rendered...,
	)
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Bottom, stack)
}

var cardStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	Width(3).
	Height(2).
	Align(lipgloss.Center).
	Padding(0, 1)

func cardToStr(cardID int) string {
	card := cardID%13 + 1
	var cardStr string

	switch card {
	case 1:
		cardStr = "A"
	case 10:
		cardStr = "X"
	case 11:
		cardStr = "J"
	case 12:
		cardStr = "Q"
	case 13:
		cardStr = "K"
	default:
		cardStr = strconv.Itoa(card)
	}
	return cardStyle.Render(cardStr)
}

func drawCard(deck *[]int, hand *[]int, amount int) {
	for i := 0; i < amount; i++ {
		if len(*deck) == 0 {
			fmt.Println("Deck is empty!")
			return
		}

		idx := rand.Intn(len(*deck))

		*hand = append(*hand, (*deck)[idx])

		*deck = append((*deck)[:idx], (*deck)[idx+1:]...)
	}
}

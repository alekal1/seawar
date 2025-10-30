package game

import (
	"aleksale/seawar/cells"
	"aleksale/seawar/statusMsg"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SeaWarModel struct {
	CoordinatesInput textinput.Model

	Board              [][]string
	OpponentGuessBoard [][]string

	ShipsPlaced         map[int]int
	OpponentsGuessShips map[int]int

	AIHits        []cells.Cell
	OpponentBoard [][]string

	PlayerTurn bool
	status     string
}

func (m *SeaWarModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *SeaWarModel) View() string {
	boardStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 5).
		Width(35)

	playerBoard := boardStyle.Render(m.DisplayPlayerBoard())
	gap := lipgloss.NewStyle().PaddingRight(4).Render("")

	if m.allShipsPlaced() {
		opponentGuessBoard := boardStyle.Render(m.DisplayOpponentGuessBoard())

		return lipgloss.JoinHorizontal(
			lipgloss.Top,
			playerBoard,
			gap,
			opponentGuessBoard)
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		playerBoard,
		gap,
		getRemainingShips(m))
}

func (m *SeaWarModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter:
			if m.CoordinatesInput.Value() == "" {
				m.status = statusMsg.NoInput()
				return m, nil
			}

			if !m.allShipsPlaced() {
				return m.placeShip(m.CoordinatesInput.Value()), nil
			}

			if m.PlayerTurn {
				return m.playerTurn(m.CoordinatesInput.Value()), nil
			}
		}
	}

	if !m.PlayerTurn && !m.allPlayerShipsSunk() && !m.allOpponentShipsSunk() {
		m.opponentTurn()
	}

	m.CoordinatesInput, cmd = m.CoordinatesInput.Update(msg)

	return m, cmd
}

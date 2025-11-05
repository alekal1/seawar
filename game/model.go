package game

import (
	"aleksale/seawar/opponent"
	"aleksale/seawar/player"
	"aleksale/seawar/statusMsg"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SeaWarModel struct {
	CoordinatesInput textinput.Model

	Player   *player.Player
	Opponent *opponent.Opponent

	status string
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

	if m.Player.AllShipsPlaced() {
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

			if !m.Player.AllShipsPlaced() {
				m.status = m.Player.PlaceShip(m.CoordinatesInput.Value())
				m.CoordinatesInput.SetValue("")
				return m, nil
			}

			if m.Player.PlayerTurn {
				m.status = m.Player.MakeTurn(m.CoordinatesInput.Value(), m.Opponent.Board)

				if m.Opponent.AllShipsSunk(m.Player.GuessBoard) {
					m.status = statusMsg.PlayerWin()
					return m, nil
				}

				m.CoordinatesInput.SetValue("")

				return m, nil
			}
		}
	}

	if !m.Player.PlayerTurn && !m.Player.AllShipsSunk() && !m.Opponent.AllShipsSunk(m.Player.Board) {
		isPlayersTurn, status := m.Opponent.MakeTurn(m.Player.Board)
		m.status = status
		if isPlayersTurn {
			m.Player.PlayerTurn = true
		}
	}

	if m.Player.AllShipsPlaced() && m.Player.AllShipsSunk() {
		m.status = statusMsg.PlayerLost()
	}

	m.CoordinatesInput, cmd = m.CoordinatesInput.Update(msg)

	return m, cmd
}

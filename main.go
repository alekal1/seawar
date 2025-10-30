package main

import (
	"aleksale/seawar/variables"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"aleksale/seawar/game"
)

func makeBoard() [][]string {
	rows, cols := 10, 10

	matrix := make([][]string, rows)
	for i := range matrix {
		matrix[i] = make([]string, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = variables.EmptySpace
		}
	}

	return matrix
}

func initialModel() *game.SeaWarModel {
	ci := textinput.New()
	ci.Placeholder = "Ship coordinates..."
	ci.Width = 20
	ci.Focus()
	ci.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
	ci.TextStyle = lipgloss.NewStyle().Bold(true)

	return &game.SeaWarModel{
		CoordinatesInput:   ci,
		Board:              makeBoard(),
		OpponentGuessBoard: makeBoard(),
		OpponentBoard:      game.GenerateRandomOpponentBoard(),
		ShipsPlaced:        make(map[int]int),
		PlayerTurn:         true,
	}
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}

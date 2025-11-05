package main

import (
	"aleksale/seawar/opponent"
	"aleksale/seawar/player"
	"aleksale/seawar/util"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"aleksale/seawar/game"
)

func initialModel() *game.SeaWarModel {
	ci := textinput.New()
	ci.Placeholder = "Ship coordinates..."
	ci.Width = 20
	ci.Focus()
	ci.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
	ci.TextStyle = lipgloss.NewStyle().Bold(true)

	return &game.SeaWarModel{
		CoordinatesInput: ci,

		Player:   player.NewPlayer(util.MakeEmptyBoard(), util.MakeEmptyBoard(), make(map[int]int)),
		Opponent: opponent.NewOpponent(),
	}
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}

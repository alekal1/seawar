package game

import (
	"aleksale/seawar/variables"
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func (m *SeaWarModel) DisplayPlayerBoard() string {
	view := variables.BoardTop
	for index, row := range m.Board {
		view += fmt.Sprintf("%v ", variables.BoardLeft[index])
		for _, col := range row {
			view += " " + colorize(col)
		}
		view += "\n"
	}
	view += fmt.Sprintf("\n %v \n", m.status)
	view += m.CoordinatesInput.View()
	return view
}

func (m *SeaWarModel) DisplayOpponentGuessBoard() string {
	view := variables.BoardTop
	for index, row := range m.OpponentGuessBoard {
		view += fmt.Sprintf("%v ", variables.BoardLeft[index])
		for _, col := range row {
			view += " " + colorize(col)
		}
		view += "\n"
	}
	return view
}

func getRemainingShips(m *SeaWarModel) string {
	cruiserNum := FleetLimits[4] - m.ShipsPlaced[4]
	destroyerNum := FleetLimits[3] - m.ShipsPlaced[3]
	submarineNum := FleetLimits[2] - m.ShipsPlaced[2]
	boatNum := FleetLimits[1] - m.ShipsPlaced[1]

	columns := []table.Column{
		{Title: "Type", Width: 20},
		{Title: "Remaining", Width: 10},
	}

	rows := []table.Row{
		{"Cruiser (4 cells)", strconv.Itoa(cruiserNum)},
		{"Destroyer (3 cells)", strconv.Itoa(destroyerNum)},
		{"Submarine (2 cells)", strconv.Itoa(submarineNum)},
		{"Boat (1 cell)", strconv.Itoa(boatNum)},
	}

	s := table.DefaultStyles()
	s.Selected = lipgloss.NewStyle()

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithHeight(7))

	t.SetStyles(s)

	return t.View()
}

func colorize(col string) string {
	switch col {
	case variables.EmptySpace:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("4")).Render(col)
	case variables.MissedGuess:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Render(col)
	case variables.Ship:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Render(col)
	case variables.DefeatedShip:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Render(col)
	default:
		return col
	}
}

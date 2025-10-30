package game

import (
	"aleksale/seawar/cells"
	"aleksale/seawar/statusMsg"
	"aleksale/seawar/variables"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var FleetLimits = map[int]int{
	1: 4,
	2: 3,
	3: 2,
	4: 1,
}

func (m *SeaWarModel) placeShip(shipCoordinates string) tea.Model {
	var shipCells []cells.Cell

	multipleCoordinates := strings.Contains(shipCoordinates, "-")
	if multipleCoordinates {
		coordinates := strings.Split(shipCoordinates, "-")
		start, end := coordinates[0], coordinates[1]

		startRow, startCol, err := cells.ParseCoordinate(start)
		endRow, endCol, err := cells.ParseCoordinate(end)

		if err != nil {
			m.CoordinatesInput.SetValue("")
			m.status = statusMsg.InvalidInput()
			return m
		}

		shipCells, _ = cells.CellsBetween(startRow, startCol, endRow, endCol)
	} else {
		row, col, err := cells.ParseCoordinate(shipCoordinates)
		shipCells, err = cells.CellsBetween(row, col, row, col)

		if err != nil {
			m.CoordinatesInput.SetValue("")
			m.status = statusMsg.InvalidInput()
			return m
		}
	}

	if ok, err := m.canPlaceShip(shipCells); !ok {
		m.CoordinatesInput.SetValue("")
		m.status = statusMsg.ErrorStatus(err)
		return m
	}

	size := len(shipCells)
	if ok, err := m.canPlaceShipOfSize(size); !ok {
		m.CoordinatesInput.SetValue("")
		m.status = statusMsg.ErrorStatus(err)
		return m
	}

	m.placeShipAtCells(shipCells)
	m.ShipsPlaced[size]++
	m.CoordinatesInput.SetValue("")

	m.status = statusMsg.ShipPlaced(shipCoordinates)
	return m
}

func (m *SeaWarModel) canPlaceShip(cells []cells.Cell) (bool, error) {
	for _, cell := range cells {
		if m.Board[cell.Row][cell.Col] != variables.EmptySpace {
			return false, fmt.Errorf("cell %s%v is already occupied", variables.BoardLeft[cell.Row], cell.Col+1)
		}
	}
	return true, nil
}

func (m *SeaWarModel) canPlaceShipOfSize(size int) (bool, error) {
	limit, exists := FleetLimits[size]
	if !exists {
		return false, fmt.Errorf("invalid ship size %d", size)
	}

	if m.ShipsPlaced[size] >= limit {
		return false, fmt.Errorf("cannot place more ships of size %d", size)
	}

	return true, nil
}

func (m *SeaWarModel) placeShipAtCells(cells []cells.Cell) {
	for _, cell := range cells {
		m.Board[cell.Row][cell.Col] = variables.Ship
	}
}

func (m *SeaWarModel) allShipsPlaced() bool {
	for size, limit := range FleetLimits {
		if m.ShipsPlaced[size] < limit {
			return false
		}
	}
	return true
}

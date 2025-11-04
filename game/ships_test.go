package game

import (
	"aleksale/seawar/cells"
	"aleksale/seawar/variables"
	"testing"
)

func TestCanPlaceShips(t *testing.T) {
	m := &SeaWarModel{
		Board: makeBoard(),
	}

	m.Board[0][0] = variables.Ship

	tests := []struct {
		name  string
		cells []cells.Cell
		want  bool
	}{
		{"All empty", []cells.Cell{{1, 1}}, true},
		{"One occupied", []cells.Cell{{0, 0}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := m.canPlaceShip(tt.cells)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanPlaceShipOfSize(t *testing.T) {
	m := &SeaWarModel{
		ShipsPlaced: map[int]int{
			1: 4,
			2: 2,
		},
	}

	tests := []struct {
		name string
		size int
		want bool
	}{
		{"Maxed out size 1", 1, false},
		{"Available size 2", 2, true},
		{"Invalid size", 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := m.canPlaceShipOfSize(tt.size)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllShipsPlaced(t *testing.T) {
	m := &SeaWarModel{
		ShipsPlaced: map[int]int{
			1: 4,
			2: 3,
			3: 2,
			4: 1,
		},
	}

	if !m.allShipsPlaced() {
		t.Errorf("expected all ship placed")
	}

	m.ShipsPlaced[2] = 2
	if m.allShipsPlaced() {
		t.Errorf("expected not all ships placed")
	}
}

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

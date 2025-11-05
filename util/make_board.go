package util

import (
	"aleksale/seawar/cells"
	"aleksale/seawar/variables"
	"math/rand"
	"time"
)

func MakeEmptyBoard() [][]string {
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

func MakeRandomlyFilledBoard() [][]string {
	rows, cols := 10, 10
	rand.NewSource(time.Now().UnixNano())

	matrix := make([][]string, rows)
	for i := range matrix {
		matrix[i] = make([]string, cols)
		for j := range matrix[i] {
			matrix[i][j] = variables.EmptySpace
		}
	}

	for size, count := range variables.FleetLimits {
		for n := 0; n < count; n++ {
			placeRandomShip(matrix, size)
		}
	}

	return matrix
}

func placeRandomShip(board [][]string, size int) {
	rows, cols := len(board), len(board[0])
	placed := false

	for !placed {
		vertical := rand.Intn(2) == 0
		row := rand.Intn(rows)
		col := rand.Intn(cols)

		shipCells := []cells.Cell{}

		if vertical {
			if row+size > rows {
				continue
			}
			for i := 0; i < size; i++ {
				shipCells = append(shipCells, cells.Cell{Row: row + i, Col: col})
			}
		} else {
			if col+size > cols {
				continue
			}
			for i := 0; i < size; i++ {
				shipCells = append(shipCells, cells.Cell{Row: row, Col: col + i})
			}
		}

		free := true
		for _, c := range shipCells {
			if board[c.Row][c.Col] != variables.EmptySpace {
				free = false
				break
			}
		}

		if free {
			for _, c := range shipCells {
				board[c.Row][c.Col] = variables.Ship
			}
			placed = true
		}
	}
}

package game

import (
	"aleksale/seawar/variables"
	"fmt"
	"testing"
)

func TestGenerateRandomOpponentBoard(t *testing.T) {
	for i := 0; i < 50; i++ {
		board := GenerateRandomOpponentBoard()

		if len(board) != 10 {
			t.Fatalf("Expected 10 rows, got %d", len(board))
		}

		for _, row := range board {
			if len(row) != 10 {
				t.Fatalf("Expected 10 columns, got %d", len(board))
			}
		}

		totalShips, err := calculateTotalShips(board)
		if err != nil {
			t.Fatalf("Unexpected cell value: %v", err)
		}

		expectedShips := calculateExpectedShips()

		if totalShips != expectedShips {
			t.Fatalf("Expected %d total ships cells, got %d", expectedShips, totalShips)
		}
	}
}

func calculateTotalShips(board [][]string) (int, error) {
	totalShips := 0
	for _, row := range board {
		for _, cell := range row {
			if cell != variables.EmptySpace && cell != variables.Ship {
				return 0, fmt.Errorf("cell: %q", cell)
			}
			if cell == variables.Ship {
				totalShips++
			}
		}
	}
	return totalShips, nil
}

func calculateExpectedShips() int {
	expectedShips := 0
	for size, count := range FleetLimits {
		expectedShips += size * count
	}
	return expectedShips
}

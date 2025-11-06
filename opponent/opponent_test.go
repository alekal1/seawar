package opponent

import (
	"aleksale/seawar/cells"
	"aleksale/seawar/util"
	"aleksale/seawar/variables"
	"testing"
)

func TestOpponent_MakeTurn_HuntMode_Miss(t *testing.T) {
	op := NewOpponent(util.MakeRandomlyFilledBoard())
	playerBoard := util.MakeEmptyBoard()

	isPlayerTurn, status := op.MakeTurn(playerBoard)

	if !isPlayerTurn {
		t.Error("Expected turn to be passed to player after miss")
	}

	if status == "" {
		t.Error("Expected a status message after miss")
	}

	foundMissed := false
	for _, row := range playerBoard {
		for _, col := range row {
			if col == variables.MissedGuess {
				foundMissed = true
			}
		}
	}

	if !foundMissed {
		t.Error("Expected a MissedGuess mark on the player's board")
	}
}

func TestOpponent_MakeTurn_HuntMode_Hit(t *testing.T) {
	op := NewOpponent(util.MakeRandomlyFilledBoard())
	playerBoard := makeFullBoard(variables.Ship)

	isPlayerTurn, status := op.MakeTurn(playerBoard)

	if isPlayerTurn {
		t.Error("Expected turn not to be passed to player after hit")
	}

	if status == "" {
		t.Error("Expected a status message after hit")
	}

	foundDefeated := false
	for _, row := range playerBoard {
		for _, col := range row {
			if col == variables.DefeatedShip {
				foundDefeated = true
			}
		}
	}

	if !foundDefeated {
		t.Error("Expected a DefeatedShip mark on the player's board")
	}
}

func TestOpponent_MakeTurn_TargetMode(t *testing.T) {
	op := NewOpponent(util.MakeRandomlyFilledBoard())
	playerBoard := util.MakeEmptyBoard()

	playerBoard[0][0] = variables.DefeatedShip
	playerBoard[0][1] = variables.Ship
	playerBoard[1][0] = variables.Ship

	op.Hits = append(op.Hits, cells.Cell{Row: 0, Col: 0})

	isPlayerTurn, status := op.MakeTurn(playerBoard)

	if isPlayerTurn {
		t.Error("Expected turn not to be passed to player after hit")
	}

	if status == "" {
		t.Error("Expected a status message after hit")
	}

	if playerBoard[0][1] != variables.DefeatedShip && playerBoard[1][0] != variables.DefeatedShip {
		t.Error("Expected at least one adjacent cell to be marked as DefeatedShip")
	}
}

func TestOpponent_AllShipsSunk(t *testing.T) {
	op := NewOpponent(util.MakeRandomlyFilledBoard())
	playerGuessBoard := util.MakeEmptyBoard()

	op.Board = util.MakeEmptyBoard()
	op.Board[0][0] = variables.Ship

	sunk := op.AllShipsSunk(playerGuessBoard)

	if sunk {
		t.Error("Expected all ships not to be sunk")
	}

	playerGuessBoard[0][0] = variables.DefeatedShip

	sunk = op.AllShipsSunk(playerGuessBoard)

	if !sunk {
		t.Error("Expected all ships to be sunk")
	}
}

func makeFullBoard(mark string) [][]string {
	rows, cols := 10, 10

	matrix := make([][]string, rows)
	for i := range matrix {
		matrix[i] = make([]string, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = mark
		}
	}

	return matrix
}

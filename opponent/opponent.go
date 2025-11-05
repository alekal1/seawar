package opponent

import (
	"aleksale/seawar/cells"
	"aleksale/seawar/statusMsg"
	"aleksale/seawar/util"
	"aleksale/seawar/variables"
)

type Opponent struct {
	Board [][]string
	Hits  []cells.Cell
}

func NewOpponent() *Opponent {
	return &Opponent{
		Board: util.MakeRandomlyFilledBoard(),
	}
}

func (op *Opponent) MakeTurn(playerBoard [][]string) (bool, string) {
	if len(op.Hits) > 0 {
		return op.targetModeTurn(playerBoard)
	}

	return op.huntModeTurn(playerBoard)
}

func (op *Opponent) huntModeTurn(playerBoard [][]string) (bool, string) {
	for {
		row := util.RandInt(0, len(playerBoard)-1)
		col := util.RandInt(0, len(playerBoard[0])-1)

		if playerBoard[row][col] == variables.EmptySpace {
			playerBoard[row][col] = variables.MissedGuess
			status := statusMsg.MISS(variables.BoardLeft[row], col+1, false)
			return true, status
		}

		if playerBoard[row][col] == variables.Ship {
			playerBoard[row][col] = variables.DefeatedShip
			status := statusMsg.HIT(variables.BoardLeft[row], col+1, false)
			op.Hits = append(op.Hits, cells.Cell{Row: row, Col: col})
			return false, status
		}
	}
}

func (op *Opponent) targetModeTurn(playerBoard [][]string) (bool, string) {
	lastKnownHit := op.Hits[len(op.Hits)-1]

	directions := []cells.Cell{
		{Row: -1, Col: 0},
		{Row: 1, Col: 0},
		{Row: 0, Col: -1},
		{Row: 0, Col: 1},
	}

	for _, d := range directions {
		r := lastKnownHit.Row + d.Row
		c := lastKnownHit.Col + d.Col

		if r < 0 || r >= len(playerBoard) || c < 0 || c >= len(playerBoard) {
			continue
		}

		cellValue := playerBoard[r][c]

		if cellValue == variables.DefeatedShip || cellValue == variables.MissedGuess {
			continue
		}

		if cellValue == variables.Ship {
			playerBoard[r][c] = variables.DefeatedShip
			status := statusMsg.HIT(variables.BoardLeft[r], c+1, false)
			op.Hits = append(op.Hits, cells.Cell{Row: r, Col: c})
			return false, status
		}

		if cellValue == variables.EmptySpace {
			playerBoard[r][c] = variables.MissedGuess
			status := statusMsg.MISS(variables.BoardLeft[r], c+1, false)
			return true, status
		}
	}

	op.Hits = nil
	return op.huntModeTurn(playerBoard)
}

func (op *Opponent) AllShipsSunk(playerGuessBoard [][]string) bool {
	for r := 0; r < len(op.Board); r++ {
		for c := 0; c < len(op.Board[r]); c++ {
			if op.Board[r][c] == variables.Ship && playerGuessBoard[r][c] != variables.DefeatedShip {
				return false
			}
		}
	}
	return true
}

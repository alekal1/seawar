package player

import (
	"aleksale/seawar/cells"
	"aleksale/seawar/statusMsg"
	"aleksale/seawar/variables"
	"fmt"
	"strings"
)

type Player struct {
	Board       [][]string
	GuessBoard  [][]string
	ShipsPlaced map[int]int
	PlayerTurn  bool
}

func NewPlayer(board [][]string, guessBoard [][]string, shipsPlaced map[int]int) *Player {
	return &Player{
		Board:       board,
		GuessBoard:  guessBoard,
		ShipsPlaced: shipsPlaced,
		PlayerTurn:  true,
	}
}

func (p *Player) MakeTurn(coord string, opponentBoard [][]string) string {
	if !p.PlayerTurn {
		return statusMsg.NotYourTurn()
	}

	row, col, err := cells.ParseCoordinate(coord)
	if err != nil {
		return statusMsg.InvalidGuess(err)
	}

	if p.GuessBoard[row][col] != variables.EmptySpace {
		return statusMsg.AlreadyGuessed()
	}

	if opponentBoard[row][col] == variables.Ship {
		p.GuessBoard[row][col] = variables.DefeatedShip
		return statusMsg.HIT(variables.BoardLeft[row], col+1, true)
	} else {
		p.GuessBoard[row][col] = variables.MissedGuess
		p.PlayerTurn = false
		return statusMsg.MISS(variables.BoardLeft[row], col+1, true)
	}
}

func (p *Player) PlaceShip(shipCoordinates string) string {
	var shipCells []cells.Cell

	multipleCoordinates := strings.Contains(shipCoordinates, "-")
	if multipleCoordinates {
		coordinates := strings.Split(shipCoordinates, "-")
		start, end := coordinates[0], coordinates[1]

		startRow, startCol, err := cells.ParseCoordinate(start)
		endRow, endCol, err := cells.ParseCoordinate(end)

		if err != nil {
			return statusMsg.InvalidInput()
		}

		shipCells, _ = cells.CellsBetween(startRow, startCol, endRow, endCol)
	} else {
		row, col, err := cells.ParseCoordinate(shipCoordinates)
		shipCells, err = cells.CellsBetween(row, col, row, col)

		if err != nil {
			return statusMsg.InvalidInput()
		}
	}

	if ok, err := p.canPlaceShip(shipCells); !ok {
		return statusMsg.ErrorStatus(err)
	}

	size := len(shipCells)
	if ok, err := p.canPlaceShipOfSize(size); !ok {
		return statusMsg.ErrorStatus(err)
	}

	p.placeShipAtCells(shipCells)
	p.ShipsPlaced[size]++

	return statusMsg.ShipPlaced(shipCoordinates)
}

func (p *Player) canPlaceShip(cells []cells.Cell) (bool, error) {
	for _, cell := range cells {
		if p.Board[cell.Row][cell.Col] != variables.EmptySpace {
			return false, fmt.Errorf("cell %s%v is already occupied", variables.BoardLeft[cell.Row], cell.Col+1)
		}
	}
	return true, nil
}

func (p *Player) canPlaceShipOfSize(size int) (bool, error) {
	limit, exists := variables.FleetLimits[size]
	if !exists {
		return false, fmt.Errorf("invalid ship size %d", size)
	}

	if p.ShipsPlaced[size] >= limit {
		return false, fmt.Errorf("cannot place more ships of size %d", size)
	}

	return true, nil
}

func (p *Player) placeShipAtCells(cells []cells.Cell) {
	for _, cell := range cells {
		p.Board[cell.Row][cell.Col] = variables.Ship
	}
}

func (p *Player) AllShipsSunk() bool {
	for r := 0; r < len(p.Board); r++ {
		for c := 0; c < len(p.Board[r]); c++ {
			if p.Board[r][c] == variables.Ship {
				return false
			}
		}
	}
	return true
}

func (p *Player) AllShipsPlaced() bool {
	for size, limit := range variables.FleetLimits {
		if p.ShipsPlaced[size] < limit {
			return false
		}
	}
	return true
}

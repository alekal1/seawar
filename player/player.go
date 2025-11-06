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
	shipCells, parseError := p.parseShipCoordinates(shipCoordinates)

	if parseError != nil {
		return statusMsg.InvalidInput()
	}

	if placeShipError := p.canPlaceShip(shipCells); placeShipError != nil {
		return statusMsg.ErrorStatus(placeShipError)
	}

	size := len(shipCells)
	if shipSizeError := p.canPlaceShipOfSize(size); shipSizeError != nil {
		return statusMsg.ErrorStatus(shipSizeError)
	}

	p.placeShipAtCells(shipCells)
	p.ShipsPlaced[size]++

	return statusMsg.ShipPlaced(shipCoordinates)
}

func (p *Player) parseShipCoordinates(shipCoordinates string) ([]cells.Cell, error) {
	var shipCells []cells.Cell

	multipleCoordinates := strings.Contains(shipCoordinates, "-")
	if multipleCoordinates {
		coordinates := strings.Split(shipCoordinates, "-")
		start, end := coordinates[0], coordinates[1]

		startRow, startCol, err := cells.ParseCoordinate(start)
		if err != nil {
			return nil, err
		}

		endRow, endCol, err := cells.ParseCoordinate(end)

		if err != nil {
			return nil, err
		}

		shipCells, _ = cells.CellsBetween(startRow, startCol, endRow, endCol)
	} else {
		row, col, err := cells.ParseCoordinate(shipCoordinates)
		if err != nil {
			return nil, err
		}

		shipCells, err = cells.CellsBetween(row, col, row, col)
		if err != nil {
			return nil, err
		}
	}

	return shipCells, nil
}

func (p *Player) canPlaceShip(cells []cells.Cell) error {
	for _, cell := range cells {
		if p.Board[cell.Row][cell.Col] != variables.EmptySpace {
			return fmt.Errorf("cell %s%v is already occupied", variables.BoardLeft[cell.Row], cell.Col+1)
		}
	}
	return nil
}

func (p *Player) canPlaceShipOfSize(size int) error {
	limit, exists := variables.FleetLimits[size]
	if !exists {
		return fmt.Errorf("invalid ship size %d", size)
	}

	if p.ShipsPlaced[size] >= limit {
		return fmt.Errorf("cannot place more ships of size %d", size)
	}

	return nil
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

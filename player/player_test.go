package player

import (
	"aleksale/seawar/statusMsg"
	"aleksale/seawar/util"
	"aleksale/seawar/variables"
	"fmt"
	"strings"
	"testing"
)

func TestPlaceShip_singleCoord_alreadyOccupied(t *testing.T) {
	player := NewPlayer(
		util.MakeEmptyBoard(),
		util.MakeEmptyBoard(),
		make(map[int]int),
	)

	player.Board[0][0] = variables.Ship

	status := player.PlaceShip("A1")

	if !strings.Contains(status, "Error") {
		t.Error("Expected error, when cell is occupied")
	}
}

func TestPlaceShip_singleCoord_limitExceeded(t *testing.T) {
	player := NewPlayer(
		util.MakeEmptyBoard(),
		util.MakeEmptyBoard(),
		make(map[int]int),
	)

	for i := 1; i < 5; i++ {
		status := player.PlaceShip(fmt.Sprintf("A%v", i))

		if strings.Contains(status, "Error") {
			t.Error("Expected no error when limit is not exceeded")
		}
	}

	status := player.PlaceShip("A5")
	if !strings.Contains(status, "Error") {
		t.Errorf("Expected error status, got %v", status)
	}
}

func TestPlaceShip_singleCoord_invalidInput(t *testing.T) {
	player := NewPlayer(
		util.MakeEmptyBoard(),
		util.MakeEmptyBoard(),
		make(map[int]int),
	)

	coordinates := []string{
		"A",
		"9",
		"A ",
		"A  ",
		"",
		"X99",
		"-",
		"X1",
		"B20",
	}

	for _, c := range coordinates {
		t.Run(c, func(t *testing.T) {
			status := player.PlaceShip(c)
			if status != statusMsg.InvalidInput() {
				t.Errorf("Expected invalid input error, got %v", status)
			}
		})
	}
}

func TestPlaceShip_singleCoord(t *testing.T) {
	player := NewPlayer(
		util.MakeEmptyBoard(),
		util.MakeEmptyBoard(),
		make(map[int]int),
	)

	status := player.PlaceShip("A1")

	if status != statusMsg.ShipPlaced("A1") {
		t.Errorf("Expected shipPlaced message, got %v", status)
	}

	if player.Board[0][0] != variables.Ship {
		t.Error("Expected ship to be placed at 0 0")
	}

	if player.ShipsPlaced[1] == 0 {
		t.Error("Expected one ship with size 1")
	}
}

func TestPlaceShip_multipleCoord_invalidInput(t *testing.T) {
	player := NewPlayer(
		util.MakeEmptyBoard(),
		util.MakeEmptyBoard(),
		make(map[int]int),
	)

	coordinates := []string{
		"1-A",
		"A-1",
		"A1-A",
		"A1- ",
		" -A1",
		" - ",
		"  -  ",
		"-",
		"X1-A4",
		"A1-B2",
		"A99-B1",
	}

	for _, c := range coordinates {
		t.Run(c, func(t *testing.T) {
			status := player.PlaceShip(c)

			if status != statusMsg.InvalidInput() {
				if !strings.Contains(status, "Error") {
					t.Errorf("Expected %v error, got %v", statusMsg.InvalidInput(), status)
				}
			}
		})
	}

}

func TestPlaceShip_multipleCoord_oneCellOccupied(t *testing.T) {
	player := NewPlayer(
		util.MakeEmptyBoard(),
		util.MakeEmptyBoard(),
		make(map[int]int),
	)

	player.PlaceShip("A2")

	status := player.PlaceShip("A1-A4")

	if !strings.Contains(status, "Error") {
		t.Errorf("Expected error status, got %v", status)
	}
}

func TestPlaceShip_multipleCoord(t *testing.T) {
	player := NewPlayer(
		util.MakeEmptyBoard(),
		util.MakeEmptyBoard(),
		make(map[int]int),
	)

	status := player.PlaceShip("A1-A4")

	if strings.Contains(status, "Error") || status == statusMsg.InvalidInput() {
		t.Errorf("Expected ShipPlaced message, got %v", status)
	}
}

func TestAllSipsPlaced(t *testing.T) {
	player := NewPlayer(
		util.MakeEmptyBoard(),
		util.MakeEmptyBoard(),
		make(map[int]int),
	)

	player.ShipsPlaced[1] = 3
	player.ShipsPlaced[2] = 3
	player.ShipsPlaced[3] = 2
	player.ShipsPlaced[4] = 1

	if shipsPlaced := player.AllShipsPlaced(); shipsPlaced {
		t.Errorf("Not all ships are placed, expected false, got %v", shipsPlaced)
	}

	player.ShipsPlaced[1] = 4

	if shipsPlaced := player.AllShipsPlaced(); !shipsPlaced {
		t.Errorf("Expecetd all ships to be placed, got %v", shipsPlaced)
	}
}

func TestAllShipsSunk(t *testing.T) {
	player := NewPlayer(
		util.MakeEmptyBoard(),
		util.MakeEmptyBoard(),
		make(map[int]int),
	)

	player.Board[0][0] = variables.Ship

	if allSunk := player.AllShipsSunk(); allSunk {
		t.Error("Not all ships are sunk")
	}

	player.Board[0][0] = variables.DefeatedShip

	if allSunk := player.AllShipsSunk(); !allSunk {
		t.Error("Expected all ships to be sunk")
	}
}

func TestMakeTurn_notPlayerTurn(t *testing.T) {
	player := NewPlayer(
		util.MakeEmptyBoard(),
		util.MakeEmptyBoard(),
		make(map[int]int),
	)

	opBoard := util.MakeEmptyBoard()

	player.PlayerTurn = false

	status := player.MakeTurn("A1", opBoard)

	if statusMsg.NotYourTurn() != status {
		t.Errorf("Expected status %v, got %v", statusMsg.NotYourTurn(), status)
	}
}

func TestMakeTurn_invalidCoordinate(t *testing.T) {
	player := NewPlayer(
		util.MakeEmptyBoard(),
		util.MakeEmptyBoard(),
		make(map[int]int),
	)

	opBoard := util.MakeEmptyBoard()

	status := player.MakeTurn("A99", opBoard)

	if !strings.Contains(status, "Invalid") {
		t.Errorf("Expected invalid status, got %v", status)
	}
}

func TestMakeTurn_alreadyGuessed(t *testing.T) {
	player := NewPlayer(
		util.MakeEmptyBoard(),
		util.MakeEmptyBoard(),
		make(map[int]int),
	)
	opBoard := util.MakeEmptyBoard()

	player.GuessBoard[0][0] = variables.DefeatedShip

	status := player.MakeTurn("A1", opBoard)

	if statusMsg.AlreadyGuessed() != status {
		t.Errorf("Expected %v, got %v", statusMsg.AlreadyGuessed(), status)
	}
}

func TestMakeTurn_hit(t *testing.T) {
	player := NewPlayer(
		util.MakeEmptyBoard(),
		util.MakeEmptyBoard(),
		make(map[int]int),
	)
	opBoard := util.MakeEmptyBoard()

	opBoard[0][0] = variables.Ship

	status := player.MakeTurn("A1", opBoard)

	if !strings.Contains(status, "HIT") {
		t.Errorf("Expected HIT status, got %v", status)
	}

	if player.GuessBoard[0][0] != variables.DefeatedShip {
		t.Error("Expected to mark cell as defeated ship at guess board")
	}
}

func TestMakeTurn_miss(t *testing.T) {
	player := NewPlayer(
		util.MakeEmptyBoard(),
		util.MakeEmptyBoard(),
		make(map[int]int),
	)
	opBoard := util.MakeEmptyBoard()

	status := player.MakeTurn("A1", opBoard)

	if !strings.Contains(status, "MISS") {
		t.Errorf("Expected MISS status, got %v", status)
	}

	if player.GuessBoard[0][0] != variables.MissedGuess {
		t.Error("Expected to mark cell as defeated ship at guess board")
	}
}

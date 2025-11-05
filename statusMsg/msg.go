package statusMsg

import "fmt"

func NoInput() string {
	return "No input!"
}

func InvalidInput() string {
	return "Invalid input!"
}

func ErrorStatus(err error) string {
	return fmt.Sprintf("Error: %v", err)
}

func ShipPlaced(shipCoordinates string) string {
	return fmt.Sprintf("Placed ship at %v", shipCoordinates)
}

func PlayerWin() string {
	return "YOU WON! All enemy ships destroyed!"
}

func PlayerLost() string {
	return "YOU LOST! All your ships were destroyed!"
}

func NotYourTurn() string {
	return "Not your turn"
}

func AlreadyGuessed() string {
	return "Already guessed this cell"
}

func InvalidGuess(err error) string {
	return fmt.Sprintf("Invalid guess: %v", err)
}

func HIT(row string, col int, player bool) string {
	if player {
		return fmt.Sprintf("HIT at %s%d", row, col)
	}
	return fmt.Sprintf("Opponent HIT at %s%d", row, col)
}

func MISS(row string, col int, player bool) string {
	if player {
		return fmt.Sprintf("MISS at %s%d", row, col)
	}
	return fmt.Sprintf("Opponent missed at %s%d", row, col)
}

package variables

var (
	BoardTop     = "   1 2 3 4 5 6 7 8 9 10 \n"
	BoardLeft    = map[int]string{0: "A", 1: "B", 2: "C", 3: "D", 4: "E", 5: "F", 6: "G", 7: "H", 8: "I", 9: "J", 10: "K"}
	Ship         = "■"
	DefeatedShip = "X"
	EmptySpace   = "~"
	MissedGuess  = "."
	FleetLimits  = map[int]int{
		1: 4,
		2: 3,
		3: 2,
		4: 1,
	}
)

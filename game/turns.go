package game

//func (m *SeaWarModel) playerTurn(coord string) tea.Model {
//	if m.allOpponentShipsSunk() {
//		m.status = statusMsg.PlayerAlreadyWon()
//		return m
//	}
//
//	if !m.PlayerTurn {
//		m.status = statusMsg.NotYourTurn()
//		return m
//	}
//
//	row, col, err := cells.ParseCoordinate(coord)
//	if err != nil {
//		m.status = statusMsg.InvalidGuess(err)
//		m.CoordinatesInput.SetValue("")
//	}
//
//	if m.OpponentGuessBoard[row][col] != variables.EmptySpace {
//		m.status = statusMsg.AlreadyGuessed()
//		m.CoordinatesInput.SetValue("")
//		return m
//	}
//
//	if m.OpponentBoard[row][col] == variables.Ship {
//		m.OpponentGuessBoard[row][col] = variables.DefeatedShip
//		m.status = statusMsg.HIT(variables.BoardLeft[row], col+1, true)
//
//		if m.allOpponentShipsSunk() {
//			m.status = statusMsg.PlayerWin()
//		}
//	} else {
//		m.OpponentGuessBoard[row][col] = variables.MissedGuess
//		m.status = statusMsg.MISS(variables.BoardLeft[row], col+1, true)
//
//		m.PlayerTurn = false
//		m.opponentTurn()
//	}
//
//	m.CoordinatesInput.SetValue("")
//	return m
//}

//func (m *SeaWarModel) opponentTurn() {
//	if len(m.AIHits) > 0 {
//		m.targetModeTurn()
//	} else {
//		m.huntModeTurn()
//	}
//
//	if m.allPlayerShipsSunk() {
//		m.status = statusMsg.PlayerLost()
//	}
//}
//
//// Random shots
//func (m *SeaWarModel) huntModeTurn() {
//	for {
//		row := util.RandInt(0, len(m.Board)-1)
//		col := util.RandInt(0, len(m.Board[0])-1)
//
//		if m.Board[row][col] == variables.EmptySpace {
//			m.Board[row][col] = variables.MissedGuess
//			m.status = statusMsg.MISS(variables.BoardLeft[row], col+1, false)
//			m.PlayerTurn = true
//			return
//		}
//
//		if m.Board[row][col] == variables.Ship {
//			m.Board[row][col] = variables.DefeatedShip
//			m.status = statusMsg.HIT(variables.BoardLeft[row], col+1, false)
//			m.AIHits = append(m.AIHits, cells.Cell{Row: row, Col: col})
//			return
//		}
//	}
//}
//
//// Adjacent cells shots
//func (m *SeaWarModel) targetModeTurn() {
//	lastKnownHit := m.AIHits[len(m.AIHits)-1]
//
//	directions := []cells.Cell{
//		{Row: -1, Col: 0},
//		{Row: 1, Col: 0},
//		{Row: 0, Col: -1},
//		{Row: 0, Col: 1},
//	}
//
//	for _, d := range directions {
//		r := lastKnownHit.Row + d.Row
//		c := lastKnownHit.Col + d.Col
//
//		if r < 0 || r >= len(m.Board) || c < 0 || c >= len(m.Board[0]) {
//			continue
//		}
//
//		cellValue := m.Board[r][c]
//
//		if cellValue == variables.DefeatedShip || cellValue == variables.MissedGuess {
//			continue
//		}
//
//		if cellValue == variables.Ship {
//			m.Board[r][c] = variables.DefeatedShip
//			m.status = statusMsg.HIT(variables.BoardLeft[r], c+1, false)
//			m.AIHits = append(m.AIHits, cells.Cell{Row: r, Col: c})
//			return
//		}
//
//		if cellValue == variables.EmptySpace {
//			m.Board[r][c] = variables.MissedGuess
//			m.status = statusMsg.MISS(variables.BoardLeft[r], c+1, false)
//			m.PlayerTurn = true
//			return
//		}
//	}
//
//	m.AIHits = nil
//	m.huntModeTurn()
//}
//
//func (m *SeaWarModel) allOpponentShipsSunk() bool {
//	for r := 0; r < len(m.OpponentBoard); r++ {
//		for c := 0; c < len(m.OpponentBoard[r]); c++ {
//			if m.OpponentBoard[r][c] == variables.Ship && m.OpponentGuessBoard[r][c] != variables.DefeatedShip {
//				return false
//			}
//		}
//	}
//	return true
//}

//func (m *SeaWarModel) allPlayerShipsSunk() bool {
//	for r := 0; r < len(m.Board); r++ {
//		for c := 0; c < len(m.Board[r]); c++ {
//			if m.Board[r][c] == variables.Ship {
//				return false
//			}
//		}
//	}
//	return true
//}

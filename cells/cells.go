package cells

import (
	"aleksale/seawar/variables"
	"fmt"
	"strconv"
	"strings"
)

type Cell struct {
	Row int
	Col int
}

func ParseCoordinate(coord string) (int, int, error) {

	rowLetter := coord[:1]
	colPart := coord[1:]

	var row int
	found := false
	for k, v := range variables.BoardLeft {
		if v == strings.ToUpper(rowLetter) {
			row = k
			found = true
			break
		}
	}

	if !found {
		return 0, 0, fmt.Errorf("invalid row letter")
	}

	col, err := strconv.Atoi(colPart)
	if err != nil || col < 1 || col > 10 {
		return 0, 0, fmt.Errorf("invalid column number")
	}
	col--

	return row, col, nil
}

func CellsBetween(startRow, startCol, endRow, endCol int) ([]Cell, error) {
	var cells []Cell

	if startRow != endRow && startCol != endCol {
		return nil, fmt.Errorf("diagonal ships not allowed")
	}

	// Horizontal
	if startRow == endRow {
		if startCol > endCol {
			startCol, endCol = endCol, startCol
		}
		for c := startCol; c <= endCol; c++ {
			cells = append(cells, Cell{startRow, c})
		}
	} else {
		// Vertical
		if startRow > endRow {
			startRow, endRow = endRow, startRow
		}

		for r := startRow; r <= endRow; r++ {
			cells = append(cells, Cell{r, startCol})
		}
	}

	return cells, nil
}

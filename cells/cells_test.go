package cells

import (
	"slices"
	"testing"
)

func TestParseCoordinate(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantX   int
		wantY   int
		wantErr bool
	}{
		{"Valid A1", "A1", 0, 0, false},
		{"Valid a1", "A1", 0, 0, false},
		{"Valid A10", "A10", 0, 9, false},
		{"Valid J10", "J10", 9, 9, false},
		{"Invalid X10", "X10", 0, 0, true},
		{"Invalid A99", "A99", 0, 0, true},
		{"Invalid A0", "A99", 0, 0, true},
		{"Invalid 99", "99", 0, 0, true},
		{"Invalid A", "A", 0, 0, true},
		{"Empty string", "", 0, 0, true},
		{"Single space", " ", 0, 0, true},
		{"Two spaces", "  ", 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY, err := ParseCoordinate(tt.input)

			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseCoordinate(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}

			if !tt.wantErr {
				if gotX != tt.wantX {
					t.Errorf("ParseCoordinate(%q) gotX = %v, want %v", tt.input, gotX, tt.wantX)
				}

				if gotY != tt.wantY {
					t.Errorf("ParseCoordinate(%q) gotY = %v, want %v", tt.input, gotY, tt.wantY)
				}
			}

		})
	}
}

func TestCellsBetween(t *testing.T) {
	tests := []struct {
		name      string
		startRow  int
		startCol  int
		endRow    int
		endCol    int
		wantCells []Cell
		wantErr   bool
	}{
		{"One cell", 0, 0, 0, 0, []Cell{
			{0, 0}}, false},
		{"Four horizontal cells", 0, 0, 0, 3, []Cell{
			{0, 0},
			{0, 1},
			{0, 2},
			{0, 3}}, false},
		{"Four vertical cells", 0, 0, 3, 0, []Cell{
			{0, 0},
			{1, 0},
			{2, 0},
			{3, 0}}, false},
		{"Invalid Diagonal cells", 0, 0, 1, 1, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotCells, err := CellsBetween(tt.startRow, tt.startCol, tt.endRow, tt.endCol)

			if (err != nil) != tt.wantErr {
				t.Fatalf("CellsBetween(%v, %v, %v, %v) error = %v, wantErr %v", tt.startRow, tt.startCol, tt.endRow, tt.endCol, err, tt.wantErr)
			}

			if !slices.Equal(gotCells, tt.wantCells) {
				t.Errorf("CellsBetween(%v, %v, %v, %v) gotCells = %v, want = %v", tt.startRow, tt.startCol, tt.endRow, tt.endCol, gotCells, tt.wantCells)
			}
		})
	}
}

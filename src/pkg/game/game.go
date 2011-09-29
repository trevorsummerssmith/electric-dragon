package game

import (
	"os"
)

const (
	// Maximum number of characters on a board
	MaxCharacters = 100
)

type Idx struct {
	Row int
	Col int
}

type RelativeIdx struct {
	X int
	Y int
}

type CompassDir int

const (
	N CompassDir = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

func NewRelativeIdx(dir CompassDir) *RelativeIdx {
	switch dir {
	case N:  return &RelativeIdx{0,   1}
	case NE: return &RelativeIdx{1,   1}
	case E:  return &RelativeIdx{1,   0}
	case SE: return &RelativeIdx{1,  -1}
	case S:  return &RelativeIdx{0,  -1}
	case SW: return &RelativeIdx{-1, -1}
	case W:  return &RelativeIdx{-1,  0}
	case NW: return &RelativeIdx{-1,  1}
	}
	return nil
}

type Board struct {
	// Number of rows in the board
	rows int

	// Number of cells in a row
	cols int

	// Slice of the Cells that make up the board
	// Dimensions are: [rows][cols]Cell
	cells [][]Cell

	// Characters the board knows about
	characters []Character
}

func NewBoard(rows int, cols int) *Board {
	cells := make([][]Cell, rows)

	for i := 0; i < rows; i++ {
		cells[i] = make([]Cell, cols)
		// Allocate the cells
		for j := 0; j < cols; j++ {
			idx := Idx{Row: i, Col: j}
			cells[i][j] = NewBasicCell(idx, UndefinedCellType)
		}
	}

	// Setup characters
	characters := make([]Character, 1, MaxCharacters)

	return &Board{
	rows: rows,
	cols: cols,
	cells: cells,
	characters: characters,
	}
}

func (b *Board) GetCell(idx Idx) (cell Cell, err os.Error) {
	if !b.ValidIdx(idx) {
		err = os.NewError("Index out of bounds")
		return
	}
	cell = b.cells[idx.Row][idx.Col]
	return
}

func (b *Board) ValidIdx(idx Idx) bool {
	return idx.Row < b.rows && idx.Col < b.cols
}

func (b *Board) ClampIdx(idx *Idx) {
	if idx.Row >= b.rows {
		idx.Row = b.rows - 1
	}
	if idx.Col >= b.cols {
		idx.Col = b.cols - 1
	}
}

func (b *Board) PlaceObject(pobj PlacedObject, idx Idx) os.Error {
	if !b.ValidIdx(idx) {
		return os.NewError("Index out of bounds")
	}
	// Tell object where it is
	pobj.SetIdx(idx)
	// Tell cell object is on it
	b.cells[idx.Row][idx.Col].AddPlacedObject(pobj)
	return nil
}

func (b *Board) AddCharacter(c Character) {
	b.characters = append(b.characters, c)
}

func (i *Idx) AddRelativeIdx(ridx RelativeIdx) {
	i.Row += ridx.Y
	i.Col += ridx.X
}

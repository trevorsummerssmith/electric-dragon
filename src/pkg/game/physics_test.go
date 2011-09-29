package game

import (
	"testing"
	"container/vector"
)

func assertEqualIdx(t *testing.T, idx Idx, expected Idx) {
	if idx.Row != expected.Row {
		t.Errorf("Row should be %d, is %d", expected.Row, idx.Row)
	}
	if idx.Col != expected.Col {
		t.Errorf("Col should be %d, is %d", expected.Col, idx.Col)
	}
}

func assertCellCharactersEqual(t *testing.T, c Cell, expected *vector.Vector) {
	guys := c.PlacedObjects()
	if guys.Len() != expected.Len() {
		t.Errorf("Cell should only have %d characters, has %d",
			expected.Len(), guys.Len())
		return
	}

	// Len is same so compare guys... expect the same order
	for i := 0; i < guys.Len(); i++ {
		if guys.At(i) != expected.At(i) {
			t.Errorf("Character at position %d should be: %v, is %v",
				i, guys.At(i), expected.At(i))
		}
	}
}

func TestMove(t *testing.T) {
	// 5x5 Board
	board := NewBoard(5, 5)
	physics := &Physics{Board: board}

	c := &BasicCharacter{totalEnergy: 10, currentEnergy: 10}
	c.SetIdx(Idx{Row: 2, Col: 1})

	// Move one index
	idxs := make([]RelativeIdx, 1)
	idxs[0].Y = 2
	idxs[0].X = 1
	m := &Move{RelativeIndices: idxs}
	physics.move(c, m)

	// Ensure move happened to character
	expectedIdx := Idx{Row: 4, Col: 2}
	assertEqualIdx(t, c.Idx(), expectedIdx)

	// Ensure move happened to cell
	cell, err := board.GetCell(expectedIdx)
	if err != nil {
		t.Errorf("Invalid index passed to board (SERIOUS PROBLEM)")
	}
	assertCellCharactersEqual(t, cell, &vector.Vector{c})
}

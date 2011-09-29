package game

import (
	"testing"
	)


func TestBoard(t *testing.T) {
	b := NewBoard(10, 20)
	if b.rows != 10 {
		t.Errorf("Rows must be 10 but is %d", b.rows)
	}

	// Make a character and place it on the board
	guy := &BasicCharacter{totalHP: 10, currentHP: 10}
	idx := Idx{Row: 4, Col: 8}

	b.PlaceObject(guy, idx)

	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			guys := b.cells[i][j].PlacedObjects()
			n := guys.Len()
			if i == 4 && j == 8 {
				if n != 1 {
					t.Errorf("This row should have 1 guy")
				}
			} else {
				if n != 0 {
					t.Errorf("This row should have 0 guys")
				}
			}
		}
	}
}

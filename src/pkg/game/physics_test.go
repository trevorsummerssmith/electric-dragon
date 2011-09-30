package game

import (
	"testing"
)

func assertEqualIdx(t *testing.T, idx Idx, expected Idx) {
	if idx.Row != expected.Row {
		t.Errorf("Row should be %d, is %d", expected.Row, idx.Row)
	}
	if idx.Col != expected.Col {
		t.Errorf("Col should be %d, is %d", expected.Col, idx.Col)
	}
}

func assertCellCharactersEqual(t *testing.T, c Cell, expected []PlacedObject) {
	guys := c.PlacedObjects()
	if len(guys) != len(expected) {
		t.Errorf("Cell should only have %d characters, has %d",
			len(expected), len(guys))
		return
	}

	// Len is same so compare guys... expect the same order
	for i := 0; i < len(guys); i++ {
		if guys[i] != expected[i] {
			t.Errorf("Character at position %d should be: %v, is %v",
				i, guys[i], expected[i])
		}
	}
}

func TestMove(t *testing.T) {
	// 5x5 Board
	board := NewBoard(5, 5)
	physics := &Physics{Board: board}

	c := &BasicCharacter{totalEnergy: 10, currentEnergy: 10}
	c.SetIdx(Idx{Row: 2, Col: 1})

	// Move two steps
	steps := make([]CompassDir, 2)
	steps[0] = NE
	steps[1] = N
	m := &Move{Steps: steps}
	physics.move(c, m)

	// Ensure move happened to character
	expectedIdx := Idx{Row: 4, Col: 2}
	assertEqualIdx(t, c.Idx(), expectedIdx)

	// Ensure move happened to cell
	cell, err := board.GetCell(expectedIdx)
	if err != nil {
		t.Errorf("Invalid index passed to board (SERIOUS PROBLEM)")
	}
	assertCellCharactersEqual(t, cell, []PlacedObject{c})
}

func TestAttack(t *testing.T) {
	// 5x5 Board
	board := NewBoard(5, 5)
	physics := &Physics{Board: board}

	// Guy at 2, 1
	c1 := &BasicCharacter{totalEnergy: 10, currentEnergy: 10}
	board.PlaceObject(c1, Idx{Row: 2, Col: 1})

	// Guy at 2, 2
	c2 := &BasicCharacter{totalEnergy: 10, currentEnergy: 10,
	                      totalHP: 20, currentHP: 20}
	board.PlaceObject(c2, Idx{Row: 2, Col: 2})

	// C1 attacks C2
	attack := &Attack{Target: RelativeIdx{Y: 0, X: 1}}
	physics.attack(c1, attack)

	if c1.currentEnergy != 8 {
		t.Errorf("Character energy should be %d, is %d", 8, c1.currentEnergy)
	}
	// C2 has 1 defense point so 9 damage is done
	if c2.currentHP != 11 {
		t.Errorf("Character 2 HP should be %d, is %d", 11, c2.currentHP)
	}
	
}
package game

import (
	"os"
)

const (
	// Maximum number of characters that can be on a single board cell
	MaxGuysOnCell int = 9
)

type Physics struct {
	Board *Board
}

func NewPhysics() *Physics {
	return &Physics{}
}

func (p *Physics) Update(c Character, a Action) os.Error {
	switch obj := a.(type) {
	case *Move:
		return p.move(c, obj)
	case *Attack:
		return p.attack(c, obj)
	default:
		return os.NewError("Unknown Action type")
	}
	return nil
}

func (p *Physics) attack(c Character, attack *Attack) os.Error {
	// Is it possible for the character to move ridx?
	return nil
}

/*
 Moves the RelativeIdx in the *Move. Each index is checked for valid placement:
 _not_ the path from the character's current position to that place.

 It is perhaps more correct to think of this function as transporting the unit
*/
func (p *Physics) move(c Character, move *Move) os.Error {

	// Is it possible for the character to move ridx?
	for i := 0; i < len(move.Steps); i++ {
		ridx := NewRelativeIdx(move.Steps[i])
		newIdx := c.Idx()
		newIdx.AddRelativeIdx(*ridx)

		// 1) Character-specific move function check (can they move?)
		// TODO check attributes of the character and other stuff.

		// 3) Ensure move is inside of board
		p.Board.ClampIdx(&newIdx)

		// Grab the relevant cell
		cell, err := p.Board.GetCell(newIdx)
		if err != nil {
			return err
		}

		// 2) Energy check
		// Get relevant param data together
		moveParams := &MoveParams{Move: move, CellType: cell.CellType()}
		cost := c.EnergyCost(moveParams)
		energy := c.Energy()
		if cost > energy {
			return os.NewError("Not enough energy")
		}

		// 4) Ensure character can fit on cell
		// If not, move is to same cell
		numGuys := cell.PlacedObjects().Len()
		if numGuys >= MaxGuysOnCell {
			newIdx = c.Idx()
		}

		// All systems go -- make the move, take away energy
		cell.AddPlacedObject(c)
		c.SetIdx(newIdx)
		c.SetEnergy(energy - cost)
	}
	return nil
}

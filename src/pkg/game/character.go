package game

import (
	"os"
)

type Character interface {
	HP() uint
	SetHP(hp uint) os.Error

	Energy() uint
	SetEnergy(e uint) os.Error

	EnergyCost(p ActionParams) uint

	Idx() Idx

	// Note this does not do any validation.
	// This is only meant to be a setter method.
	SetIdx(idx Idx)
}

type BasicCharacter struct {
	totalHP uint
	currentHP uint
	totalEnergy uint
	currentEnergy uint
	idx Idx
}

func (b *BasicCharacter) HP() uint {
	return b.currentHP
}

func (b *BasicCharacter) SetHP(hp uint) os.Error {
	if hp > b.totalHP {
		return os.NewError("HP out of range")
	}
	b.currentHP = hp
	return nil
}

func (b *BasicCharacter) Energy() uint {
	return b.currentEnergy
}

func (b *BasicCharacter) SetEnergy(e uint) os.Error {
	if e > b.totalEnergy {
		return os.NewError("Energy out of range")
	}
	b.currentEnergy = e
	return nil
}

func (b *BasicCharacter) EnergyCost(p ActionParams) uint {
	switch obj := p.(type) {
	case *MoveParams:
		// This is just an example for now
		// This character moves easily through everything except swamp
		switch obj.CellType {
		case Grass:
			return 1
		case Mountain:
			return 1
		case Sand:
			return 1
		case Swamp:
			return 2
		default:
			return 1 // TODO(trevor) error handle
		 }
	case *AttackParams:
		return 2
	}

	return 3
}

func (b *BasicCharacter) Idx() Idx {
	return b.idx
}

func (b *BasicCharacter) SetIdx(idx Idx) {
	b.idx = idx
}
package game

import (
	"container/vector"
)

type PlacedObject interface {
	Idx() Idx
	SetIdx(idx Idx)
}

type CellType int

const (
	UndefinedCellType CellType = iota
	Grass
	Mountain
	Sand
	Swamp
	)

type Cell interface {
	PlacedObjects() *vector.Vector
	SetPlacedObjects(p *vector.Vector)
	AddPlacedObject(p PlacedObject)
	RemovePlacedObject(p PlacedObject)
	Idx() Idx
	SetIdx(idx Idx)
	CellType() CellType
}

type BasicCell struct {
	objects *vector.Vector
	idx Idx
	cellType CellType
}

func NewBasicCell(idx Idx, cellType CellType) *BasicCell {
	v := new(vector.Vector)
	return &BasicCell{objects: v, idx: idx, cellType: cellType}
}

func (c *BasicCell) PlacedObjects() *vector.Vector {
	return c.objects
}

func (c *BasicCell) SetPlacedObjects(p *vector.Vector) {
	c.objects = p
}

func (c *BasicCell) AddPlacedObject(p PlacedObject) {
	c.objects.Push(p)
}

func (c *BasicCell) RemovePlacedObject(p PlacedObject) {
	for i := 0; i < c.objects.Len(); i++ {
		if c.objects.At(i) == p {
			c.objects.Delete(i)
			break
		}
	}
}

func (c *BasicCell) Idx() Idx {
	return c.idx
}

func (c *BasicCell) SetIdx(idx Idx) {
	c.idx = idx
}

func (c *BasicCell) CellType() CellType {
	return c.cellType
}
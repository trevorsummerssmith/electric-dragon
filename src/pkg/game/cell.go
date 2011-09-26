package game

import (
	"container/vector"
)

type PlacedObject interface {
	Idx() Idx
	SetIdx(idx Idx)
}

type Cell interface {
	PlacedObjects() *vector.Vector
	SetPlacedObjects(p *vector.Vector)
	AddPlacedObject(p PlacedObject)
	RemovePlacedObject(p PlacedObject)
	Idx() Idx
	SetIdx(idx Idx)
}

type BasicCell struct {
	objects *vector.Vector
	idx Idx
}

func NewBasicCell(idx Idx) *BasicCell {
	v := new(vector.Vector)
	return &BasicCell{objects: v, idx: idx}
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
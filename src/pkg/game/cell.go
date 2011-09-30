package game

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
	PlacedObjects() []PlacedObject
	SetPlacedObjects(ps []PlacedObject)
	AddPlacedObject(p PlacedObject)
	RemovePlacedObject(p PlacedObject)

	Idx() Idx
	SetIdx(idx Idx)
	CellType() CellType
}

type BasicCell struct {
	placedObjects []PlacedObject
	idx Idx
	cellType CellType
}

func NewBasicCell(idx Idx, cellType CellType) *BasicCell {
	ps := make([]PlacedObject, 0)
	return &BasicCell{placedObjects: ps, idx: idx, cellType: cellType}
}

func (c *BasicCell) PlacedObjects() []PlacedObject {
	return c.placedObjects
}

func (c *BasicCell) SetPlacedObjects(ps []PlacedObject) {
	c.placedObjects = ps
}

func (c *BasicCell) AddPlacedObject(p PlacedObject) {
	c.placedObjects = append(c.placedObjects, p)
}

func (c *BasicCell) RemovePlacedObject(p PlacedObject) {
	for i := 0; i < len(c.placedObjects); i++ {
		if c.placedObjects[i] == p {
			c.placedObjects = append(c.placedObjects[:i],
				                 c.placedObjects[i+1:]...)
			// Optimization: bail after we found a guy.
			return
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
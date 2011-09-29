package game

type Action interface {
}

type Move struct {
	RelativeIndices []RelativeIdx
}

type Attack struct {
}

// Physics -> Character Param bundles
type ActionParams interface {
	Action() Action
}

type MoveParams struct {
	Move *Move
	CellType CellType
}

type AttackParams struct {
	Attack *Attack
	CellType CellType
	Characters []Character
}

func (mp *MoveParams) Action() Action {
	return mp.Move
}

func (ap *AttackParams) Action() Action {
	return ap.Attack
}
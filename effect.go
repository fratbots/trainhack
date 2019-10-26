package main

type EffectTile struct {
	Position Position
	Rune     rune
}

type Effect interface {
	Alive() bool
	Render() []EffectTile
	SetPosition(Position)
}

type EffectAura struct {
	position Position
}

func NewEffectAura() *EffectAura {
	return &EffectAura{}
}

func (e *EffectAura) Alive() bool {
	return true
}

func (e *EffectAura) SetPosition(position Position) {
	e.position = position
}

func (e *EffectAura) Render() []EffectTile {
	return []EffectTile{
		EffectTile{
			Position: Position{e.position.X - 1, e.position.Y},
			Rune:     '-',
		},
		EffectTile{
			Position: Position{e.position.X + 1, e.position.Y},
			Rune:     '-',
		},
		EffectTile{
			Position: Position{e.position.X, e.position.Y - 1},
			Rune:     '|',
		},
		EffectTile{
			Position: Position{e.position.X, e.position.Y + 1},
			Rune:     '|',
		},
	}
}

package main

type EffectTile struct {
	Position Position
	Rune     rune
}

type Effect interface {
	Update() bool
	Render() []EffectTile
	SetPosition(Position)
}

type EffectAura struct {
	position Position
	age      int
}

func NewEffectAura(age int) *EffectAura {
	return &EffectAura{
		age: age,
	}
}

func (e *EffectAura) Update() bool {
	e.age = e.age - 1
	if e.age <= 0 {
		return false
	}
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

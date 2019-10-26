package main

type EffectTile struct {
	Position Position
	Rune     rune
}

type Effect interface {
	Update() bool
	Render() []EffectTile
	GetBehavior() EffectBehavior
	Equals(Effect) bool
}

type EffectAura struct {
	behavior EffectBehavior
	age      int
	target   *Actor
}

func NewEffectAura(behavior EffectBehavior, age int, target *Actor) *EffectAura {
	return &EffectAura{
		behavior: behavior,
		age:      age,
		target:   target,
	}
}

func (e *EffectAura) GetBehavior() EffectBehavior {
	return e.behavior
}

func (e *EffectAura) GetTarget() *Actor {
	return e.target
}

func (e *EffectAura) Equals(other Effect) bool {
	if otherEffectAura, ok := other.(*EffectAura); ok {
		if e.GetTarget() == otherEffectAura.GetTarget() {
			return true
		}
	}
	return false
}

func (e *EffectAura) Update() bool {
	e.age = e.age - 1
	if e.age <= 0 {
		return false
	}
	return true
}

func (e *EffectAura) Render() []EffectTile {
	return []EffectTile{
		EffectTile{
			Position: Position{e.target.Position.X - 1, e.target.Position.Y},
			Rune:     '-',
		},
		EffectTile{
			Position: Position{e.target.Position.X + 1, e.target.Position.Y},
			Rune:     '-',
		},
		EffectTile{
			Position: Position{e.target.Position.X, e.target.Position.Y - 1},
			Rune:     '|',
		},
		EffectTile{
			Position: Position{e.target.Position.X, e.target.Position.Y + 1},
			Rune:     '|',
		},
	}
}

package main

// EffectAura is an Aura effect around actor.
type EffectAura struct {
	behavior EffectBehavior
	age      int
	target   *Actor
}

// NewEffectAura creates Aura effect.
func NewEffectAura(behavior EffectBehavior, age int, target *Actor) *EffectAura {
	return &EffectAura{
		behavior: behavior,
		age:      age,
		target:   target,
	}
}

// GetBehavior returns Aura effect behavior.
func (e *EffectAura) GetBehavior() EffectBehavior {
	return e.behavior
}

// GetTarget returns actor that Aura effect assigned to.
func (e *EffectAura) GetTarget() *Actor {
	return e.target
}

// Equals checks whether other effect is also Aura effect that is assigned to the same actor.
func (e *EffectAura) Equals(other Effect) bool {
	if otherEffectAura, ok := other.(*EffectAura); ok {
		if e.GetTarget() == otherEffectAura.GetTarget() {
			return true
		}
	}
	return false
}

// Update moves forward effect animation progress and returns false if animation has ended.
func (e *EffectAura) Update() bool {
	e.age = e.age - 1
	if e.age <= 0 {
		return false
	}
	return true
}

func (e *EffectAura) Render() []EffectTile {
	var runeLeft rune
	var runeRight rune
	var runeTop rune
	var runeBottom rune

	if e.age <= 6 {
		runeLeft = '+'
		runeRight = '+'
		runeTop = '+'
		runeBottom = '+'
	} else {
		runeLeft = '*'
		runeRight = '*'
		runeTop = '*'
		runeBottom = '*'
	}

	return []EffectTile{
		EffectTile{
			Position: Position{e.target.Position.X - 1, e.target.Position.Y},
			Rune:     runeLeft,
		},
		EffectTile{
			Position: Position{e.target.Position.X + 1, e.target.Position.Y},
			Rune:     runeRight,
		},
		EffectTile{
			Position: Position{e.target.Position.X, e.target.Position.Y - 1},
			Rune:     runeTop,
		},
		EffectTile{
			Position: Position{e.target.Position.X, e.target.Position.Y + 1},
			Rune:     runeBottom,
		},
	}
}

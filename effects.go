package main

// Effects is a manageable collection of currently active effects.
type Effects struct {
	effects []LabelledEffect
}

type LabelledEffect struct {
	Label  string
	Effect Effect
}

// Set creates effect with specified label. It replaces existing effects in case of labels collision.
func (e *Effects) Set(label string, effect Effect) {
	for idx, existingEffect := range e.effects {
		if existingEffect.Label == label {
			e.effects[idx].Effect = effect
			return
		}
	}
	e.effects = append(
		e.effects,
		LabelledEffect{
			Label:  label,
			Effect: effect,
		},
	)
}

// Update updates the state of all effects in collection and removes completed effects.
func (e *Effects) Update() {
	var updated []LabelledEffect
	for _, effect := range e.effects {
		if effect.Effect.Update() {
			updated = append(updated, effect)
		}
	}
	e.effects = nil
	e.effects = updated
}

// Count returns the number of active effects.
func (e *Effects) Count() int {
	return len(e.effects)
}

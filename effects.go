package main

// Effects is a manageable collection of currently active effects.
type Effects struct {
	effects []Effect
}

// Add inserts an effect taking into account it's adding behavior.
func (e *Effects) Add(effect Effect) {
	switch effect.GetBehavior().GetAddBehavior() {
	case EffectBehaviorAddOnce:
		for _, existingEffect := range e.effects {
			if existingEffect.Equals(effect) {
				return
			}
		}
		e.effects = append(e.effects, effect)
	case EffectBehaviorAddRefresh:
		for idx, existingEffect := range e.effects {
			if existingEffect.Equals(effect) {
				e.effects[idx] = effect
				return
			}
		}
		e.effects = append(e.effects, effect)
	case EffectBehaviorAddMultiple:
		e.effects = append(e.effects, effect)
	}
}

// Update updates the state of all effects in collection and removes completed effects.
func (e *Effects) Update() {
	var updated []Effect
	for _, effect := range e.effects {
		if effect.Update() {
			updated = append(updated, effect)
		}
	}
	e.effects = nil
	e.effects = updated
}

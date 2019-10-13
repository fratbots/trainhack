package main

type Energy struct {
	Value float64
}

func (e *Energy) Gain(speed float64) bool {
	result := e.CanTakeTurn()
	if !result {
		e.Value += speed * energyGain
		return e.CanTakeTurn()
	}
	return result
}

func (e *Energy) Spend() {
	if e.Value >= energyAction {
		e.Value -= energyAction
	}
}

func (e *Energy) CanTakeTurn() bool {
	return e.Value >= energyAction
}

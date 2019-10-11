package main

const (
	energyGain   = 100
	energyAction = 300
)

type Energy struct {
	Value float64
}

func (e *Energy) Gain(speed float64) bool {
	e.Value += speed * energyGain
	return e.CanTakeTurn()
}

func (e *Energy) Spend() {
	if e.Value >= energyAction {
		e.Value -= energyAction
	}
}

func (e *Energy) CanTakeTurn() bool {
	return e.Value >= energyAction
}

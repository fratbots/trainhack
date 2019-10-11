package main

type Behavior func() *Action

type Actor struct {
	IsHero   bool
	Position Position
	Energy   Energy
	Speed    float64
	Behavior Behavior
}

func (h *Actor) nextAction(action Action) {
	h.Behavior = func() *Action {
		return &action
	}
}

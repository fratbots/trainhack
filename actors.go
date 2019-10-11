package main

type Actor struct {
	IsHero   bool
	Position Position
	Energy   Energy
	Speed    float64
	Behavior func() Action
}

func (h *Actor) nextAction(action Action) {
	h.Behavior = func() Action {
		return action
	}
}

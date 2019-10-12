package main

type Actor struct {
	IsHero   bool
	Rune     rune
	Position Position
	Energy   Energy
	Speed    float64
	Behavior func() *Action
}

func (h *Actor) NextAction(action *Action) {
	h.Behavior = func() *Action {
		h.Behavior = nil
		return action
	}
}

func NewHero() *Actor {
	return &Actor{
		IsHero:   true,
		Rune:     '@',
		Position: Position{0, 0},
		Energy:   Energy{Value: energyAction},
		Speed:    1,
		Behavior: nil,
	}
}

func NewActor(pos Position, speed float64) *Actor {
	return &Actor{
		IsHero:   false,
		Rune:     '§',
		Position: pos,
		Energy:   Energy{Value: 0},
		Speed:    speed,
		Behavior: nil,
	}
}

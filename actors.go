package main

type Actor struct {
	IsHero   bool
	Rune     rune
	Position Position
	Energy   Energy
	Speed    float64

	Behavior    Behavior
	Interaction Interaction
}

type Behavior func() *Action

type Interaction func(actor *Actor) *Action

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

func NewActor(pos Position, speed float64, rune rune) *Actor {
	return &Actor{
		IsHero:   false,
		Rune:     rune,
		Position: pos,
		Energy:   Energy{Value: 0},
		Speed:    speed,
		Behavior: nil,
	}
}

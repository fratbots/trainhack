package main

type Class struct {
	Name   string
	IsHero bool
	Rune   rune
	Speed  float64

	Behavior    func(actor *Actor, stage *Stage) Behavior
	Interaction func(actor *Actor, stage *Stage) Interaction
}

const (
	ClassHero   = "hero"
	ClassPursue = "pursue"
)

var Classes = map[string]Class{

	ClassHero: {
		Name:     ClassHero,
		IsHero:   true,
		Rune:     '@',
		Speed:    1,
		Behavior: nil,
	},

	ClassPursue: {
		Name:  ClassPursue,
		Rune:  '@',
		Speed: 0.3,
		Behavior: func(actor *Actor, stage *Stage) Behavior {
			return BehaviorPursue2(actor, stage, stage.Hero)
		},
	},
}

package main

type Class struct {
	Name   string
	IsHero bool
	Rune   rune
	Speed  float64

	// creates behavior by default
	BehaviorInit func(stage *Stage, actor *Actor) Behavior

	// creates interaction by default
	InteractionInit func(stage *Stage, actor *Actor) Interaction
}

const (
	ClassHero        = "hero"
	ClassPursue      = "pursue"
	ClassDialog      = "dialog"
	ClassBattle      = "battle"
	ClassThink       = "think"
	ClassMirrorDown  = "mirrorDown"
	ClassMirrorUp    = "mirrorUp"
	ClassMirrorLeft  = "mirrorLeft"
	ClassMirrorRight = "mirrorRight"
)

func Classes() map[string]Class {

	return map[string]Class{

		ClassHero: {
			Name:   ClassHero,
			IsHero: true,
			Rune:   '@',
			Speed:  1,
		},

		ClassThink: {
			Name:  ClassThink,
			Rune:  'd',
			Speed: 1,
			BehaviorInit: func(stage *Stage, actor *Actor) Behavior {
				return BehaviorThink(stage, actor, '↑', '↓', '←', '→')
			},
		},

		ClassMirrorDown: {
			Name:  ClassMirrorDown,
			Rune:  '↓',
			Speed: 0,
		},

		ClassMirrorLeft: {
			Name:  ClassMirrorLeft,
			Rune:  '←',
			Speed: 0,
		},

		ClassMirrorRight: {
			Name:  ClassMirrorRight,
			Rune:  '→',
			Speed: 0,
		},

		ClassMirrorUp: {
			Name:  ClassMirrorDown,
			Rune:  '↑',
			Speed: 0,
		},

		ClassPursue: {
			Name:  ClassPursue,
			Rune:  '$',
			Speed: 0.3,
			BehaviorInit: func(stage *Stage, actor *Actor) Behavior {
				return BehaviorPursue(stage, actor, stage.Hero)
			},
		},

		ClassDialog: {
			Name:  ClassDialog,
			Rune:  'D',
			Speed: 0,
			InteractionInit: func(stage *Stage, actor *Actor) Interaction {
				return BehaviorDialog(stage, actor)
			},
		},

		ClassBattle: {
			Name:  ClassBattle,
			Rune:  'B',
			Speed: 0,
			InteractionInit: func(stage *Stage, actor *Actor) Interaction {
				return BehaviorBattle(stage, actor)
			},
		},
	}
}

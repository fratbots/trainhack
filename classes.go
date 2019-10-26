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
	ClassHero   = "hero"
	ClassPursue = "pursue"
	ClassDialog = "dialog"
	ClassBattle = "battle"
)

var Classes = map[string]Class{

	ClassHero: {
		Name:   ClassHero,
		IsHero: true,
		Rune:   '@',
		Speed:  1,
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
			// extract interaction function
			return func(target *Actor) *Action {
				return &Action{
					Actor: actor,
					Perform: func() Result {
						return AlternativeAction(&Action{
							Actor: actor,
							Perform: func() Result {
								stage.Game.Sound.SetTheme(SoundThemePursuit)
								stage.Game.SetScreen(NewDialogScreen("a_dialog", 0, NewScreenStage(stage.Game, "map2", nil)))
								return Result{}
							},
						}, false)
					},
				}
			}
		},
	},

	ClassBattle: {
		Name:  ClassBattle,
		Rune:  'B',
		Speed: 0,
		InteractionInit: func(stage *Stage, actor *Actor) Interaction {
			// extract interaction function
			return func(target *Actor) *Action {
				return &Action{
					Actor: actor,
					Perform: func() Result {
						return AlternativeAction(&Action{
							Actor: actor,
							Perform: func() Result {
								stage.Game.Sound.SetTheme(SoundThemePursuit)
								stage.Game.SetScreen(NewBattleScreen(stage.Hero, target, NewScreenStage(stage.Game, "map2", nil), nil))
								return Result{}
							},
						}, false)
					},
				}
			}
		},
	},
}

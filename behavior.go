package main

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func BehaviorPursue(stage *Stage, actor *Actor, target *Actor) Behavior {
	return func() *Action {
		dx := target.Position.X - actor.Position.X
		dy := target.Position.Y - actor.Position.Y
		dxAbs := abs(dx)
		dyAbs := abs(dy)

		// do not follow at 5 distance
		if dxAbs > 5 || dyAbs > 5 {
			return nil
		}

		// Show aura effect around Hero when he is being pursued.
		heroAuraLabel := "hero_aura"
		auraEffect := NewEffectAura(16, stage.Hero)
		stage.Effects.Set(heroAuraLabel, auraEffect)

		if dxAbs > dyAbs {
			if dx < 0 {
				return ActionMove(stage, actor, DirectionLeft)
			} else if dy > 0 {
				return ActionMove(stage, actor, DirectionRight)
			}
		} else {
			if dy < 0 {
				return ActionMove(stage, actor, DirectionTop)
			} else if dy > 0 {
				return ActionMove(stage, actor, DirectionDown)
			}
		}

		return nil
	}
}

func BehaviorThink(stage *Stage, actor *Actor, mirrorUp, mirrorDown, mirrorLeft, mirrorRight rune) Behavior {
	endRune := '4'
	return func() *Action {
		switch actor.Direction {
		case DirectionTop:
			findPos := Position{actor.Position.X, actor.Position.Y - 1}
			if value, ok := stage.RuneCoords[findPos]; ok {
				switch value {
				case mirrorLeft:
					actor.Direction = DirectionLeft
					return ActionMove(stage, actor, DirectionLeft)
				case mirrorRight:
					actor.Direction = DirectionRight
					return ActionMove(stage, actor, DirectionRight)
				case endRune:
					actor.MustBeDeleted = true
					return nil
				}
			}
		case DirectionDown:
			findPos := Position{actor.Position.X, actor.Position.Y + 1}
			if value, ok := stage.RuneCoords[findPos]; ok {
				switch value {
				case mirrorLeft:
					actor.Direction = DirectionLeft
					return ActionMove(stage, actor, DirectionLeft)
				case mirrorRight:
					actor.Direction = DirectionRight
					return ActionMove(stage, actor, DirectionRight)
				case endRune:
					actor.MustBeDeleted = true
					return nil
				}
			}
		case DirectionLeft:
			findPos := Position{actor.Position.X - 1, actor.Position.Y}
			if value, ok := stage.RuneCoords[findPos]; ok {
				switch value {
				case mirrorUp:
					actor.Direction = DirectionTop
					return ActionMove(stage, actor, DirectionTop)
				case mirrorDown:
					actor.Direction = DirectionDown
					return ActionMove(stage, actor, DirectionDown)
				case endRune:
					actor.MustBeDeleted = true
					return nil
				}
			}
		case DirectionRight:
			findPos := Position{actor.Position.X + 1, actor.Position.Y}
			if value, ok := stage.RuneCoords[findPos]; ok {
				switch value {
				case mirrorUp:
					actor.Direction = DirectionTop
					return ActionMove(stage, actor, DirectionTop)
				case mirrorDown:
					actor.Direction = DirectionDown
					return ActionMove(stage, actor, DirectionDown)
				case endRune:
					actor.MustBeDeleted = true
					return nil
				}
			}
		}

		return ActionMove(stage, actor, actor.Direction)
	}
}

func BehaviorGhost(actor *Actor, stage *Stage, target *Actor) *Actor {
	actor.Behavior = func() *Action {
		return &Action{
			Actor: actor,
			Perform: func() Result {
				pos := actor.Position.FollowGap(target.Position, 7)
				if pos.IsOn(stage.Level.Dimensions) {
					actor.Position = pos
					return UpdatedResult
				}

				return FailureResult
			},
		}
	}

	return actor
}

func BehaviorDialog(stage *Stage, actor *Actor) Interaction {
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
}

func BehaviorBattle(stage *Stage, actor *Actor) Interaction {
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
}

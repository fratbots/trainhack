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

func BehaviorThink(stage *Stage, actor *Actor, mirrors []*Actor) Behavior {
	return func() *Action {

		return ActionMove(stage, actor, DirectionLeft)
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

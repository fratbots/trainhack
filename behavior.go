package main

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Pursue(actor *Actor, stage *Stage, target *Actor) *Actor {

	actor.Behavior = func() Action {
		if abs(target.Position.X-actor.Position.X) > abs(target.Position.Y-actor.Position.Y) {
			if target.Position.X-actor.Position.X < 0 {
				return ActionMove(stage, actor, DirectionLeft)
			} else {
				return ActionMove(stage, actor, DirectionRight)
			}
		} else {
			if target.Position.Y-actor.Position.Y < 0 {
				return ActionMove(stage, actor, DirectionTop)
			} else {
				return ActionMove(stage, actor, DirectionDown)
			}
		}
	}

	return actor
}

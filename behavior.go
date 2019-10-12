package main

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Pursue(actor *Actor, stage *Stage, target *Actor) *Actor {

	actor.Behavior = func() *Action {
		xd := target.Position.X - actor.Position.X
		yd := target.Position.Y - actor.Position.Y
		absXd := abs(xd)
		absYd := abs(yd)

		if absXd > 5 || absYd > 5 {
			return nil
		}

		if absXd > absYd {
			if yd < 0 {
				return ActionMove(stage, actor, DirectionLeft)
			} else {
				return ActionMove(stage, actor, DirectionRight)
			}
		} else {
			if yd < 0 {
				return ActionMove(stage, actor, DirectionTop)
			} else {
				return ActionMove(stage, actor, DirectionDown)
			}
		}
	}

	return actor
}

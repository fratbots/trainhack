package main

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Pursue(stage *Stage, who, target *Actor) func() Action {
	return func() Action {
		if abs(target.Position.X-who.Position.X) > abs(target.Position.Y-who.Position.Y) {
			if target.Position.X-who.Position.X < 0 {
				return ActionMove(stage, who, DirectionLeft)
			} else {
				return ActionMove(stage, who, DirectionRight)
			}
		} else {
			if target.Position.Y-who.Position.Y < 0 {
				return ActionMove(stage, who, DirectionTop)
			} else {
				return ActionMove(stage, who, DirectionDown)
			}
		}
	}
}

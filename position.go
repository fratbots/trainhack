package main


type Vec2 struct {
	X, Y int
}

type Direction = Vec2

type Position = Vec2

type Dimensions = Vec2

func (p Position) Shift(d Direction) Position {
	return Position{p.X + d.X, p.Y + d.Y}
}

func (p Position) FollowGap(n Position, gap int) (result Position) {
	result = p

	dx := n.X - p.X
	if abs(dx) > gap {
		if dx < 0 {
			result.X = n.X + gap
		} else {
			result.X = n.X - gap
		}
	}

	dy := n.Y - p.Y
	if abs(dy) > gap {
		if dy < 0 {
			result.Y = n.Y + gap
		} else {
			result.Y = n.Y - gap
		}
	}

	return
}

func (p Position) IsOn(d Dimensions) bool {
	if p.X >= 0 && p.X < d.X &&
		p.Y >= 0 && p.Y < d.Y {
		return true
	}

	return false
}

var (
	DirectionTop   = Direction{X: 0, Y: -1}
	DirectionDown  = Direction{X: 0, Y: +1}
	DirectionLeft  = Direction{X: -1, Y: 0}
	DirectionRight = Direction{X: +1, Y: 0}
)

package main

import (
	"container/list"
)

type Action struct {
	Actor   *Actor
	Perform func() Result
}

type Result struct {
	Success     bool
	Alternative *Action
}

type Actions struct {
	list *list.List
}

func NewActions() *Actions {
	return &Actions{list: list.New()}
}

func (a *Actions) Add(action *Action) {
	a.list.PushBack(action)
}

func (a *Actions) Get() *Action {
	f := a.list.Front()
	if f != nil {
		a.list.Remove(f)
		return f.Value.(*Action)
	}
	return nil
}

func (a *Actions) Reset() {
	for {
		e := a.list.Front()
		if e == nil {
			return
		}
		a.list.Remove(e)
	}
}

// ------------------------------------------------------------ //

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

	successResult = Result{
		Success:     true,
		Alternative: nil,
	}
)

// ============================================================ //

func ActionMove(stage *Stage, actor *Actor, dir Direction) *Action {
	return &Action{
		Actor: actor,
		Perform: func() Result {

			pos := actor.Position.Shift(dir)

			target := stage.ActorAt(pos)

			if target != nil {
				return successResult // rest
			}

			if !pos.IsOn(stage.LevelMap.Dimensions) {
				return successResult // rest, TODO: hit
			}

			actor.Position = pos
			return successResult
		},
	}
}

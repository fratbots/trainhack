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

// ------------------------------------------------------------ //

type Vec2 struct {
	X, Y int
}

type Direction = Vec2

type Position = Vec2

func (p Position) Shift(d Direction) Position {
	return Position{p.X + d.X, p.Y + d.Y}
}

var (
	DirectionTop   = Direction{X: 0, Y: -1}
	DirectionDown  = Direction{X: 0, Y: +1}
	DirectionLeft  = Direction{X: -1, Y: 0}
	DirectionRight = Direction{X: +1, Y: 0}
)

func success() Result {
	return Result{
		Success:     true,
		Alternative: nil,
	}
}

func alternate(alt *Action) Result {
	return Result{
		Success:     true,
		Alternative: alt,
	}
}

// ============================================================ //

func ActionMove(stage *Stage, actor *Actor, dir Direction) *Action {
	return &Action{
		Actor: actor,
		Perform: func() Result {

			pos := actor.Position.Shift(dir)

			target := stage.ActorAt(pos)

			if target != nil {
				return success() // rest
			}

			// TODO: collision

			actor.Position = pos
			return success()
		},
	}
}

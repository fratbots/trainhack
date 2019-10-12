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

var (
	successResult = Result{
		Success:     true,
		Alternative: nil,
	}

	failureResult = Result{
		Success:     false,
		Alternative: nil,
	}
)

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

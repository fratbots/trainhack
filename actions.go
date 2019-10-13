package main

import (
	"container/list"
	"sync"
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
	mu   sync.Mutex
}

func NewActions() *Actions {
	return &Actions{list: list.New()}
}

func (a *Actions) Add(action *Action) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.list.PushBack(action)
}

func (a *Actions) Get() *Action {
	a.mu.Lock()
	defer a.mu.Unlock()

	f := a.list.Front()
	if f != nil {
		a.list.Remove(f)
		return f.Value.(*Action)
	}
	return nil
}

func (a *Actions) Reset() {
	a.mu.Lock()
	defer a.mu.Unlock()

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

func alternativeAction(action *Action) Result {
	return Result{
		Success:     false,
		Alternative: action,
	}
}

func ActionMove(stage *Stage, actor *Actor, dir Direction) *Action {
	return &Action{
		Actor: actor,
		Perform: func() Result {

			pos := actor.Position.Shift(dir)

			target := stage.ActorAt(pos)
			if target != nil {
				// target interacts to actor
				if target.Interaction != nil {
					return alternativeAction(target.Interaction(actor))
				}

				return successResult // rest
			}

			if !pos.IsOn(stage.Level.Dimensions) {
				return successResult // rest, TODO: hit
			}

			tile := stage.Level.GetTile(pos)
			if !tile.IsWalkable {
				return successResult // rest
			}

			actor.Position = pos

			if tile.Interaction != nil {
				stage.Actions.Add(tile.Interaction(actor))
			}

			return successResult
		},
	}
}

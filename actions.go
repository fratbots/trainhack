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
	Deferred    bool
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

func (a *Actions) Len() int {
	return a.list.Len()
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
	SuccessResult = Result{
		Success:     true,
		Alternative: nil,
	}

	FailureResult = Result{
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

				return SuccessResult // rest
			}

			if !pos.IsOn(stage.Level.Dimensions) {
				return SuccessResult // rest, TODO: hit
			}

			tile := stage.Level.GetTile(pos)
			if !tile.IsWalkable {
				return SuccessResult // rest
			}

			actor.Position = pos

			if tile.Interaction != nil {
				return alternativeAction(tile.Interaction(actor))
				// stage.Actions.Add(tile.Interaction(actor))
			}

			return SuccessResult
		},
	}
}

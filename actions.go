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
	Updated               bool    // true - need to draw and etc.
	Alternative           *Action // alternative action
	AlternativeIsDeferred bool    // Alternative is for next update
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
	UpdatedResult = Result{
		Updated:               true,
		Alternative:           nil,
		AlternativeIsDeferred: false,
	}

	FailureResult = Result{
		Updated:     false,
		Alternative: nil,
	}
)

func AlternativeAction(action *Action, updated, deferred bool) Result {
	return Result{
		Updated:               updated,
		Alternative:           action,
		AlternativeIsDeferred: deferred,
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
					return AlternativeAction(target.Interaction(actor), false, false)
				}

				return UpdatedResult // rest
			}

			if !pos.IsOn(stage.Level.Dimensions) {
				return UpdatedResult // rest, TODO: hit
			}

			tile := stage.Level.GetTile(pos)
			if !tile.IsWalkable {
				return UpdatedResult // rest
			}

			actor.Position = pos

			if tile.Interaction != nil {
				return AlternativeAction(tile.Interaction(actor), false, false)
				// stage.Actions.Add(tile.Interaction(actor))
			}

			return UpdatedResult
		},
	}
}

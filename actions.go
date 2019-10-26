package main

import (
	"container/list"
	"sync"
)

type Action struct {
	Actor    *Actor
	Deferred bool
	Perform  func() Result
}

type Result struct {
	Updated     bool    // need to draw and etc.
	Alternative *Action // alternative action
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
		Updated:     true,
		Alternative: nil,
	}

	FailureResult = Result{
		Updated:     false,
		Alternative: nil,
	}
)

func AlternativeAction(action *Action, updated bool) Result {
	return Result{
		Updated:     updated,
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
					return AlternativeAction(target.Interaction(actor), false)
				}

				return UpdatedResult // rest
			}

			if !pos.IsOn(stage.Level.Dimensions) {
				return UpdatedResult // rest, TODO: hit
			}

			tile := stage.Level.GetTile(pos)
			if !tile.GetWalkable() {
				return UpdatedResult // rest
			}

			actor.Position = pos

			interaction := tile.GetInteraction()
			if interaction != nil {
				return AlternativeAction(interaction(actor), true)
			}

			return UpdatedResult
		},
	}
}

func ActionInteract(stage *Stage, actor *Actor) *Action {
	return &Action{
		Actor: actor,
		Perform: func() Result {

			tile := stage.Level.GetTile(stage.Hero.Position)
			interaction := tile.GetInteraction()
			if interaction != nil {
				return AlternativeAction(interaction(actor), false)
			}

			return UpdatedResult
		},
	}
}

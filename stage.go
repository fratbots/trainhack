package main

import (
	"time"
)

const (
	tickTimeout  = time.Second / 5
	tickTimeoutF = float64(tickTimeout)
)

type Stage struct {
	Game *Game

	Hero    *Actor
	Actors  []*Actor
	Actions *Actions

	ticker *Ticker
}

func NewStage(g *Game) *Stage {
	hero := NewHero()
	return &Stage{
		Game: g,
		Hero: hero,

		Actors:  []*Actor{hero},
		Actions: NewActions(),
	}
}

func (s *Stage) StartTime() {
	s.ticker = NewTicker(tickTimeout, func(d time.Duration) {
		if s.Update(d) {
			s.Game.View.Draw()
		}
	})
}

func (s *Stage) StopTime() {
	if s.ticker != nil {
		s.ticker.Done()
	}
}

func (s *Stage) ActorAt(pos Position) *Actor {
	for _, a := range s.Actors {
		if a.Position == pos {
			return a
		}
	}

	return nil
}

func (s *Stage) AddActor(actor *Actor) {
	s.Actors = append(s.Actors, actor)
}

func (s *Stage) Update(d time.Duration) bool {

	if d > tickTimeout {
		d = tickTimeout
	}
	timeFactor := float64(d) / tickTimeoutF

	l := len(s.Actors)
	for i := 0; i < l; i++ {
		actor := s.Actors[i]
		if actor.Behavior == nil {
			continue
		}

		if actor.Energy.CanTakeTurn() || actor.Energy.Gain(timeFactor*actor.Speed) {
			action := actor.Behavior()
			if action != nil {
				s.Actions.Add(action)
			}
		}
	}

	needToDraw := false

	for {
		action := s.Actions.Get()
		if action == nil {
			break
		}

		result := action.Perform()

		for result.Alternative != nil {
			result = result.Alternative.Perform()
		}

		if result.Success {
			needToDraw = true

			if action.Actor != nil {
				action.Actor.Energy.Spend()
			}
		}
	}

	return needToDraw
}

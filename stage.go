package main

import (
	"time"
)

type Stage struct {
	Name string

	Game    *Game
	Hero    *Actor
	Actors  []*Actor
	Actions *Actions
	Level   *Level

	Effects Effects

	UpdateCallback func(d time.Duration) bool

	deferred *Actions
	ticker   *Ticker
	frame    int

	RuneCoords map[Position]rune
}

func (s *Stage) RegisterRune(name rune, coords Position) {
	s.RuneCoords[coords] = name
}

func NewStage(g *Game) *Stage {
	hero := NewHero()
	return &Stage{
		Game: g,
		Hero: hero,

		Actors:   []*Actor{hero},
		Actions:  NewActions(),
		deferred: NewActions(),
		frame:    0,

		RuneCoords: make(map[Position]rune),
	}
}

func (s *Stage) Load(name string, location *rune) *Stage {
	s.Stop()

	level := LoadLevel(s.Game, name)
	if level == nil {
		panic("cannot load level " + name)
		// TODO: handle error
		return s
	}

	// reset
	s.Name = name
	s.Level = level
	s.Actions.Reset()
	s.deferred.Reset()

	// load state
	if !LoadState(s) {
		// TODO: create actors from levelMap
		s.Hero = NewHero()
		s.Hero.Position = Position{X: 39, Y: 5}
		s.Actors = []*Actor{s.Hero}
	}

	if location != nil {
		if pos, ok := level.Doors[*location]; ok {
			s.Hero.Position = pos
		}
	}

	return s
}

func (s *Stage) Start() {
	s.ticker = NewTicker(tickTimeout, func(d time.Duration) {
		if s.Update(d) {
			s.Game.Draw()
		}
	})
}

func (s *Stage) Stop() {
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
	s.frame = s.frame + 1

	if d > tickTimeout {
		d = tickTimeout
	}
	timeFactor := float64(d) / tickTimeoutFloat

	l := len(s.Actors)

	for i := 0; i < l; i++ {

		actor := s.Actors[i]
		if actor == nil {
			continue
		}
		if actor.MustBeDeleted {
			// FIXME
			//s.Actors[i] = nil
			continue
		}
		if actor.Behavior == nil {
			continue
		}

		if actor.Energy.CanTakeTurn() || actor.Energy.Gain(timeFactor*actor.Class.Speed) {
			action := actor.Behavior()
			if action != nil {
				if action.Deferred {
					s.deferred.Add(action)
				} else {
					s.Actions.Add(action)
				}
			}
		}
	}
	needToDraw := false

	for {
		action := s.Actions.Get()
		if action == nil {
			break
		}

		if action.Deferred {
			s.deferred.Add(action)
			continue
		}

		result := action.Perform()

		for result.Alternative != nil {
			if result.Alternative.Deferred {
				s.deferred.Add(result.Alternative)
				break
			}

			result = result.Alternative.Perform()
		}

		if result.Updated {
			needToDraw = true

			if action.Actor != nil {
				action.Actor.Energy.Spend()
			}
		}
	}

	if s.UpdateCallback != nil {
		if s.UpdateCallback(d) {
			needToDraw = true
		}
	}

	for {
		a := s.deferred.Get()
		if a == nil {
			break
		}
		a.Deferred = false
		s.Actions.Add(a)
	}
	return needToDraw
}

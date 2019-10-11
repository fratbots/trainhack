package main

type Stage struct {
	Hero    *Actor
	Actors  []*Actor
	Actions *Actions
}

func NewStage(hero *Actor) *Stage {
	return &Stage{
		Hero:    hero,
		Actors:  []*Actor{hero},
		Actions: NewActions(),
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

func (s *Stage) Update() {
	for {
		action := s.Actions.Get()
		if action == nil {
			break
		}

		result := (*action)()

		for result.Alternative != nil {
			result = (*result.Alternative)()
		}
	}

	for i := 0; i < len(s.Actors); i++ {
		actor := s.Actors[i]
		if actor.Behavior == nil {
			continue
		}

		if actor.Energy.CanTakeTurn() || actor.Energy.Gain(actor.Speed) {
			action := actor.Behavior()
			if action != nil {
				s.Actions.Add(action)
			}
		}
	}
}

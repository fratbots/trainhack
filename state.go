package main

type State struct {
	Stages map[string]StateStage
}

type StateStage struct {
	Hero   ActorState
	Actors []ActorState
}

type ActorState struct {
	ClassName string
	Position  Position
	Energy    Energy
}

func SaveState(stage *Stage) {
	if stage.Name == "" {
		return
	}

	stateStage := StateStage{
		Hero:   GetActorState(stage.Hero),
		Actors: []ActorState{},
	}

	for _, actor := range stage.Actors {
		if actor != stage.Hero {
			stateStage.Actors = append(stateStage.Actors, GetActorState(actor))
		}
	}

	stage.Game.State.Stages[stage.Name] = stateStage
}

func LoadState(stage *Stage) bool {
	state, ok := stage.Game.State.Stages[stage.Name]
	if !ok {
		return false
	}

	hero := NewClassActor(stage, Position{}, Direction{}, state.Hero.ClassName)
	hero.Position = state.Hero.Position

	for _, a := range state.Actors {
		actor := NewClassActor(stage, Position{}, Direction{}, a.ClassName)
		actor.Position = a.Position
		stage.AddActor(actor)
	}

	return true
}

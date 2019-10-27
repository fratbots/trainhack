package main

import (
	"encoding/json"
	"io/ioutil"
)

type State struct {
	Stage  string
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

	stage.Game.State.Stage = stage.Name
	stage.Game.State.Stages[stage.Name] = stateStage
}

func LoadState(stage *Stage) bool {
	state, ok := stage.Game.State.Stages[stage.Name]
	if !ok {
		return false
	}

	stage.Hero.Position = state.Hero.Position
	stage.Hero.Energy = state.Hero.Energy

	for _, a := range state.Actors {
		actor := NewClassActor(stage, a.Position, Direction{}, a.ClassName)
		stage.AddActor(actor)
	}

	return true
}

func SaveToFile(state *State) {
	b, _ := json.Marshal(state)
	_ = ioutil.WriteFile("./progress.json", b, 0666)
}

func LoadFromFile(state *State) {
	r, _ := ioutil.ReadFile("./progress.json")
	_ = json.Unmarshal(r, state)
}

package main

type State struct {
	Stages map[string]StateStage
}

type StateStage struct {
	HeroPosition Position
	Actors       []*Actor
}

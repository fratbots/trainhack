package main

import (
	"fmt"
)

type Screen interface {
	Do(g *Game)
}

type HelloScreen struct {
}

func (hs *HelloScreen) Do(g *Game) {
	fmt.Printf("hello")
}

func init() {

}

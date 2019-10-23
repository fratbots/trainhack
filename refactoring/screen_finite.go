package main

import (
	"github.com/rivo/tview"
)

type finite struct{}

func NewFiniteScreen() Screen {
	return &finite{}
}

func (s *finite) Finalize() {}

func (s *finite) Init(game *Game) tview.Primitive {

	var modal *tview.Modal

	modal = tview.NewModal().
		SetText("Bye bye!").
		AddButtons([]string{"See Ya!"}).
		SetDoneFunc(
			func(buttonIndex int, buttonLabel string) {
				if buttonIndex == 0 {
					game.Finalize()
					return
				}
			},
		)

	return modal

}

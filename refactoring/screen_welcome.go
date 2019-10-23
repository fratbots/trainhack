package main

import (
	"github.com/rivo/tview"
)

type welcome struct{}

func NewWelcomeScreen() Screen {
	return &welcome{}
}

func (s *welcome) Finalize() {}

func (s *welcome) Init(game *Game) tview.Primitive {
	var modal *tview.Modal
	modal = tview.NewModal().
		SetText("Hello!").
		AddButtons([]string{"Play", "Exit"}).
		SetDoneFunc(
			func(buttonIndex int, buttonLabel string) {
				if buttonIndex == 0 {
					game.SetScreen(NewFiniteScreen())
					return
				}
				if buttonIndex == 1 {
					game.SetScreen(NewFiniteScreen())
					return
				}

				modal.SetText("You win!")
			},
		)

	return modal
}

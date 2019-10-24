package main

import (
	"github.com/rivo/tview"
)

func NewScreenFinal() Screen {
	return NewScreen(
		func(game *Game) tview.Primitive {
			return NewUIModal("The End!", "OK", func() {
				game.Finish()
			})
		},
	)
}

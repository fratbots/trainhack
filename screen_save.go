package main

import (
	"github.com/rivo/tview"
)

func NewScreenSave(currentMap string) Screen {
	return NewScreen(
		func(game *Game) tview.Primitive {
			return NewUIModal("-= Save your progress =-",
				"Save", func() {
					SaveToFile(&game.State)
					game.SetScreen(NewScreenStage(game, currentMap, nil))
				},
				"Load", func() {
					LoadFromFile(&game.State)
					game.SetScreen(NewScreenStage(game, game.State.Stage, nil))
				},
				"Cancel", func() {
					game.SetScreen(NewScreenStage(game, currentMap, nil))
				},
			)
		},
	)
}

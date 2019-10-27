package main

import (
	"github.com/rivo/tview"
)

func NewScreenSave(currentMap string) Screen {
	return NewScreen(
		func(game *Game) tview.Primitive {
			game.Sound.Pause()
			return NewUIModal("-= Save your progress =-",
				"Save", func() {
					SaveToFile(&game.State)
					game.Sound.SetTheme(SoundThemeAutumn)
					game.SetScreen(NewScreenStage(game, currentMap, nil))
				},
				"Load", func() {
					LoadFromFile(&game.State)
					game.Sound.SetTheme(SoundThemeAutumn)
					game.SetScreen(NewScreenStage(game, game.State.Stage, nil))
				},
				"Continue", func() {
					game.Sound.SetTheme(SoundThemeAutumn)
					game.SetScreen(NewScreenStage(game, currentMap, nil))
				},
				"Exit", func() {
					game.Sound.SetTheme(SoundThemePursuit)
					game.SetScreen(NewScreenFinal())
				},
			)
		},
	)
}

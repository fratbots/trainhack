package main

import (
	"github.com/rivo/tview"
)

type HelloScreen struct {
}

func (s *HelloScreen) Do(g *Game, end func(next Screen)) tview.Primitive {

	var modal *tview.Modal
	modal = tview.NewModal().
		SetText("Hello, чувак!").
		AddButtons([]string{"Play", "Exit"}).
		SetDoneFunc(
			func(buttonIndex int, buttonLabel string) {
				if buttonIndex == 0 {
					end(&ScreensStage{})
					return
				}
				if buttonIndex == 1 {
					end(&ScreenFinal{})
					return
				}

				modal.SetText("You win!")
			},
		)

	return modal
}

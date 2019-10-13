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
		AddButtons([]string{"Play", "Exit", "Dialog"}).
		SetDoneFunc(
			func(buttonIndex int, buttonLabel string) {
				if buttonIndex == 0 {
					end(NewScreenStage(g, "map2"))
					return
				}
				if buttonIndex == 1 {
					end(&ScreenFinal{})
					return
				}
				if buttonIndex == 2 {
					end(NewDialogScreen("pika_dialog1", 0, &ScreenFinal{}))
					return
				}

				modal.SetText("You win!")
			},
		)

	return modal
}

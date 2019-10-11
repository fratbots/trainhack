package main

import (
	"github.com/rivo/tview"
)

type HelloScreen struct {
}

func (s *HelloScreen) Do(g *Game, end func(next Screen)) tview.Primitive {

	modal := tview.NewModal().
		SetText("Hello, чувак!").
		AddButtons([]string{"Play", "Exit", "DIALOG"}).
		SetDoneFunc(
			func(buttonIndex int, buttonLabel string) {
				if buttonIndex == 0 {
					end(&ScreenFinal{})
					return
				} else if buttonLabel == "DIALOG" {
					end(NewDialogScreen("pika_dialog1", 0, &ScreenFinal{}))
					return
				}

				g.Final()
			},
		)

	return modal
}

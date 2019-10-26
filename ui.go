package main

import (
	"log"

	"github.com/rivo/tview"
)

type UI struct {
	app    *tview.Application
	screen Screen
}

func (u *UI) SetScreen(game *Game, screen Screen) {

	// start app
	if u.app == nil {

		u.screen = screen
		primitive := u.screen.Init(game)

		u.app = tview.NewApplication()
		u.app.SetRoot(primitive, true)

		err := u.app.Run()
		if err != nil {
			log.Fatalf("Failed to run app", err)
		}

		return
	}

	// update screen
	u.app.QueueUpdate(func() {
		if u.screen != nil {
			u.screen.Finalize()
		}

		u.screen = screen
		primitive := u.screen.Init(game)

		u.app.SetRoot(primitive, true)
		u.app.SetFocus(primitive)
		u.app.Draw()
	})
}

func (u *UI) Finalize() {
	if u.screen != nil {
		u.screen.Finalize()
	}
	if u.app != nil {
		u.app.Stop()
	}
}

func (u *UI) Draw() {
	if u.app != nil {
		u.app.Draw()
	}
}

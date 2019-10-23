package main

import (
	"github.com/rivo/tview"
)

type UI interface {
	SetScreen(*Game, Screen)
	Finalize()
}

type ui struct {
	app    *tview.Application
	screen Screen
}

func NewUI() UI {
	return &ui{}
}

func (u *ui) SetScreen(g *Game, screen Screen) {

	// start app
	if u.app == nil {

		u.screen = screen
		p := u.screen.Init(g)

		u.app = tview.NewApplication()
		u.app.SetRoot(p, true)

		err := u.app.Run()
		if err != nil {
			panic(err)
		}

		return
	}

	// update screen
	u.app.QueueUpdate(func() {
		if u.screen != nil {
			u.screen.Finalize()
		}

		u.screen = screen
		p := u.screen.Init(g)
		u.app.SetRoot(p, true)
		u.app.SetFocus(p)
		u.app.Draw()
	})
}

func (u *ui) Finalize() {
	if u.screen != nil {
		u.screen.Finalize()
	}
	if u.app != nil {
		u.app.Stop()
	}
}

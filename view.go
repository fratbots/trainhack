package main

import (
	"github.com/rivo/tview"
)

type View struct {
	App *tview.Application
}

func NewVew() *View {
	view := View{
		App: tview.NewApplication(),
	}

	return &view
}

func (v *View) Final() {
	v.App.Stop()
}

func (v *View) Run(p tview.Primitive) {
	// v.App.SetAfterDrawFunc() TODO: use to get flow
	v.App.SetRoot(p, true)
	if err := v.App.Run(); err != nil {
		panic(err)
	}
}

func (v *View) Focus(p tview.Primitive) {
	v.App.SetRoot(p, true)
	v.App.SetFocus(p)
	v.App.Draw()
}

func (v *View) Draw() {
	v.App.Draw()
}

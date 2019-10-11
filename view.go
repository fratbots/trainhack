package main

import (
	"github.com/rivo/tview"
)

type View struct {
	App   *tview.Application
	Pages *tview.Pages
}

func NewVew() *View {
	view := View{
		App:   tview.NewApplication(),
		Pages: tview.NewPages(),
	}

	view.App.SetRoot(view.Pages, true)
	return &view
}

func (v *View) Final() {
	v.App.Stop()
}

func (v *View) Page(name string) bool {
	if v.Pages.HasPage("hello") {
		v.Pages.SwitchToPage("hello")
		return true
	}

	return false
}

func (v *View) Run(p tview.Primitive) {
	v.App.SetRoot(p, true)
	if err := v.App.Run(); err != nil {
		panic(err)
	}
}

func (v *View) Focus(p tview.Primitive) {
	v.App.SetRoot(p, true)
	v.App.SetFocus(p)
}

func (v *View) Draw() {
	v.App.Draw()
}

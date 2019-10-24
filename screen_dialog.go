package main

import (
	"github.com/rivo/tview"
)

type DialogScreen struct {
	name       string
	sceneId    int
	lastScreen Screen
}

func (d *DialogScreen) Init(g *Game) tview.Primitive {
	callback := func(s Screen) {
		g.SetScreen(s)
	}
	primitive := g.Dialoger.GetDialogPrimitive(d.name, d.sceneId, callback, d.lastScreen)
	return primitive
}

func (d *DialogScreen) Finalize() {}

func NewDialogScreen(dialogName string, sceneId int, lastScreen Screen) Screen {
	return &DialogScreen{
		dialogName,
		sceneId,
		lastScreen,
	}
}

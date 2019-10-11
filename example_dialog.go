package main

import (
	"github.com/rivo/tview"
)

type DialogScreen struct {
	name       string
	sceneId    int
	lastScreen Screen
}

func (d *DialogScreen) Do(g *Game, callback func(s Screen)) tview.Primitive {
	primitive := g.Dialoger.GetDialogPrimitive(d.name, d.sceneId, callback, d.lastScreen)
	return primitive
}

func NewDialogScreen(dialogName string, sceneId int, lastScreen Screen) Screen {
	return &DialogScreen{
		dialogName,
		sceneId,
		lastScreen,
	}
}

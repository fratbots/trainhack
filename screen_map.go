package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type ScreenMap struct {
}

func (s *ScreenMap) Do(g *Game, end func(s Screen)) tview.Primitive {

	// Time & Life
	// Stage
	// Box
	// Input

	var box *tview.Box

	box = tview.NewBox().
		SetTitle("title").
		SetBorder(true).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

			end(nil)
			return nil
		}).
		SetBackgroundColor(tcell.ColorGreen).
		SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			x, y, width, height = box.GetInnerRect()

			screen.SetContent(x+(width/2), y+(height/2), '@', nil, tcell.StyleDefault)

			return x, y, width, height
		})

	return box
}


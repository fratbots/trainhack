package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type ScreenFinal struct {
}

func (s *ScreenFinal) Do(g *Game, callback func(next Screen)) tview.Primitive {
	text := tview.NewTextView().
		SetChangedFunc(func() {
			g.View.Draw()
		}).
		SetText("bye bye").
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetTitle("title").
		SetBorder(true).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			g.Final()
			return nil
		})

	return text
}

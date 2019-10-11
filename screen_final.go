package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type ScreenFinal struct {
}

func (s *ScreenFinal) Do(g *Game, callback func(next Screen)) tview.Primitive {
	text := tview.NewModal().
		SetText("Bye Bye!").
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			g.Final()
			return nil
		})

	return text
}

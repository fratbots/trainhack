package main

import (
	"github.com/rivo/tview"
)

type Screen interface {
	Init(game *Game) tview.Primitive
	Finalize()
}

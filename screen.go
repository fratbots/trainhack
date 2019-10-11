package main

import (
	"github.com/rivo/tview"
)

type Screen interface {
	Do(g *Game, end func(next Screen)) tview.Primitive
}

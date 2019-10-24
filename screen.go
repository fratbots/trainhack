package main

import (
	"github.com/rivo/tview"
)

type Screen interface {
	Init(game *Game) tview.Primitive
	Finalize()
}

func NewScreen(fn func(game *Game) tview.Primitive) Screen {
	return &screen{fn: fn}
}

type screen struct{ fn func(game *Game) tview.Primitive }

func (s *screen) Init(game *Game) tview.Primitive {
	return s.fn(game)
}

func (s *screen) Finalize() {}

package main

import "github.com/gdamore/tcell"

type Tile interface {
	GetAppearance(frame int, position Position) (rune, tcell.Style)
	GetWalkable() bool
	SetWalkable(bool)
	GetInteraction() Interaction
	SetInteraction(Interaction)
	Copy() Tile
}

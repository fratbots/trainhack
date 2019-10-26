package main

import "github.com/gdamore/tcell"

type TileAnimated struct {
	walkable    bool
	interaction Interaction
	tileRune    rune
	style       tcell.Style
}

func (t TileAnimated) GetAppearance(frame int, x int, y int) (rune, tcell.Style) {
	return t.tileRune, t.style
}

func (t TileAnimated) GetWalkable() bool {
	return t.walkable
}

func (t TileAnimated) SetWalkable(walkable bool) {
	t.walkable = walkable
}

func (t TileAnimated) GetInteraction() Interaction {
	return t.interaction
}

func (t TileAnimated) SetInteraction(interaction Interaction) {
	t.interaction = interaction
}

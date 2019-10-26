package main

import "github.com/gdamore/tcell"

type TileStatic struct {
	walkable    bool
	interaction Interaction
	tileRune    rune
	style       tcell.Style
}

func (t *TileStatic) GetAppearance(frame int, position Position) (rune, tcell.Style) {
	return t.tileRune, t.style
}

func (t *TileStatic) GetWalkable() bool {
	return t.walkable
}

func (t *TileStatic) SetWalkable(walkable bool) {
	t.walkable = walkable
}

func (t *TileStatic) GetInteraction() Interaction {
	return t.interaction
}

func (t *TileStatic) SetInteraction(interaction Interaction) {
	t.interaction = interaction
}

func (t *TileStatic) Copy() Tile {
	c := *t
	return &c
}

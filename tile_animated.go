package main

import "github.com/gdamore/tcell"

type TileAnimatedFrame struct {
	tileRune rune
	style    tcell.Style
}

type TileAnimated struct {
	walkable       bool
	interaction    Interaction
	frames         []TileAnimatedFrame
	animationSpeed int
}

func (t *TileAnimated) GetAppearance(frame int, position Position) (rune, tcell.Style) {
	idx := frame % len(t.frames)
	return t.frames[idx].tileRune, t.frames[idx].style
}

func (t *TileAnimated) GetWalkable() bool {
	return t.walkable
}

func (t *TileAnimated) SetWalkable(walkable bool) {
	t.walkable = walkable
}

func (t *TileAnimated) GetInteraction() Interaction {
	return t.interaction
}

func (t *TileAnimated) SetInteraction(interaction Interaction) {
	t.interaction = interaction
}

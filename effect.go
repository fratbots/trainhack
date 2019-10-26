package main

import "github.com/gdamore/tcell"

// EffectTile is a tile that was created during effect rendering.
type EffectTile struct {
	Position   Position
	Rune       rune
	Foreground tcell.Color
}

// Effect describes methods common for all effects.
type Effect interface {
	Update() bool
	Render() []EffectTile
}

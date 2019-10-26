package main

// EffectTile is a tile that was created during effect rendering.
type EffectTile struct {
	Position Position
	Rune     rune
}

// Effect describes methods common for all effects.
type Effect interface {
	Update() bool
	Render() []EffectTile
	GetBehavior() EffectBehavior
	Equals(Effect) bool
}

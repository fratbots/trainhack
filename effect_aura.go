package main

import "github.com/gdamore/tcell"

// EffectAura is an Aura effect around actor.
type EffectAura struct {
	longevity int
	frame     int
	target    *Actor
}

// NewEffectAura creates Aura effect.
func NewEffectAura(longevity int, target *Actor) *EffectAura {
	return &EffectAura{
		frame:     0,
		longevity: longevity,
		target:    target,
	}
}

// Update moves forward effect animation progress and returns false if animation has ended.
func (e *EffectAura) Update() bool {
	e.frame = e.frame + 1
	if e.frame >= e.longevity {
		return false
	}
	return true
}

// Render creates a set of tiles representing current frame of effect animation.
func (e *EffectAura) Render() []EffectTile {
	var runeLeft rune
	var runeRight rune
	var runeTop rune
	var runeBottom rune

	color := tcell.ColorGreen

	if e.frame >= e.longevity/2 {
		runeLeft = '+'
		runeRight = '+'
		runeTop = '+'
		runeBottom = '+'
	} else {
		runeLeft = '*'
		runeRight = '*'
		runeTop = '*'
		runeBottom = '*'
	}

	return []EffectTile{
		EffectTile{
			Position:   Position{e.target.Position.X - 1, e.target.Position.Y},
			Rune:       runeLeft,
			Foreground: color,
		},
		EffectTile{
			Position:   Position{e.target.Position.X + 1, e.target.Position.Y},
			Rune:       runeRight,
			Foreground: color,
		},
		EffectTile{
			Position:   Position{e.target.Position.X, e.target.Position.Y - 1},
			Rune:       runeTop,
			Foreground: color,
		},
		EffectTile{
			Position:   Position{e.target.Position.X, e.target.Position.Y + 1},
			Rune:       runeBottom,
			Foreground: color,
		},
	}
}

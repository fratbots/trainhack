package main

import "github.com/gdamore/tcell"

/*
This is helicopter
  0123456789012
             #
2           #    2
1      #   #     1
0 -x---==[()]>   0
1         #   #  1
2        #       2
        #
  0123456789012
*/

type EffectHelicopter struct {
	position Position
	speed    int
	frame    int
}

func NewEffectHelicopter() *EffectHelicopter {
	return &EffectHelicopter{
		speed:    1,
		frame:    0,
		position: Position{0, 12},
	}
}

func (e *EffectHelicopter) Update() bool {
	length := len(e.getBody())
	e.frame = e.frame + 1
	e.position.X = e.position.X + e.speed
	// TODO remove hardcode, get viewport width.
	if e.position.X > 100+length {
		return false
	}
	return true
}

func (e *EffectHelicopter) Render() []EffectTile {
	heli := append(e.getBody(), e.getBlades(e.frame%4)...)
	return e.normalizeCoords(heli)
}

func (e *EffectHelicopter) normalizeCoords(tiles []EffectTile) []EffectTile {
	length := len(e.getBody())
	for idx, tile := range tiles {
		tiles[idx].Position.X = tile.Position.X - length + e.position.X
		tiles[idx].Position.Y = tile.Position.Y + e.position.Y
	}
	return tiles
}

func (e *EffectHelicopter) getBody() []EffectTile {
	bodyColor := tcell.ColorDarkOrange
	return []EffectTile{
		EffectTile{
			Position:   Position{0, 0},
			Rune:       '-',
			Foreground: bodyColor,
		},
		EffectTile{
			Position:   Position{1, 0},
			Rune:       'x',
			Foreground: bodyColor,
		},
		EffectTile{
			Position:   Position{2, 0},
			Rune:       '-',
			Foreground: bodyColor,
		},
		EffectTile{
			Position:   Position{3, 0},
			Rune:       '-',
			Foreground: bodyColor,
		},
		EffectTile{
			Position:   Position{4, 0},
			Rune:       '-',
			Foreground: bodyColor,
		},
		EffectTile{
			Position:   Position{5, 0},
			Rune:       '=',
			Foreground: bodyColor,
		},
		EffectTile{
			Position:   Position{6, 0},
			Rune:       '=',
			Foreground: bodyColor,
		},
		EffectTile{
			Position:   Position{7, 0},
			Rune:       '[',
			Foreground: bodyColor,
		},
		EffectTile{
			Position:   Position{8, 0},
			Rune:       '(',
			Foreground: bodyColor,
		},
		EffectTile{
			Position:   Position{9, 0},
			Rune:       ')',
			Foreground: bodyColor,
		},
		EffectTile{
			Position:   Position{10, 0},
			Rune:       ']',
			Foreground: bodyColor,
		},
		EffectTile{
			Position:   Position{11, 0},
			Rune:       '>',
			Foreground: bodyColor,
		},
	}
}

func (e *EffectHelicopter) getBlades(phase int) []EffectTile {
	bladeRune := '#'
	bladeColor := tcell.ColorDarkOrange
	positions := [][]Position{
		[]Position{
			Position{5, -2},
			Position{12, -2},
			Position{7, -1},
			Position{10, -1},
			Position{7, 1},
			Position{10, 1},
			Position{5, 2},
			Position{12, 2},
		},
		[]Position{
			Position{5, -1},
			Position{9, -1},
			Position{10, -2},
			Position{11, -3},
			Position{6, 3},
			Position{7, 2},
			Position{8, 1},
			Position{12, 1},
		},
		[]Position{
			Position{5, -2},
			Position{12, -2},
			Position{7, -1},
			Position{10, -1},
			Position{7, 1},
			Position{10, 1},
			Position{5, 2},
			Position{12, 2},
		},
		[]Position{
			Position{12, -1},
			Position{8, -1},
			Position{7, -2},
			Position{6, -3},
			Position{11, 3},
			Position{10, 2},
			Position{9, 1},
			Position{5, 1},
		},
	}
	result := make([]EffectTile, 8)
	for idx, position := range positions[phase] {
		result[idx] = EffectTile{
			Position:   position,
			Rune:       bladeRune,
			Foreground: bladeColor,
		}
	}
	return result
}

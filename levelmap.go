package main

import "errors"

// LevelMap is a 2d map of a level.
type LevelMap struct {
	Width   int
	Height  int
	Texture Texture
}

// GetTile returns tile of a map from the specified coord.
func (m *LevelMap) GetTile(x int, y int) (Tile, error) {
	if x >= m.Width {
		return Tile('X'), errors.New("X coord out of range.")
	}
	if y >= m.Height {
		return Tile('X'), errors.New("Y coord out of range.")
	}
	idx := y*m.Width + x
	return m.Texture[idx], nil
}

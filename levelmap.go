package main

// LevelMap is a 2d map of a level.
type LevelMap struct {
	Width   int
	Height  int
	Texture Texture
}

// GetTile returns tile of a map from the specified coord.
func (m *LevelMap) GetTile(x int, y int) Tile {
	idx := y*m.Width + x
	return m.Texture[idx]
}

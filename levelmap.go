package main

// LevelMap is a 2d map of a level.
type LevelMap struct {
	Width     int
	Height    int
	Texture   Texture
	Obstacles []Obstacle
}

// GetTile returns tile of a map from the specified coord.
func (m *LevelMap) GetTile(x int, y int) *Tile {
	if x >= m.Width {
		return nil
	}
	if y >= m.Height {
		return nil
	}
	idx := y*m.Width + x
	return &m.Texture[idx]
}

package main

// LevelMap is a 2d map of a level.
type LevelMap struct {
	Dimensions Dimensions
	Width      int
	Height     int
	Texture    Texture
	Obstacles  []Obstacle
}

// GetTile returns tile of a map from the specified coord.
func (m *LevelMap) GetTile(pos Position) *Tile {
	if pos.X >= m.Dimensions.X || pos.X < 0 {
		return nil
	}
	if pos.Y >= m.Dimensions.Y || pos.Y < 0 {
		return nil
	}

	i := pos.Y*m.Dimensions.X + pos.X

	if i >= 0 && i < len(m.Texture) {
		return &m.Texture[i]
	}

	return nil
}

package main

// Actor is a behaving entity, for example, hero or enemy.
type ViewActor struct {
	X       int
	Y       int
	Width   int
	Height  int
	Texture Texture
}

// GetTile returns Actor tile based on relative coordinates from (0,0) to (a.Width,a.Height).
func (a *ViewActor) GetTile(x int, y int) Tile {
	idx := y*a.Width + x
	return a.Texture[idx]
}

package main

// Mixer is the layers compositor. It knows, how to place objects on the level map.
type Mixer struct {
	buffer   *RenderBuffer
	levelMap *LevelMap
	viewport *Viewport
}

// NewMixer returns new Mixer.
func NewMixer(buffer *RenderBuffer, levelMap *LevelMap, viewport *Viewport) *Mixer {
	return &Mixer{
		buffer:   buffer,
		levelMap: levelMap,
		viewport: viewport,
	}
}

// Mix fills the render buffer with ready to render image.
func (m *Mixer) Mix(actors []*Actor) {
	// First layer: the map.
	for x := 0; x < m.viewport.Width; x++ {
		for y := 0; y < m.viewport.Height; y++ {
			mapTile := m.levelMap.GetTile(
				m.viewport.ToMapCoordX(x),
				m.viewport.ToMapCoordY(y),
			)
			m.buffer.SetTile(x, y, mapTile)
		}
	}
	// Second layer: actors.
	for _, actor := range actors {
		for x := 0; x < actor.Width; x++ {
			for y := 0; y < actor.Height; y++ {
				vpX := m.viewport.ToViewportCoordX(actor.X)
				vpY := m.viewport.ToViewportCoordY(actor.Y)
				actorTile := actor.GetTile(x, y)
				if vpX >= 0 && vpY >= 0 {
					m.buffer.SetTile(
						vpX,
						vpY,
						actorTile,
					)
				}
			}
		}
	}
}

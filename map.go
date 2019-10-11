package main

import "fmt"

// Map is a global 2d map of a level.
type Map struct {
	Width   int
	Height  int
	Texture Texture
}

// GetTile returns tile of a map from the specified coord.
func (m *Map) GetTile(x int, y int) Tile {
	return m.Texture[y*m.Width+x]
}

// Viewport is a visible part of a map. It uses Map coords.
type Viewport struct {
	X      int
	Y      int
	Width  int
	Height int
}

// Actor is a behaving entity, for example, hero or enemy.
type Actor struct {
	X       int
	Y       int
	Width   int
	Height  int
	Texture Texture
}

// RenderBuffer contains rendered image ready to be printed out.
type RenderBuffer struct {
	Width  int
	Height int
	buffer []Tile
}

// NewRenderBuffer returns new render buffer.
func NewRenderBuffer(width int, height int) *RenderBuffer {
	return &RenderBuffer{
		Width:  width,
		Height: height,
		buffer: make([]Tile, width*height),
	}
}

// Set places tile to specified position in render buffer.
func (rb *RenderBuffer) SetTile(x int, y int, t Tile) {
	idx := y*rb.Width + x
	rb.buffer[idx] = t
}

// Get returns all the contents from buffer.
func (rb *RenderBuffer) Get() []Tile {
	return rb.buffer
}

// Tile is the smallest part of a texture.
type Tile string

// Texture is a matrix of symbols representing a look of an entity.
type Texture []Tile

// Mixer is the layers compositor. It knows, how to place objects on the level map.
type Mixer struct {
	buffer   *RenderBuffer
	levelMap *Map
	viewport *Viewport
}

// NewMixer returns new Mixer.
func NewMixer(buffer *RenderBuffer, levelMap *Map, viewport *Viewport) *Mixer {
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
			m.buffer.SetTile(x, y, m.levelMap.GetTile(x, y))
		}
	}
	// Second layer: actors.
	for _, actor := range actors {
		for x := 0; x < actor.Width; x++ {
			for y := 0; y < actor.Height; y++ {
				idx := y*actor.Width + x
				m.buffer.SetTile(actor.X, actor.Y, actor.Texture[idx])
			}
		}
	}
}

// Renderer renders game objects.
type Renderer struct {
	buffer *RenderBuffer
}

// NewRenderer returns new Renderer.
func NewRenderer(buffer *RenderBuffer) *Renderer {
	return &Renderer{
		buffer: buffer,
	}
}

// Render draws the frame based on map and game objects.
func (r *Renderer) Render() {
	tiles := r.buffer.Get()
	for x := 0; x < r.buffer.Width; x++ {
		for y := 0; y < r.buffer.Height; y++ {
			idx := y*r.buffer.Width + x
			fmt.Printf("%s", tiles[idx])
		}
		fmt.Printf("\n")
	}
}

func main() {
	levelMapTexture := []Tile{
		Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."),
		Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."),
		Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."),
		Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."),
		Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."),
		Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."),
	}
	levelMap := &Map{
		Width:   8,
		Height:  6,
		Texture: levelMapTexture,
	}
	viewport := &Viewport{
		X:      0,
		Y:      0,
		Width:  8,
		Height: 6,
	}
	hero := &Actor{
		X:       4,
		Y:       4,
		Width:   1,
		Height:  1,
		Texture: []Tile{Tile("@")},
	}
	actors := []*Actor{
		hero,
	}

	rb := NewRenderBuffer(8, 6)
	mixer := NewMixer(rb, levelMap, viewport)
	mixer.Mix(actors)
	r := NewRenderer(rb)
	r.Render()
}

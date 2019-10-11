package main

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

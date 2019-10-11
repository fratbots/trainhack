package main

import "fmt"

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
	for y := 0; y < r.buffer.Height; y++ {
		for x := 0; x < r.buffer.Width; x++ {
			idx := y*r.buffer.Width + x
			fmt.Printf("%s", tiles[idx])
		}
		fmt.Printf("\n")
	}
}

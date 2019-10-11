package main

// Viewport is a visible part of a map. It uses LevelMap coords.
type Viewport struct {
	X        int
	Y        int
	Width    int
	Height   int
	levelMap *LevelMap
}

// NewViewport returns new Viewport.
func NewViewport(x int, y int, width int, height int, levelMap *LevelMap) *Viewport {
	return &Viewport{
		X:        x,
		Y:        y,
		Width:    width,
		Height:   height,
		levelMap: levelMap,
	}
}

// Returns LevelMap coord X based on Viewport coord X.
func (v *Viewport) ToMapCoordX(vpX int) int {
	return v.X + vpX
}

// Returns LevelMap coord Y based on Viewport coord Y.
func (v *Viewport) ToMapCoordY(vpY int) int {
	return v.Y + vpY
}

// Returns Viewport coord X based on LevelMap coord X.
func (v *Viewport) ToViewportCoordX(mapX int) int {
	return mapX - v.X
}

// Returns Viewport coord Y based on LevelMap coord Y.
func (v *Viewport) ToViewportCoordY(mapY int) int {
	return mapY - v.Y
}

package main

// Tile is the smallest part of a texture.
type Tile struct {
	Symbol rune
}

// Texture is a matrix of symbols representing a look of an entity.
type Texture []Tile

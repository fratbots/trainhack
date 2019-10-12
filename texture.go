package main

// Tile is the smallest part of a texture.
type Tile struct {
	Symbol  rune
	FgColor string
	BgColor string
}

// NewTile returns new tile.
func NewTile(symbol rune, fgColor string, bgColor string) Tile {
	return Tile{
		Symbol:  symbol,
		FgColor: fgColor,
		BgColor: bgColor,
	}
}

// Texture is a matrix of symbols representing a look of an entity.
type Texture []Tile

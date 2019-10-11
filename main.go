package main

func main() {
	game := NewGame()
	game.Start(&HelloScreen{})

	// mapTest()
}

func mapTest() {
	levelMapTexture := []Tile{
		Tile("r"), Tile("-"), Tile("-"), Tile("-"), Tile("-"), Tile("-"), Tile("-"), Tile("-"), Tile("-"), Tile("-"), Tile("-"), Tile("7"),
		Tile("|"), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("|"),
		Tile("|"), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("|"),
		Tile("|"), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("|"),
		Tile("|"), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("|"),
		Tile("|"), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("|"),
		Tile("|"), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("|"),
		Tile("|"), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("|"),
		Tile("|"), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("."), Tile("|"),
		Tile("L"), Tile("_"), Tile("_"), Tile("_"), Tile("_"), Tile("_"), Tile("_"), Tile("_"), Tile("_"), Tile("_"), Tile("_"), Tile("J"),
	}
	levelMap := &LevelMap{
		Width:   12,
		Height:  10,
		Texture: levelMapTexture,
	}

	viewportWidth := 8
	viewportHeight := 6
	viewportX := 0
	viewportY := 0
	viewport := NewViewport(
		viewportX,
		viewportY,
		viewportWidth,
		viewportHeight,
		levelMap,
	)

	hero := &Actor{
		X:       5,
		Y:       5,
		Width:   1,
		Height:  1,
		Texture: []Tile{Tile("@")},
	}
	actors := []*Actor{
		hero,
	}

	renderBufWidth := viewportWidth
	renderBufHeight := viewportHeight
	renderBuf := NewRenderBuffer(renderBufWidth, renderBufHeight)
	mixer := NewMixer(renderBuf, levelMap, viewport)
	mixer.Mix(actors)
	renderer := NewRenderer(renderBuf)
	renderer.Render()
}

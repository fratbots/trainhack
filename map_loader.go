package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// MapLoader loads level maps from filesystem.
type MapLoader struct {
}

// FileTilePoint is a FileTile coords container.
type FileTileCoords struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// FileTile represents single tile stored in a map file.
type FileTile struct {
	Point   FileTileCoords `json:"point"`
	Symbol  string         `json:"character"`
	FgColor string         `json:"foreground"`
	BgColor string         `json:"background"`
}

// Load reads level map from file.
func (m *MapLoader) Load(path string) (LevelMap, error) {
	fileReader, err := os.Open(path)
	if err != nil {
		return LevelMap{}, err
	}
	defer fileReader.Close()

	var tiles [][]FileTile
	err = json.NewDecoder(fileReader).Decode(&tiles)
	if err != nil {
		return LevelMap{}, fmt.Errorf("failed to load map texture from file %s: %v", path, err)
	}

	var width int
	height := len(tiles)
	var tex Texture
	for _, row := range tiles {
		width = len(row)
		for i := 0; i < width; i++ {
			tex = append(tex, NewTile(rune(row[i].Symbol[0]), row[i].FgColor, row[i].BgColor))
		}
	}
	return LevelMap{
		Width:   width,
		Height:  height,
		Texture: tex,
	}, nil
}

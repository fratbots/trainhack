package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	MapMeta      = "meta"
	MapTexture   = "texture"
	MapObstacles = "obstacles"
	MapDir       = "./levels"
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
func (m *MapLoader) Load(mapName string) (LevelMap, error) {
	path := fmt.Sprintf("%s/%s/%s.json", MapDir, mapName, MapTexture)
	fileReader, err := os.Open(path)
	defer fileReader.Close()
	if err != nil {
		return LevelMap{}, err
	}
	decoder := json.NewDecoder(fileReader)
	var tiles [][]FileTile
	err = decoder.Decode(&tiles)
	if err != nil {
		return LevelMap{}, fmt.Errorf("Failed to load map texture from file %s: %v", path, err)
	}
	var width int
	height := len(tiles)
	var textureTiles []Tile
	for _, row := range tiles {
		width = len(row)
		for i := 0; i < width; i++ {
			textureTiles = append(
				textureTiles,
				NewTile(rune(row[i].Symbol[0]), row[i].FgColor, row[i].BgColor),
			)
		}
	}
	return LevelMap{
		Width:   width,
		Height:  height,
		Texture: Texture(textureTiles),
	}, nil
}

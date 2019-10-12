package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/naoina/toml"
)

const (
	MapMeta           = "meta"
	MapLayerTexture   = "texture"
	MapLayerObstacles = "obstacles"
	MapDir            = "./levels"
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

type FileObstacles struct {
	Obstacles []Obstacle
}

// Load reads level map from file.
func (m *MapLoader) Load(mapName string) (LevelMap, error) {
	// Load texture.
	texturePath := fmt.Sprintf("%s/%s/%s.json", MapDir, mapName, MapLayerTexture)
	width, height, textureTiles, err := m.loadTexture(texturePath)
	if err != nil {
		return LevelMap{}, fmt.Errorf("Failed to load map texture: %v", err)
	}

	// Load obstacles.
	obstaclesPath := fmt.Sprintf("%s/%s/%s.toml", MapDir, mapName, MapLayerObstacles)
	obstacles, err := m.loadObstacles(obstaclesPath)
	if err != nil {
		return LevelMap{}, fmt.Errorf("Failed to load map obstacles: %v", err)
	}

	return LevelMap{
		Width:     width,
		Height:    height,
		Texture:   Texture(textureTiles),
		Obstacles: obstacles,
	}, nil
}

// loadTexture loads map texture from file.
func (m *MapLoader) loadTexture(path string) (int, int, []Tile, error) {
	fileReader, err := os.Open(path)
	defer fileReader.Close()
	var textureTiles []Tile
	if err != nil {
		return 0, 0, Texture(textureTiles), err
	}
	decoder := json.NewDecoder(fileReader)
	var tiles [][]FileTile
	err = decoder.Decode(&tiles)
	if err != nil {
		return 0, 0, Texture(textureTiles), fmt.Errorf("Failed to load map texture from file %s: %v", path, err)
	}
	var width int
	height := len(tiles)
	for _, row := range tiles {
		width = len(row)
		for i := 0; i < width; i++ {
			textureTiles = append(
				textureTiles,
				NewTile(rune(row[i].Symbol[0]), row[i].FgColor, row[i].BgColor),
			)
		}
	}
	return width, height, Texture(textureTiles), nil
}

// loadObstacles loads obstacles from file.
func (m *MapLoader) loadObstacles(path string) ([]Obstacle, error) {
	var fo FileObstacles
	var obstacles []Obstacle
	f, err := os.Open(path)
	if err != nil {
		return obstacles, err
	}
	if err := toml.NewDecoder(f).Decode(&fo); err != nil {
		return obstacles, err
	}
	fmt.Println(fo)
	return fo.Obstacles, nil
}

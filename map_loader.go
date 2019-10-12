package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// MapLoader loads level maps from filesystem.
type MapLoader struct {
}

// Load reads level map from file.
func (m *MapLoader) Load(path string) (LevelMap, error) {
	fileReader, err := os.Open(path)
	defer fileReader.Close()
	if err != nil {
		return LevelMap{}, err
	}
	bytesReader := bufio.NewReader(fileReader)
	header, err := bytesReader.ReadBytes('\n')
	if err != nil {
		return LevelMap{}, err
	}

	width, height, err := m.parseHeader(header)
	if err != nil {
		return LevelMap{}, fmt.Errorf("Failed to parse map file header: %v", err)
	}

	texture, err := m.loadTexture(bytesReader, width, height)
	if err != nil {
		return LevelMap{}, fmt.Errorf("Failed to load map contents: %v", err)
	}

	return LevelMap{
		Width:   width,
		Height:  height,
		Texture: texture,
	}, nil
}

// parseHeader returns level map metainformation such as width and height.
func (m *MapLoader) parseHeader(header []byte) (int, int, error) {
	if len(header) < 4 || string(header[:4]) != "map|" {
		return 0, 0, errors.New("Map file header corrupted. Expected format: \"map|80x60\".")
	}
	str := string(header)
	splitted := strings.Split(str[4:len(str)-1], "x")
	if len(splitted) != 2 {
		return 0, 0, errors.New("Map file header corrupted [x]. Expected format: \"map|80x60\".")
	}
	width, err := strconv.Atoi(splitted[0])
	if err != nil {
		return 0, 0, fmt.Errorf("Map file header corrupted [width]. Expected format: \"map|80x60\". Error: %v", err)
	}
	height, err := strconv.Atoi(splitted[1])
	if err != nil {
		return 0, 0, fmt.Errorf("Map file header corrupted [height]. Expected format: \"map|80x60\". Error: %v", err)
	}
	return width, height, nil
}

// loadTexture reads texture from map file.
func (m *MapLoader) loadTexture(reader *bufio.Reader, width int, height int) (Texture, error) {
	buf := make([]byte, width*height)
	if _, err := io.ReadFull(reader, buf); err != nil {
		return Texture{}, err
	}
	tiles := make([]Tile, width*height)
	for i := 0; i < len(buf); i++ {
		r := rune(buf[i])
		tiles[i] = Tile{Symbol: r}
	}
	return Texture(tiles), nil
}

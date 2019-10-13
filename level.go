package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/gdamore/tcell"
)

type Til struct {
	Rune       rune
	Style      tcell.Style
	IsWalkable bool
}

type Level struct {
	Dimensions Dimensions
	Tiles      []Til
}

func (l *Level) GetTile(pos Position) *Til {
	if !pos.IsOn(l.Dimensions) {
		return nil
	}

	i := pos.Y*l.Dimensions.X + pos.X
	if i >= 0 && i < len(l.Tiles) {
		return &l.Tiles[i]
	}

	return nil
}

func LoadLevel(name string) *Level {
	path := fmt.Sprintf("./levels/%s/texture.txt", name)

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}

	b = b[3:]

	l := Level{
		Dimensions: Dimensions{},
		Tiles:      nil,
	}

	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		w := len(scanner.Text())

		if l.Dimensions.X != 0 && l.Dimensions.X != w {
			return nil
		}
		l.Dimensions.X = w

		for _, r := range scanner.Text() {
			l.Tiles = append(l.Tiles, TileParser(r))
		}
		l.Dimensions.Y++
	}

	return &l
}

var (
	colorDarkGreen   = tcell.NewRGBColor(0, 200, 0)
	colorGreen       = tcell.NewRGBColor(0, 255, 0)
	colorForestGreen = tcell.NewRGBColor(0, 200, 0)
	colorDarkBlue    = tcell.NewRGBColor(0, 0, 200)
	colorBlue        = tcell.NewRGBColor(0, 0, 255)

	styleGrass  = tcell.StyleDefault.Background(colorGreen).Foreground(colorDarkGreen)
	styleForest = tcell.StyleDefault.Background(colorDarkGreen).Foreground(colorForestGreen)
	styleWater  = tcell.StyleDefault.Background(colorDarkBlue).Foreground(colorBlue)

	tiles = map[rune]Til{
		'.': {
			Rune:       '.',
			Style:      styleGrass,
			IsWalkable: true,
		},
		'u': {
			Rune:       tcell.RuneBlock,
			Style:      styleForest,
			IsWalkable: false,
		},
		'y': {
			Rune:       tcell.RuneBoard,
			Style:      styleForest,
			IsWalkable: false,
		},
		'k': {
			Rune:       tcell.RuneCkBoard,
			Style:      styleForest,
			IsWalkable: false,
		},
		'h': {
			Rune:       tcell.RuneCkBoard,
			Style:      styleForest,
			IsWalkable: false,
		},
		'/': {
			Rune:       ' ',
			Style:      styleWater,
			IsWalkable: false,
		},
		'^': {
			Rune:       '^',
			Style:      styleWater,
			IsWalkable: true,
		},
	}
)

func TileParser(r rune) Til {
	if t, ok := tiles[r]; ok {
		return t
	}

	return Til{
		Rune:  r,
		Style: tcell.StyleDefault,
	}
}

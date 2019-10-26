package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gdamore/tcell"
)

type Tile struct {
	Rune        rune
	Style       tcell.Style
	IsWalkable  bool
	Interaction Interaction
}

type Level struct {
	Dimensions Dimensions
	Tiles      []Tile
	Doors      map[rune]Position
}

type door struct {
	Rune     rune
	Map      string
	Door     rune
	Position Position
}

func (l *Level) GetTile(pos Position) *Tile {
	if !pos.IsOn(l.Dimensions) {
		return nil
	}

	i := pos.Y*l.Dimensions.X + pos.X
	if i >= 0 && i < len(l.Tiles) {
		return &l.Tiles[i]
	}

	return nil
}

func LoadLevel(g *Game, name string) *Level {
	path := fmt.Sprintf("./levels/%s/texture.txt", name)

	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic("cannot load map " + name)
	}

	// skip BOM
	if bytes.HasPrefix(b, []byte{0xEF, 0xBB, 0xBF}) {
		b = b[3:]
	}

	level := Level{
		Dimensions: Dimensions{},
		Tiles:      nil,
		Doors:      map[rune]Position{},
	}

	doors := scanDoors(b)

	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		if len(scanner.Text()) == 0 || strings.HasPrefix(scanner.Text(), "// ") {
			continue
		}

		w := len(scanner.Text())

		if level.Dimensions.X != 0 && level.Dimensions.X != w {
			return nil
		}
		level.Dimensions.X = w

		for x, r := range scanner.Text() {
			pos := Position{X: x, Y: level.Dimensions.Y}
			level.Tiles = append(level.Tiles, TileParser(g, r, pos, doors))
		}
		level.Dimensions.Y++
	}

	// doors has positions after TileParser
	for k, door := range doors {
		level.Doors[k] = door.Position
	}

	return &level
}

func scanDoors(b []byte) map[rune]door {
	doors := map[rune]door{}

	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "// ") {
			var run, displayRune, mapLocation rune
			var mapName string
			n, err := fmt.Sscanf(scanner.Text()[3:], "rune:%c display:%c map:%s location:%c",
				&run, &displayRune, &mapName, &mapLocation)

			if err == nil && n == 4 {
				doors[run] = door{
					Rune: displayRune,
					Map:  mapName,
					Door: mapLocation,
				}
			}
		}
	}

	return doors
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
	styleGround = tcell.StyleDefault.Background(tcell.ColorSandyBrown).Foreground(tcell.ColorSandyBrown)

	tiles = map[rune]Tile{
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
		'_': {
			Rune:       ' ',
			Style:      styleGround,
			IsWalkable: true,
		},
	}
)

func TileParser(g *Game, r rune, pos Position, doors map[rune]door) Tile {

	isWalkable := false
	var interaction Interaction
	// rune is door
	if door, ok := doors[r]; ok {

		// set position
		door.Position = pos
		doors[r] = door

		r = door.Rune
		isWalkable = true

		openTheDoor := func() Result {
			d := door.Door
			g.Sound.PlayContext(SoundContextDoor)
			g.SetScreen(NewScreenStage(g, door.Map, &d))
			return Result{}
		}

		interaction = func(actor *Actor) *Action {
			return &Action{
				Actor:    actor,
				Deferred: true,
				Perform:  openTheDoor, /*func() Result {
					return Result{
						Updated:               true,
						AlternativeIsDeferred: true,
						Alternative:           &openTheDoor,
					}
				},*/
			}
		}
	}

	if t, ok := tiles[r]; ok {
		if !t.IsWalkable && isWalkable {
			t.IsWalkable = true
		}
		t.Interaction = interaction
		return t
	}

	return Tile{
		Rune:        r,
		Style:       tcell.StyleDefault,
		Interaction: interaction,
		IsWalkable:  isWalkable,
	}
}

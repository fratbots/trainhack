package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gdamore/tcell"
)

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

func (l *Level) GetTile(pos Position) Tile {
	if !pos.IsOn(l.Dimensions) {
		return nil
	}

	i := pos.Y*l.Dimensions.X + pos.X
	if i >= 0 && i < len(l.Tiles) {
		return l.Tiles[i]
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
	colorDarkBlue    = tcell.NewRGBColor(0, 212, 255)
	colorBlue        = tcell.NewRGBColor(50, 220, 255)
	colorLightBlue   = tcell.NewRGBColor(89, 227, 255)
	colorWhite       = tcell.NewRGBColor(255, 255, 255)

	styleGrass       = tcell.StyleDefault.Background(colorGreen).Foreground(colorDarkGreen)
	styleForest      = tcell.StyleDefault.Background(colorDarkGreen).Foreground(colorForestGreen)
	styleWaterDark   = tcell.StyleDefault.Background(colorDarkBlue).Foreground(colorBlue)
	styleWaterMedium = tcell.StyleDefault.Background(colorBlue).Foreground(colorLightBlue)
	styleWaterLight  = tcell.StyleDefault.Background(colorLightBlue).Foreground(colorWhite)
	styleGround      = tcell.StyleDefault.Background(tcell.ColorSandyBrown).Foreground(tcell.ColorSandyBrown)

	tiles = map[rune]Tile{
		'.': &TileStatic{
			tileRune: '.',
			style:    styleGrass,
			walkable: true,
		},
		'u': &TileStatic{
			tileRune: tcell.RuneBlock,
			style:    styleForest,
			walkable: false,
		},
		'y': &TileStatic{
			tileRune: tcell.RuneBoard,
			style:    styleForest,
			walkable: false,
		},
		'k': &TileStatic{
			tileRune: tcell.RuneCkBoard,
			style:    styleForest,
			walkable: false,
		},
		'h': &TileStatic{
			tileRune: tcell.RuneCkBoard,
			style:    styleForest,
			walkable: false,
		},
		'/': &TileAnimated{
			frames: []TileAnimatedFrame{
				TileAnimatedFrame{
					tileRune: '-',
					style:    styleWaterDark,
				},
				TileAnimatedFrame{
					tileRune: '^',
					style:    styleWaterDark,
				},
				TileAnimatedFrame{
					tileRune: ' ',
					style:    styleWaterDark,
				},
				TileAnimatedFrame{
					tileRune: '.',
					style:    styleWaterDark,
				},
				TileAnimatedFrame{
					tileRune: '-',
					style:    styleWaterMedium,
				},
				TileAnimatedFrame{
					tileRune: '^',
					style:    styleWaterMedium,
				},
				TileAnimatedFrame{
					tileRune: ' ',
					style:    styleWaterMedium,
				},
				TileAnimatedFrame{
					tileRune: '.',
					style:    styleWaterMedium,
				},
			},
			animationSpeed: 1,
			walkable:       false,
		},
		'^': &TileAnimated{
			frames: []TileAnimatedFrame{
				TileAnimatedFrame{
					tileRune: '.',
					style:    styleWaterMedium,
				},
				TileAnimatedFrame{
					tileRune: '-',
					style:    styleWaterMedium,
				},
				TileAnimatedFrame{
					tileRune: ' ',
					style:    styleWaterMedium,
				},
				TileAnimatedFrame{
					tileRune: '.',
					style:    styleWaterLight,
				},
				TileAnimatedFrame{
					tileRune: '-',
					style:    styleWaterLight,
				},
				TileAnimatedFrame{
					tileRune: ' ',
					style:    styleWaterLight,
				},
			},
			animationSpeed: 1,
			walkable:       true,
		},
		'_': &TileStatic{
			tileRune: ' ',
			style:    styleGround,
			walkable: true,
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

		interaction = func(actor *Actor) *Action {
			return &Action{
				Actor:    actor,
				Deferred: true,
				Perform: func() Result {
					if !actor.Class.IsHero {
						return FailureResult
					}
					d := door.Door
					g.Sound.PlayContext(SoundContextDoor)
					g.SetScreen(NewScreenStage(g, door.Map, &d))
					return Result{}
				},
			}
		}
	}

	if t, ok := tiles[r]; ok {
		tile := t.Copy()
		if !tile.GetWalkable() && isWalkable {
			tile.SetWalkable(true)
		}
		tile.SetInteraction(interaction)
		return tile
	}

	return &TileStatic{
		tileRune:    r,
		style:       tcell.StyleDefault,
		interaction: interaction,
		walkable:    isWalkable,
	}
}

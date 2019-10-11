package main

import (
	"math/rand"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type ScreenMap struct {
}

func (s *ScreenMap) Do(g *Game, end func(s Screen)) tview.Primitive {
	box := tview.NewBox()
	box.
		SetDrawFunc(func(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
			_, _, innerRectWidth, innerRectHeight := box.GetInnerRect()
			draw(screen, innerRectWidth, innerRectHeight)
			return 0, 0, 0, 0
		}).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Rune() == 'a' {
				g.View.Draw()
			}
			return event
		})

	return box
}

// draw is a callback function that draws the frame.
func draw(screen tcell.Screen, screenWidth int, screenHeight int) {
	levelMap := getMap(100, 100)

	viewportX := 0
	viewportY := 0
	viewport := NewViewport(
		viewportX,
		viewportY,
		screenWidth,
		screenHeight,
		levelMap,
	)

	hero := Actor{
		X:       rand.Intn(100),
		Y:       rand.Intn(100),
		Width:   1,
		Height:  1,
		Texture: []Tile{Tile('@')},
	}
	actors := []Actor{
		hero,
	}

	composer := NewComposer(screen, viewport)
	composer.RenderLevelMap(levelMap)
	composer.RenderActors(actors)
	composer.Finalize()
}

// Composer is a manager of layers that is capable of layers ordering.
type Composer struct {
	screen   tcell.Screen
	viewport Viewport
}

// NewComposer returns new Composer.
func NewComposer(screen tcell.Screen, viewport Viewport) Composer {
	return Composer{
		screen:   screen,
		viewport: viewport,
	}
}

// RenderLevelMap renders layer of level map.
func (c *Composer) RenderLevelMap(levelMap LevelMap) {
	for vpX := 0; vpX < c.viewport.Width; vpX++ {
		for vpY := 0; vpY < c.viewport.Height; vpY++ {
			mapTile := levelMap.GetTile(
				c.viewport.ToMapCoordX(vpX),
				c.viewport.ToMapCoordY(vpY),
			)
			c.screen.SetContent(vpX, vpY, rune(mapTile), nil, tcell.StyleDefault)
		}
	}
}

// RenderActors applies actors on top of levelMap.
func (c *Composer) RenderActors(actors []Actor) {
	for _, actor := range actors {
		for actorX := 0; actorX < actor.Width; actorX++ {
			for actorY := 0; actorY < actor.Height; actorY++ {
				vpX := c.viewport.ToViewportCoordX(actor.X)
				vpY := c.viewport.ToViewportCoordY(actor.Y)
				actorTile := actor.GetTile(actorX, actorY)
				if vpX >= 0 && vpY >= 0 {
					c.screen.SetContent(vpX, vpY, rune(actorTile), nil, tcell.StyleDefault)
				}
			}
		}
	}
}

// Finalize completes frame rendering.
func (c *Composer) Finalize() {
	c.screen.Show()
}

// getMap returns pregenerated map of the level.
func getMap(width int, height int) LevelMap {
	levelMapTexture := make([]Tile, width*height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			symbol := '.'
			if x%10 == 0 {
				symbol = ','
			}
			levelMapTexture[x*y] = Tile(symbol)
		}
	}
	return LevelMap{
		Width:   12,
		Height:  10,
		Texture: levelMapTexture,
	}
}

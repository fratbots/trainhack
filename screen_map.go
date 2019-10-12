package main

import (
	"log"
	"math/rand"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const (
	Map2 = "map2"
)

type ScreenMap struct {
}

func (s *ScreenMap) Do(g *Game, end func(s Screen)) tview.Primitive {
	box := tview.NewBox()
	box.
		SetDrawFunc(func(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
			_, _, innerRectWidth, innerRectHeight := box.GetInnerRect()
			err := draw(screen, innerRectWidth, innerRectHeight)
			if err != nil {
				log.Fatalf("Failed to draw the frame: %v", err)
			}
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
func draw(screen tcell.Screen, screenWidth int, screenHeight int) error {
	mapLoader := MapLoader{}
	levelMap, err := mapLoader.Load(Map2)
	if err != nil {
		return err
	}

	hero := ViewActor{
		X:      rand.Intn(levelMap.Width),
		Y:      rand.Intn(levelMap.Height),
		Width:  1,
		Height: 1,
		Texture: []Tile{
			Tile{Symbol: '@'},
		},
	}

	viewportX := hero.X - screenWidth/2
	viewportY := hero.Y - screenHeight/2
	viewport := NewViewport(
		viewportX,
		viewportY,
		screenWidth,
		screenHeight,
	)

	actors := []ViewActor{
		hero,
	}

	composer := NewComposer(screen, viewport)
	composer.RenderLevelMap(levelMap)
	composer.RenderActors(actors)
	composer.Finalize()

	return nil
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
	for vpY := 0; vpY < c.viewport.Height; vpY++ {
		for vpX := 0; vpX < c.viewport.Width; vpX++ {
			mapTile := levelMap.GetTile(Position{
				X: c.viewport.ToMapCoordX(vpX),
				Y: c.viewport.ToMapCoordY(vpY),
			})
			if mapTile == nil {
				continue
			}
			c.screen.SetContent(vpX, vpY, mapTile.Symbol, nil, tcell.StyleDefault)
		}
	}
}

// RenderActors applies actors on top of levelMap.
func (c *Composer) RenderActors(actors []ViewActor) {
	for _, actor := range actors {
		for actorY := 0; actorY < actor.Height; actorY++ {
			for actorX := 0; actorX < actor.Width; actorX++ {
				vpX := c.viewport.ToViewportCoordX(actor.X)
				vpY := c.viewport.ToViewportCoordY(actor.Y)
				actorTile := actor.GetTile(actorX, actorY)
				if vpX >= 0 && vpY >= 0 {
					c.screen.SetContent(vpX, vpY, actorTile.Symbol, nil, tcell.StyleDefault)
				}
			}
		}
	}
}

// Finalize completes frame rendering.
func (c *Composer) Finalize() {
	c.screen.Show()
}

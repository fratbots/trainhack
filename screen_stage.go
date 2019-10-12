package main

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type ScreenStage struct {
	Stage *Stage
}

func (s *ScreenStage) Do(g *Game, end func(next Screen)) tview.Primitive {

	s.Stage = NewStage(g)
	s.Stage.Hero.Position = Position{X: 20, Y: 10}
	s.Stage.Load("map2", "")
	s.Stage.AddActor(Pursue(NewActor(Position{X: 7, Y: 5}, 0.3), s.Stage, s.Stage.Hero))

	lookAt := s.Stage.Hero.Position

	box := tview.NewBox()
	box.SetTitle("title")
	box.SetBorder(true)
	box.SetBackgroundColor(tcell.ColorGreen)

	box.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp:
			s.Stage.Hero.NextAction(ActionMove(s.Stage, s.Stage.Hero, DirectionTop))
		case tcell.KeyDown:
			s.Stage.Hero.NextAction(ActionMove(s.Stage, s.Stage.Hero, DirectionDown))
		case tcell.KeyLeft:
			s.Stage.Hero.NextAction(ActionMove(s.Stage, s.Stage.Hero, DirectionLeft))
		case tcell.KeyRight:
			s.Stage.Hero.NextAction(ActionMove(s.Stage, s.Stage.Hero, DirectionRight))
		}

		return nil
	})

	box.SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		x, y, width, height = box.GetInnerRect()

		lookAt = lookAt.FollowGap(s.Stage.Hero.Position, 5)
		dimensions := Dimensions{width, height}
		shift := Dimensions{x, y}
		port := NewPort(dimensions, shift, lookAt)

		// Tiles
		for iy := 0; iy < height; iy++ {
			for ix := 0; ix < width; ix++ {
				mapPos := port.ToMap(Position{ix, iy})
				if tile := s.Stage.LevelMap.GetTile(mapPos); tile != nil {
					s.drawTile(screen, tile, mapPos, Position{X: ix + shift.X, Y: iy + shift.Y})
				}
			}
		}

		// Actors
		for _, actor := range s.Stage.Actors {
			screenPos := port.ToScreen(actor.Position)
			if screenPos.IsOn(dimensions) {
				s.drawActor(screen, actor, screenPos)
			}
		}

		// HUD
		box.SetTitle(fmt.Sprintf("%dx%d", s.Stage.Hero.Position.X, s.Stage.Hero.Position.Y))

		return x, y, width, height
	})

	s.Stage.Start()

	return box
}

var backStyle = tcell.StyleDefault.Background(tcell.ColorGreen).Foreground(tcell.ColorForestGreen)

func (s *ScreenStage) drawTile(screen tcell.Screen, tile *Tile, mapPos, screenPos Position) {

	screen.SetContent(screenPos.X, screenPos.Y, tile.Symbol, nil, backStyle)

}

func (s *ScreenStage) drawActor(screen tcell.Screen, actor *Actor, screenPos Position) {

	screen.SetContent(screenPos.X, screenPos.Y, actor.Rune, nil, tcell.StyleDefault)

}

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
	s.Stage.Hero.Position = Position{X: 5, Y: 5}
	s.Stage.AddActor(Pursue(NewActor(Position{X: 7, Y: 5}, 0.3), s.Stage, s.Stage.Hero))

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

		for _, actor := range s.Stage.Actors {
			if actor.IsHero {
				box.SetTitle(fmt.Sprintf("%dx%d", actor.Position.X, actor.Position.Y))
				screen.SetContent(actor.Position.X, actor.Position.Y, '@', nil, tcell.StyleDefault)
			} else {
				screen.SetContent(actor.Position.X, actor.Position.Y, '&', nil, tcell.StyleDefault)
			}
		}

		return x, y, width, height
	})

	s.Stage.StartTime()

	return box
}

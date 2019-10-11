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

	s.Stage = NewStage()
	s.Stage.Hero.Position = Position{X: 5, Y: 5}

	skeleton := Actor{
		IsHero:   false,
		Position: Position{X: 10, Y: 6},
		Energy:   Energy{0},
		Speed:    0.3,
		Behavior: nil,
	}
	skeleton.Behavior = Pursue(s.Stage, &skeleton, s.Stage.Hero)

	s.Stage.AddActor(&skeleton)

	box := tview.NewBox()
	box.SetTitle("title")
	box.SetBorder(true)
	box.SetBackgroundColor(tcell.ColorGreen)

	box.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		var action Action

		switch event.Key() {
		case tcell.KeyUp:
			action = ActionMove(s.Stage, s.Stage.Hero, DirectionTop)
		case tcell.KeyDown:
			action = ActionMove(s.Stage, s.Stage.Hero, DirectionDown)
		case tcell.KeyLeft:
			action = ActionMove(s.Stage, s.Stage.Hero, DirectionLeft)
		case tcell.KeyRight:
			action = ActionMove(s.Stage, s.Stage.Hero, DirectionRight)
		}

		if action != nil {
			s.Stage.HeroAction(action)
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

	return box
}

package main

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type ScreensStage struct {
	Stage *Stage
}

func (s *ScreensStage) Do(g *Game, end func(next Screen)) tview.Primitive {

	hero := Actor{
		IsHero:   true,
		Position: Position{X: 5, Y: 5},
		Energy:   Energy{energyAction},
		Speed:    1,
		Behavior: nil,
	}

	s.Stage = NewStage(&hero)

	box := tview.NewBox()
	box.SetTitle("title")
	box.SetBorder(true)
	box.SetBackgroundColor(tcell.ColorGreen)

	box.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		var action Action

		switch event.Key() {
		case tcell.KeyUp:
			action = ActionMove(s.Stage, DirectionTop)
		case tcell.KeyDown:
			action = ActionMove(s.Stage, DirectionDown)
		case tcell.KeyLeft:
			action = ActionMove(s.Stage, DirectionLeft)
		case tcell.KeyRight:
			action = ActionMove(s.Stage, DirectionRight)
		}

		if action != nil {
			hero.nextAction(action)
			s.Stage.Update() // TODO: use in tick
		}

		return event
	})

	box.SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		x, y, width, height = box.GetInnerRect()

		box.SetTitle(fmt.Sprintf("%dx%d", s.Stage.Hero.Position.X, s.Stage.Hero.Position.Y))
		screen.SetContent(s.Stage.Hero.Position.X, s.Stage.Hero.Position.Y, '@', nil, tcell.StyleDefault)

		return x, y, width, height
	})

	return box
}

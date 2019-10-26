package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func NewScreenStage(g *Game, mapName string, location *rune) *ScreenStage {
	stage := NewStage(g)
	stage.Load(mapName, location)

	return &ScreenStage{
		Stage: stage,
	}
}

type ScreenStage struct {
	Stage *Stage
}

func (s *ScreenStage) Finalize() {
	SaveState(s.Stage)
	s.Stage.Stop()
}

var onceMap2 = sync.Once{}
var onceMap3 = sync.Once{}

func CreateThink(s *ScreenStage) {
	for i := 0; i < 10; i++ {
		time.Sleep(2000)
		s.Stage.AddActor(NewClassActor(s.Stage, Position{X: 78, Y: 13}, DirectionLeft, ClassThink))
	}
}

func (s *ScreenStage) Init(game *Game) tview.Primitive {
	// TODO: move to levels' meta data
	if s.Stage.Name == "map2" {
		onceMap2.Do(func() {
			s.Stage.AddActor(NewClassActor(s.Stage, Position{X: 22, Y: 6}, Direction{}, ClassDialog))
			s.Stage.AddActor(NewClassActor(s.Stage, Position{X: 46, Y: 6}, Direction{}, ClassBattle))
			s.Stage.AddActor(NewClassActor(s.Stage, Position{X: 7, Y: 5}, Direction{}, ClassPursue))
		})
		s.Stage.AddActor(NewClassActor(s.Stage, Position{X: 22, Y: 6}, Direction{}, ClassDialog))
		s.Stage.AddActor(NewClassActor(s.Stage, Position{X: 46, Y: 6}, Direction{}, ClassBattle))
		s.Stage.AddActor(NewClassActor(s.Stage, Position{X: 7, Y: 5}, Direction{}, ClassPursue))
	}

	if s.Stage.Name == "map3" {
		for y := 5; y <= 20; y = y + 3 {
			for x := 5; x <= 68; x = x + 3 {
				s.Stage.AddActor(NewClassActor(s.Stage, Position{X: x, Y: y}, Direction{}, ClassPursue))
			}
		}
		onceMap3.Do(func() {
			for y := 5; y <= 20; y = y + 3 {
				for x := 5; x <= 68; x = x + 3 {
					s.Stage.AddActor(NewClassActor(s.Stage, Position{X: x, Y: y}, Direction{}, ClassPursue))
				}
			}
		})
	}

	if s.Stage.Name == "mapMiniGame" {
		s.Stage.RegisterRune('4', Position{0, 13})
		s.Stage.RegisterRune('4', Position{29, 0})
		go CreateThink(s)
	}

	// stage stuff:

	lookAt := s.Stage.Hero.Position

	box := tview.NewBox()
	box.SetTitle("title")
	box.SetBorder(true)
	box.SetBackgroundColor(tcell.ColorGreen)

	box.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp:
			s.Stage.Hero.SetNextAction(ActionMove(s.Stage, s.Stage.Hero, DirectionTop))
			return nil
		case tcell.KeyDown:
			s.Stage.Hero.SetNextAction(ActionMove(s.Stage, s.Stage.Hero, DirectionDown))
			return nil
		case tcell.KeyLeft:
			s.Stage.Hero.SetNextAction(ActionMove(s.Stage, s.Stage.Hero, DirectionLeft))
			return nil
		case tcell.KeyRight:
			s.Stage.Hero.SetNextAction(ActionMove(s.Stage, s.Stage.Hero, DirectionRight))
			return nil
		case tcell.KeyCtrlA:
			if s.Stage.Name == "mapMiniGame" {
				s.Stage.RegisterRune('←', Position{s.Stage.Hero.Position.X + 1, s.Stage.Hero.Position.Y})
				s.Stage.AddActor(NewClassActor(s.Stage, Position{s.Stage.Hero.Position.X + 1, s.Stage.Hero.Position.Y}, Direction{}, ClassMirrorLeft))
			}
		case tcell.KeyCtrlD:
			if s.Stage.Name == "mapMiniGame" {
				s.Stage.RegisterRune('→', Position{s.Stage.Hero.Position.X + 1, s.Stage.Hero.Position.Y})
				s.Stage.AddActor(NewClassActor(s.Stage, Position{s.Stage.Hero.Position.X + 1, s.Stage.Hero.Position.Y}, Direction{}, ClassMirrorRight))
			}
		case tcell.KeyCtrlS:
			if s.Stage.Name == "mapMiniGame" {
				s.Stage.RegisterRune('↓', Position{s.Stage.Hero.Position.X + 1, s.Stage.Hero.Position.Y})
				s.Stage.AddActor(NewClassActor(s.Stage, Position{s.Stage.Hero.Position.X + 1, s.Stage.Hero.Position.Y}, Direction{}, ClassMirrorDown))
			}
		case tcell.KeyCtrlW:
			if s.Stage.Name == "mapMiniGame" {
				s.Stage.RegisterRune('↑', Position{s.Stage.Hero.Position.X + 1, s.Stage.Hero.Position.Y})
				s.Stage.AddActor(NewClassActor(s.Stage, Position{s.Stage.Hero.Position.X + 1, s.Stage.Hero.Position.Y}, Direction{}, ClassMirrorUp))
			}
		}

		switch event.Rune() {
		case ' ':
			s.Stage.Hero.SetNextAction(ActionInteract(s.Stage, s.Stage.Hero))
			return nil
		}

		return event
	})

	// Add helicopter effect.
	helicopterEffect := NewEffectHelicopter()
	s.Stage.Effects.Set("helicopter", helicopterEffect)

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
				if tile := s.Stage.Level.GetTile(mapPos); tile != nil {
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

		// Effects
		s.Stage.Effects.Update()
		for _, effect := range s.Stage.Effects.effects {
			s.drawEffect(port, screen, width, height, effect.Effect)
		}

		// HUD
		box.SetTitle(fmt.Sprintf("%dx%d", s.Stage.Hero.Position.X, s.Stage.Hero.Position.Y))

		return x, y, width, height
	})

	s.Stage.UpdateCallback = func(d time.Duration) bool {
		s.Stage.Effects.Update()
		if len(s.Stage.Effects.effects) > 0 {
			return true
		}
		return false
	}

	s.Stage.Start()

	return box
}

func (s *ScreenStage) drawTile(screen tcell.Screen, tile Tile, mapPos, screenPos Position) {
	tileRune, tileStyle := tile.GetAppearance(s.Stage.frame, Position{mapPos.X, mapPos.Y})
	screen.SetContent(screenPos.X, screenPos.Y, tileRune, nil, tileStyle)
}

func (s *ScreenStage) drawActor(screen tcell.Screen, actor *Actor, screenPos Position) {
	style := tcell.StyleDefault.Foreground(tcell.ColorRed)
	tile := s.Stage.Level.GetTile(actor.Position)
	_, tileStyle := tile.GetAppearance(s.Stage.frame, Position{screenPos.X, screenPos.Y})
	_, bg, _ := tileStyle.Decompose()
	style = style.Background(bg)
	screen.SetContent(screenPos.X, screenPos.Y, actor.Class.Rune, nil, style)
}

func (s *ScreenStage) drawEffect(port Port, screen tcell.Screen, width int, height int, effect Effect) {
	renderedEffect := effect.Render()
	for _, effectTile := range renderedEffect {
		screenPos := port.ToScreen(effectTile.Position)
		style := tcell.StyleDefault.Foreground(effectTile.Foreground)
		tile := s.Stage.Level.GetTile(effectTile.Position)
		if tile == nil {
			continue
		}
		_, tileStyle := tile.GetAppearance(s.Stage.frame, Position{screenPos.X, screenPos.Y})
		_, bg, _ := tileStyle.Decompose()
		style = style.Background(bg)
		screen.SetContent(screenPos.X, screenPos.Y, effectTile.Rune, nil, style)
	}
}

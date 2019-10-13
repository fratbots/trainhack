package main

import "log"

type Game struct {
	Sound    *SoundLibrary
	View     *View
	Screen   Screen
	Dialoger DialogManager

	State State
}

func NewGame() *Game {
	soundLibrary, err := NewSoundLibrary()
	if err != nil {
		log.Fatalf("Failed to init sound library: %v", err)
	}
	return &Game{
		Sound:    soundLibrary,
		View:     NewVew(),
		Dialoger: NewDialoger("./example/dialogs", "./example/bah.jpeg", "Иоган Себастья Бах"),

		State: State{Stages: map[string]StateStage{}},
	}
}

func (g *Game) Start(s Screen) {
	g.Sound.SetTheme(SoundThemePursuit)
	end := g.endCallback()
	p := s.Do(g, end)
	g.View.Run(p)
}

func (g *Game) endCallback() func(Screen) {
	var callback func(Screen)

	callback = func(next Screen) {
		if next != nil {
			g.Screen = next
			p := g.Screen.Do(g, callback)
			if p != nil {
				g.View.Focus(p)
				return
			}
		}
		g.Final()
	}

	return callback
}

func (g *Game) DoScreen(s Screen) {
	g.endCallback()(s)
}

func (g *Game) Final() {
	g.View.Final()
}

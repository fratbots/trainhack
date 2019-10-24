package main

type Game struct {
	UI UI

	State State

	Dialoger DialogManager
	Sound    *SoundLibrary
}

func NewGame() *Game {
	sound, _ := NewSoundLibrary()
	sound.SetTheme(SoundThemeAutumn)

	return &Game{
		UI: UI{},

		Dialoger: NewDialoger("./example/dialogs", "./example/hero.png", "Великий поработитель Уно"),
		Sound:    sound,
	}
}

func (g *Game) SetScreen(screen Screen) {
	g.UI.SetScreen(g, screen)
}

func (g *Game) Draw() {
	g.UI.Draw()
}

func (g *Game) Finish() {
	g.UI.Finalize()
}

/*
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
		Dialoger: NewDialoger("./example/dialogs", "./example/hero.png", "Великий поработитель Уно"),

		State: State{Stages: map[string]StateStage{}},
	}
}

func (g *Game) Start(s Screen) {
	g.Sound.SetTheme(SoundThemeAutumn)
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
*/

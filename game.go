package main

import "fmt"

type Game struct {
	UI UI

	State State

	Dialoger DialogManager
	Sound    SoundLibrary
}

func NewGame(playSounds bool) (*Game, error) {
	var sound SoundLibrary
	var err error
	if playSounds {
		sound, err = NewSoundLibraryDefault()
		if err != nil {
			return nil, fmt.Errorf("Failed to initialize sound library: %v", err)
		}
	} else {
		sound = NewSoundLibraryNoop()
	}
	sound.SetTheme(SoundThemeAutumn)

	return &Game{
		UI: UI{},

		Dialoger: NewDialoger("./example/dialogs", "./example/hero.png", "Великий поработитель Уно"),
		Sound:    sound,
	}, nil
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

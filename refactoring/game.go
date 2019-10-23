package main

type Game struct {
	UI UI
}

func NewGame(UI UI) *Game {
	return &Game{UI: UI}
}

func (g *Game) SetScreen(screen Screen) {
	g.UI.SetScreen(g, screen)
}

func (g *Game) Finalize() {
	g.UI.Finalize()
}

package main

type Game struct {
	View   *View
	Screen Screen
}

func NewGame() *Game {
	return &Game{
		View: NewVew(),
	}
}

func (g *Game) Start(s Screen) {
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
		g.View.Final()
	}

	return callback
}

func (g *Game) DoScreen(s Screen) {
	g.endCallback()(s)
}

func (g *Game) Final() {
	g.View.Final()
}

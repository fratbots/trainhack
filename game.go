package main

type Game struct {
	View     *View
	Screen   Screen
	Dialoger DialogManager
}

func NewGame() *Game {
	return &Game{
		View:     NewVew(),
		Dialoger: NewDialoger("./example/dialogs", "./example/bah.jpeg", "Иоган Себастья Бах"),
	}
}

func (g *Game) Start(s Screen) {
	g.View.Run(s.Do(g, g.callback()))
}

func (g *Game) callback() func(Screen) {
	var callback func(Screen)

	callback = func(next Screen) {
		if next == nil {
			return
		}
		g.Screen = next
		p := g.Screen.Do(g, callback)
		if p != nil {
			g.View.Focus(p)
			return
		}
		g.View.Final()
	}

	return callback
}

func (g *Game) DoScreen(s Screen) {
	g.callback()(s)
}

func (g *Game) Final() {
	g.View.Final()
}

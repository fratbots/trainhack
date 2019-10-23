package main

func main() {
	ui := NewUI()
	game := NewGame(ui)

	game.SetScreen(NewWelcomeScreen())
}

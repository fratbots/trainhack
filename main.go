package main

func main() {
	game := NewGame()
	//game.Start(&HelloScreen{})
	game.Start(&ScreenMap{})
}

package main

func main() {
	go Sound()
	game := NewGame()
	game.Start(&HelloScreen{})
	//game.Start(&ScreenMap{})
}

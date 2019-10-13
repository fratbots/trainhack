package main

import (
	"os"
)

func main() {
	go Sound()
	os.Setenv("TERM", "xterm-256color")
	game := NewGame()
	game.Start(&HelloScreen{})
	// game.Start(&ScreenMap{})
}

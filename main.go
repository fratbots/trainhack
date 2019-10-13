package main

import (
	"os"
)

func main() {
	os.Setenv("TERM", "xterm-256color")

	game := NewGame()
	game.Start(&HelloScreen{})
}

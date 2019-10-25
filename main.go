package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	os.Setenv("TERM", "xterm-256color")

	var flagSounds = flag.Bool("sounds", true, "Play sounds.")
	flag.Parse()

	game, err := NewGame(*flagSounds)
	if err != nil {
		log.Fatalf("Failed to initialize game: %v", err)
	}
	game.SetScreen(NewScreenHello())
}

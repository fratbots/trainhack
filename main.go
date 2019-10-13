package main

import (
	"log"
	"os"
	"time"
)

func main() {
	os.Setenv("TERM", "xterm-256color")

	// themeSoundsTest()

	game := NewGame()
	game.Start(&HelloScreen{})
}

func themeSoundsTest() {
	soundLibrary, err := NewSoundLibrary()
	if err != nil {
		log.Fatalf("Failed to init sound library: %v", err)
	}
	soundLibrary.SetTheme(SoundThemeAutumn)
	time.Sleep(time.Second * 2)
	soundLibrary.SetTheme(SoundThemePursuit)
	time.Sleep(time.Second * 2)
}

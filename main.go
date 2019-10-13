package main

import (
	"log"
	"os"
)

func main() {
	os.Setenv("TERM", "xterm-256color")
	soundTest()

	//game := NewGame()
	//game.Start(&HelloScreen{})
}

func soundTest() {
	snd, err := InitSound()
	if err != nil {
		log.Fatalf("Failed to init sound: %v", err)
	}
	snd.PlayTheme(SoundThemeAutumn)
}

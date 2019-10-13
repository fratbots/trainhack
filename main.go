package main

import (
	"log"
	"os"
	"time"
)

func main() {
	os.Setenv("TERM", "xterm-256color")

	//contextSoundsTest()
	themeSoundsTest()

	//game := NewGame()
	//game.Start(&HelloScreen{})
}

func contextSoundsTest() {
	soundLibrary, err := NewSoundLibrary()
	if err != nil {
		log.Fatalf("Failed to init sound library: %v", err)
	}
	soundLibrary.PlayContext(SoundContextDoor)
	time.Sleep(time.Second * 1)
	soundLibrary.PlayContext(SoundContextDoor)
	time.Sleep(time.Second * 1)
}

func themeSoundsTest() {
	soundLibrary, err := NewSoundLibrary()
	if err != nil {
		log.Fatalf("Failed to init sound library: %v", err)
	}
	soundLibrary.SetTheme(SoundThemePursuit)
	time.Sleep(time.Second * 5)
}

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

const (
	SoundDir          = "./music"
	SoundThemeAutumn  = "theme-autumn.mp3"
	SoundThemePursuit = "theme-pursuit.mp3"
	SoundContextDoor  = "context-door.mp3"
)

// Sound is a sound manager.
type Sound struct {
}

// NewSound returns new sound manager.
func NewSound() (*Sound, error) {
	if err != nil {
		return &Sound{}, fmt.Errorf("Failed to initialize speaker: %v", err)
	}
	s := &Sound{}
	err = s.loadSoundToBuffer(SoundContextDoor)
	if err != nil {
		return &Sound{}, fmt.Errorf("Failed to buffer sound file %s: %v", SoundContextDoor, err)
	}
	return s, nil
}

func loadSoundToBuffer(filename string) error {
	path := fmt.Printf("%s/%s", SoundDir, filename)
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Failed to open sound file %s: %v", path, err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return fmt.Errorf("Failed to init streamer for sound file %s: %v", path, err)
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()

	return nil
}

// playTheme infinitely plays theme soundtrack.
func (s *Sound) PlayTheme(theme string) error {
	// FIXME implement
	fmt.Printf("playTheme %s\n", theme)
	return nil
}

// playContext plays context sound once.
func (s *Sound) PlayContext(context string) error {
	// FIXME implement
	fmt.Printf("playContext %s\n", context)
	return nil
}

/*
func Sound() {
	fname := "./music/1.wav"
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("failed to open %s: %v", fname, err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatalf("failed to decode %s: %v", fname, err)
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	loop := beep.Loop(-1, streamer)
	done := make(chan bool)
	speaker.Play(beep.Seq(loop, beep.Callback(func() {
		done <- true
	})))
	<-done
}
*/

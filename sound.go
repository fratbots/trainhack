package main

import (
	"fmt"
	"os"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

const (
	SoundSampleRate        = 44100
	SoundSpeakerBufferSize = 10000

	SoundDir          = "./music"
	SoundThemeAutumn  = "theme-autumn"
	SoundThemePursuit = "theme-pursuit"
	SoundContextDoor  = "context-door"
)

// SoundLibrary is a collection of all the game sounds.
type SoundLibrary struct {
	themes        map[string]SoundTheme
	contextSounds map[string]SoundContext
}

type SoundTheme struct {
	ctrl *beep.Ctrl
}

func (s SoundTheme) Unpause() {
	s.ctrl.Paused = true
}

func (s SoundTheme) Pause() {
	s.ctrl.Paused = false
}

func (s SoundTheme) Paused() bool {
	return s.ctrl.Paused
}

type SoundContext struct {
	buffer *beep.Buffer
}

func (s SoundContext) Play() {
	streamer := s.buffer.Streamer(0, s.buffer.Len())
	speaker.Play(streamer)
}

// NewSoundLibrary creates a library containing all the preloaded game sounds.
func NewSoundLibrary() (*SoundLibrary, error) {
	// List game sounds: themes and context sounds.
	themeFiles := map[string]string{
		SoundThemeAutumn:  "theme-autumn.mp3",
		SoundThemePursuit: "theme-pursuit.mp3",
	}
	contextFiles := map[string]string{
		SoundContextDoor: "context-door.mp3",
	}

	soundLibrary := &SoundLibrary{
		themes:        make(map[string]SoundTheme, len(themeFiles)),
		contextSounds: make(map[string]SoundContext, len(contextFiles)),
	}

	// Init speaker.
	speaker.Init(SoundSampleRate, SoundSpeakerBufferSize)

	// Load themes.
	for soundID, filename := range themeFiles {
		err := soundLibrary.loadThemeSound(soundID, filename)
		if err != nil {
			return nil, fmt.Errorf("Failed to load theme sound %s: %v", soundID, err)
		}
	}
	// Load context sounds.
	for soundID, filename := range contextFiles {
		err := soundLibrary.loadContextSound(soundID, filename)
		if err != nil {
			return nil, fmt.Errorf("Failed to load context sound %s: %v", soundID, err)
		}
	}

	return soundLibrary, nil
}

// loadThemeSound reads the theme sound from file and stores it in a memory buffer.
func (l *SoundLibrary) loadThemeSound(soundID string, filename string) error {
	path := fmt.Sprintf("%s/%s", SoundDir, filename)
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Failed to open sound file %s: %v", path, err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return fmt.Errorf("Failed to init streamer for sound file %s: %v", path, err)
	}

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()
	loop := beep.Loop(-1, buffer.Streamer(0, buffer.Len()))
	ctrl := &beep.Ctrl{Streamer: loop, Paused: false}

	speaker.Play(ctrl)

	l.themes[soundID] = SoundTheme{
		ctrl: ctrl,
	}

	return nil
}

// loadContextSound reads the context sound from file and stores it in a memory buffer.
func (l *SoundLibrary) loadContextSound(soundID string, filename string) error {
	path := fmt.Sprintf("%s/%s", SoundDir, filename)
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Failed to open sound file %s: %v", path, err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return fmt.Errorf("Failed to init streamer for sound file %s: %v", path, err)
	}

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()

	l.contextSounds[soundID] = SoundContext{
		buffer: buffer,
	}

	return nil
}

// SetTheme starts to play theme sound infinitely.
func (l *SoundLibrary) SetTheme(soundID string) {
	for theme, sound := range l.themes {
		if theme != soundID && !sound.Paused() {
			fmt.Printf("pausing %s\n", soundID)
			l.themes[soundID].Pause()
		}
	}
	fmt.Printf("unpausing %s\n", soundID)
	l.themes[soundID].Unpause()
}

// PlayContext plays context sound once.
func (l *SoundLibrary) PlayContext(soundID string) {
	sound := l.contextSounds[soundID]
	sound.Play()
}

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
	SoundSpeakerBufferSize = 200

	SoundDir          = "./music"
	SoundThemeAutumn  = "theme-autumn"
	SoundThemePursuit = "theme-pursuit"
	SoundContextDoor  = "context-door"
)

// SoundLibrary is an interface of sound library.
type SoundLibrary interface {
	SetTheme(soundID string)
	PlayContext(soundID string)
	Pause()
}

// SoundLibraryNoop is an implementation of a sound library that just mimics an interface and does nothing.
type SoundLibraryNoop struct {
}

// NewSoundLibraryNoop returns a sound library that does nothing.
func NewSoundLibraryNoop() *SoundLibraryNoop {
	return &SoundLibraryNoop{}
}

// SetTheme does nothing by design.
func (s *SoundLibraryNoop) SetTheme(soundID string) {
	// does nothing
}

// PlayContext does nothing by design.
func (s *SoundLibraryNoop) PlayContext(soundID string) {
	// does nothing
}

func (s *SoundLibraryNoop) Pause() {
	// does nothing
}

// SoundLibraryDefault is a default sound library implementation that actually plays sounds.
type SoundLibraryDefault struct {
	themes        map[string]SoundTheme
	contextSounds map[string]SoundContext
}

// SoundTheme is a background music.
type SoundTheme struct {
	ctrl *beep.Ctrl
}

// Unpause resumes sound theme playback.
func (s SoundTheme) Unpause() {
	s.ctrl.Paused = false
}

// Pause temporarily stops sound theme playback.
func (s SoundTheme) Pause() {
	s.ctrl.Paused = true
}

// Paused returns pause state of sound theme.
func (s SoundTheme) Paused() bool {
	return s.ctrl.Paused
}

// SoundContext is a decently short sound that is played in case of specified event.
type SoundContext struct {
	buffer *beep.Buffer
}

// Play asynchronously plays context sound from start to end.
func (s SoundContext) Play() {
	streamer := s.buffer.Streamer(0, s.buffer.Len())
	speaker.Play(streamer)
}

// NewSoundLibraryDefault creates a sound library able to play sounds.
func NewSoundLibraryDefault() (*SoundLibraryDefault, error) {
	// List game sounds: themes and context sounds.
	themeFiles := map[string]string{
		SoundThemeAutumn:  "theme-autumn.mp3",
		SoundThemePursuit: "theme-pursuit.mp3",
	}
	contextFiles := map[string]string{
		SoundContextDoor: "context-door.mp3",
	}

	soundLibrary := &SoundLibraryDefault{
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
func (l *SoundLibraryDefault) loadThemeSound(soundID string, filename string) error {
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
	ctrl := &beep.Ctrl{Streamer: loop, Paused: true}

	speaker.Play(ctrl)

	l.themes[soundID] = SoundTheme{
		ctrl: ctrl,
	}

	return nil
}

// loadContextSound reads the context sound from file and stores it in a memory buffer.
func (l *SoundLibraryDefault) loadContextSound(soundID string, filename string) error {
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
func (l *SoundLibraryDefault) SetTheme(soundID string) {
	for theme, _ := range l.themes {
		l.themes[theme].Pause()
	}
	l.themes[soundID].Unpause()
}

func (l *SoundLibraryDefault) Pause() {
	for theme, _ := range l.themes {
		l.themes[theme].Pause()
	}
}

// PlayContext plays context sound once.
func (l *SoundLibraryDefault) PlayContext(soundID string) {
	sound := l.contextSounds[soundID]
	sound.Play()
}

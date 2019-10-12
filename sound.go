package main

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

func Sound() {
	fname := "./music/1.wav"
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("failed to open %s: %v", fname, err)
	}
	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatalf("failed to decode %s: %v", fname, err)
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	if true {
		speaker.Play(beep.Seq(streamer, beep.Callback(func() {
			done <- true
		})))
	}
	<-done
}

package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"log"
	"os"
	"time"
)

func main() {
	os.Setenv("TERM", "xterm-256color")

	// themeSoundsTest()
	//   lolkek()
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

func lolkek() {
	app := tview.NewApplication()
	frame := tview.NewFrame(tview.NewBox().SetBackgroundColor(tcell.ColorBlue)).
		SetBorders(2, 2, 2, 2, 4, 4).
		AddText("Header left", true, tview.AlignLeft, tcell.ColorWhite).
		AddText("Header middle", true, tview.AlignCenter, tcell.ColorWhite).
		AddText("Header right", true, tview.AlignRight, tcell.ColorWhite).
		AddText("Header second middle", true, tview.AlignCenter, tcell.ColorRed).
		AddText("Footer middle", false, tview.AlignCenter, tcell.ColorGreen).
		AddText("Footer second middle", false, tview.AlignCenter, tcell.ColorGreen)
	if err := app.SetRoot(frame, true).SetFocus(frame).Run(); err != nil {
		panic(err)
	}
}

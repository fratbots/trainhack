The following two code samples illustrate how one could write an application that shows the current time.

User [@lnxbil](https://github.com/lnxbil) provides us with an example of a [`Modal`](https://godoc.org/github.com/rivo/tview#Modal) whose text is changed in a goroutine to display the current time:

```go
// Demo code for a timer based update
package main

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
)

const refreshInterval = 500 * time.Millisecond

var (
	view *tview.Modal
	app  *tview.Application
)

func currentTimeString() string {
	t := time.Now()
	return fmt.Sprintf(t.Format("Current time is 15:04:05"))
}

func updateTime() {
	for {
		time.Sleep(refreshInterval)
		app.QueueUpdateDraw(func() {
			view.SetText(currentTimeString())
		})
	}
}

func main() {
	app = tview.NewApplication()
	view = tview.NewModal().
		SetText(currentTimeString()).
		AddButtons([]string{"Quit", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Quit" {
				app.Stop()
			}
		})

	go updateTime()
	if err := app.SetRoot(view, false).Run(); err != nil {
		panic(err)
	}
}
```

Similar in functionality but different in execution is the following example by [@ardnew](https://github.com/ardnew). Here, a screen refresh is triggered every 500 milliseconds and the current time is drawn in response to that screen refresh event. Calling `Application.QueueUpdate()` is not needed because `Application.Draw()` can be called from any goroutine. The example also shows how you may provide your own drawing routine.

```go
// Demo code for a timer based update
package main

import (
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const refreshInterval = 500 * time.Millisecond

var (
	view *tview.Box
	app  *tview.Application
)

func drawTime(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
	timeStr := time.Now().Format("Current time is 15:04:05")
	tview.Print(screen, timeStr, x, height/2, width, tview.AlignCenter, tcell.ColorLime)
	return 0, 0, 0, 0
}

func refresh() {
	tick := time.NewTicker(refreshInterval)
	for {
		select {
		case <-tick.C:
			app.Draw()
		}
	}
}

func main() {
	app = tview.NewApplication()
	view = tview.NewBox().SetDrawFunc(drawTime)

	go refresh()
	if err := app.SetRoot(view, true).Run(); err != nil {
		panic(err)
	}
}
```
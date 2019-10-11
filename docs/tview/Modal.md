For modal message boxes, there is a built-in primitive called `Modal`:

[[https://github.com/rivo/tview/blob/master/demos/modal/screenshot.png]]

Code:

```go
package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	modal := tview.NewModal().
		SetText("Do you want to quit the application?").
		AddButtons([]string{"Quit", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Quit" {
				app.Stop()
			}
		})
	if err := app.SetRoot(modal, false).SetFocus(modal).Run(); err != nil {
		panic(err)
	}
}
```

If you need more functionality than the `Modal` class provides, you can use any other primitive instead. To center your primitive, use the `Flex` class with the border elements set to `nil`. To make it appear on top of other elements, as with the `Modal` class itself, use the `Pages` class. Here is an example:

[[https://github.com/rivo/tview/blob/master/demos/modal/centered.png]]

```go
package main

import (
	"strings"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	// Returns a new primitive which puts the provided primitive in the center and
	// sets its size to the given width and height.
	modal := func(p tview.Primitive, width, height int) tview.Primitive {
		return tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(p, height, 1, false).
				AddItem(nil, 0, 1, false), width, 1, false).
			AddItem(nil, 0, 1, false)
	}

	background := tview.NewTextView().
		SetTextColor(tcell.ColorBlue).
		SetText(strings.Repeat("background ", 1000))

	box := tview.NewBox().
		SetBorder(true).
		SetTitle("Centered Box")

	pages := tview.NewPages().
		AddPage("background", background, true, true).
		AddPage("modal", modal(box, 40, 10), true, true)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
```

The same effect can be achieved by using the `Grid` class instead of `Flex`:

```go
// Returns a new primitive which puts the provided primitive in the center and
// sets its size to the given width and height.
modal := func(p tview.Primitive, width, height int) tview.Primitive {
	return tview.NewGrid().
		SetColumns(0, width, 0).
		SetRows(0, height, 0).
		AddItem(p, 1, 1, 1, 1, 0, 0, true)
}
```
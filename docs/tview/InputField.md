A one line InputField:

[[https://github.com/rivo/tview/blob/master/demos/inputfield/screenshot.png]]

Code:

```go
package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	inputField := tview.NewInputField().
		SetLabel("Enter a number: ").
		SetFieldWidth(10).
		SetAcceptanceFunc(tview.InputFieldInteger).
		SetDoneFunc(func(key tcell.Key) {
			app.Stop()
		})
	if err := app.SetRoot(inputField, true).SetFocus(inputField).Run(); err != nil {
		panic(err)
	}
}
```

The `InputField` class also provides autocomplete functionality. The repo includes two examples for this. The first example, [`autocomplete.go`](https://github.com/rivo/tview/blob/master/demos/inputfield/autocomplete.go), illustrates synchronous autocomplete functionality where the data is readily available in the same goroutine.

The second example, [`autocompleteasync.go`](https://github.com/rivo/tview/blob/master/demos/inputfield/autocompleteasync.go), illustrates how autocomplete functionality can be implemented when the autocomplete entries need to be retrieved in a separate goroutine, e.g. by querying an external API.
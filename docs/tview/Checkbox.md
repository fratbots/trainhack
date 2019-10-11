A simple Checkbox:

[[https://github.com/rivo/tview/blob/master/demos/checkbox/screenshot.png]]

Code:

```go
package main

import "github.com/rivo/tview"

func main() {
	app := tview.NewApplication()
	checkbox := tview.NewCheckbox().SetLabel("Hit Enter to check box: ")
	if err := app.SetRoot(checkbox, true).SetFocus(checkbox).Run(); err != nil {
		panic(err)
	}
}
```
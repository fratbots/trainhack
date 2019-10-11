A simple Box with a title:

[[https://github.com/rivo/tview/blob/master/demos/box/screenshot.png]]

Code:

```go
package main

import "github.com/rivo/tview"

func main() {
	box := tview.NewBox().
		SetBorder(true).
		SetTitle("Box Demo")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
```
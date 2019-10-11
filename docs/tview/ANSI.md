`tview` provides the functions [`ANSIWriter()`](https://godoc.org/github.com/rivo/tview#ANSIWriter) and [`TranslateANSI()`](https://godoc.org/github.com/rivo/tview#TranslateANSI) which translate common ANSI escape sequences into color tags used by `tview`. This can be useful when you want to display colorized output from other programs in `tview`.

Here is an example that displays output from `stdin`:

```go
package main

import (
	"io"
	"os"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	textView.SetBorder(true).SetTitle("Stdin")
	go func() {
		w := tview.ANSIWriter(textView)
		if _, err := io.Copy(w, os.Stdin); err != nil {
			panic(err)
		}
	}()
	if err := app.SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}
}
```

When you run this program, you must pipe text into it. To illustrate the handling of ANSI escape sequences, you could do the following on Linux:

```bash
ls -l --color=always | go run main.go
```

Any colored text produced by the `ls` command will be displayed in the same color in `tview`.
package main

import (
	"github.com/rivo/tview"
)

func NewUIModal(text string, buttons ...interface{}) tview.Primitive {
	p := tview.NewModal()
	p.SetText(text)

	var functions []func()

	for i := 0; i < len(buttons); i += 2 {
		if s, ok := buttons[i].(string); ok {
			p.AddButtons([]string{s})
		}
		if p, ok := buttons[i+1].(func()); ok {
			functions = append(functions, p)
		}
	}

	p.SetDoneFunc(func(i int, l string) {
		if i < len(functions) {
			functions[i]()
		}
	})

	return p
}

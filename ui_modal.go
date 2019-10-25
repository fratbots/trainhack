package main

import (
	"github.com/rivo/tview"
)

func NewUIModal(text string, buttons ...interface{}) tview.Primitive {
	p := tview.NewModal()
	p.SetText(text)

	var functions []func()
	var texts []string

	for i := 0; i < len(buttons); i += 2 {
		if s, ok := buttons[i].(string); ok {
			texts = append(texts, s)
		}
		if f, ok := buttons[i+1].(func()); ok {
			functions = append(functions, f)
		}
	}

	p.AddButtons(texts)
	p.SetDoneFunc(func(i int, l string) {
		if i < len(functions) {
			functions[i]()
		}
	})

	return p
}

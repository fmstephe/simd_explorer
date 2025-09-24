package stackapp

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type StackApp struct {
	app   *tview.Application
	stack []tview.Primitive
}

func NewStackApp() *StackApp {
	a := &StackApp{
		app:   tview.NewApplication(),
		stack: []tview.Primitive{},
	}
	// Enable mouse capture
	a.app.EnableMouse(true)
	// 'q' or ESC always quits the application
	a.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyESC:
			a.app.Stop()
		}
		switch event.Rune() {
		case 'q':
			a.app.Stop()
		}
		return event
	})

	return a
}

func (a *StackApp) Push(p tview.Primitive) {
	a.stack = append(a.stack, p)

	a.setRoot()
}

func (a *StackApp) Pop() {
	l := len(a.stack)
	if l == 1 {
		// Popping the last Primitive is a noop
		return
	}

	// Remove reference to avoid memmory leaks
	a.stack[l-1] = nil
	a.stack = a.stack[:l-1]

	a.setRoot()
}

func (a *StackApp) SetFocus(p tview.Primitive) {
	a.app.SetFocus(p)
}

func (a *StackApp) Run() error {
	return a.app.Run()
}

func (a *StackApp) setRoot() {
	a.app.SetRoot(a.stack[len(a.stack)-1], true)
}

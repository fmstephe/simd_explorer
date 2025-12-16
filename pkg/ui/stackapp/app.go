package stackapp

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type StackApp struct {
	app   *tview.Application
	idx   int
	stack []tview.Primitive
}

func NewStackApp() *StackApp {
	a := &StackApp{
		app:   tview.NewApplication(),
		idx:   -1,
		stack: []tview.Primitive{},
	}
	// Disable mouse capture
	a.app.EnableMouse(false)
	// 'q' or ESC always quits the application
	a.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyESC:
			a.app.Stop()
		case tcell.KeyF1:
			a.Pop()
		case tcell.KeyF2:
			a.Recover()
		}
		return event
	})

	return a
}

func (a *StackApp) Push(p tview.Primitive) {
	a.clearRecover()
	a.stack = append(a.stack, p)
	a.idx++

	a.setRoot()
}

func (a *StackApp) Pop() {
	// This method is a noop for an unused StackApp
	if a.idx == -1 {
		return
	}

	// If we are not at the bottom of the stack, go down one element. NB:
	// This preserves the current primitive as a recoverable historical
	// primitive, until clearRecover() is called.
	if a.idx != 0 {
		a.idx--
	}
	a.setRoot()
}

func (a *StackApp) Recover() {
	// This method is a noop for an unused StackApp
	if a.idx == -1 {
		return
	}

	// If recoverable primitives live above our current level in the stack then go up one level and return to that previous primitive.
	if a.idx < len(a.stack)-1 {
		a.idx++
	}

	a.setRoot()
}

func (a *StackApp) SetFocus(p tview.Primitive) {
	a.app.SetFocus(p)
}

func (a *StackApp) Run() error {
	return a.app.Run()
}

// This resizes the stack to only include primitives up to and including the
// current primitive. All recovery primitives are lost.
func (a *StackApp) clearRecover() {
	// This method is a noop for an unused StackApp
	if a.idx == -1 {
		return
	}

	for i := a.idx + 1; i < len(a.stack); i++ {
		a.stack[i] = nil
	}

	a.stack = a.stack[:a.idx+1]
}

func (a *StackApp) setRoot() {
	// This method is a noop for an unused StackApp
	if a.idx == -1 {
		return
	}

	a.app.SetRoot(a.stack[a.idx], true)
}

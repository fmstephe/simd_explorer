package ui

import (
	"github.com/fmstephe/simd_explorer/pkg/ui/register"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Run() {
	app := tview.NewApplication()
	app.EnableMouse(true)

	register256 := register.NewUIRegisterSet(app, 256)
	primitive := register256.Base10.GetPrimitive()
	// Setup the application with the components defined above

	app.SetRoot(primitive, true)
	app.SetFocus(primitive)
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyESC:
			app.Stop()
		}
		switch event.Rune() {
		case 'q':
			app.Stop()
		}
		return event
	})

	if err := app.Run(); err != nil {
		panic(err)
	}
}

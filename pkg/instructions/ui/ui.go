package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Run() {
	app := tview.NewApplication()

	cBroadcaster := newChangeBroadcaster()

	output64 := NewRegisterOutputs(app, 64, 256)
	cBroadcaster.addReceiver(output64)

	input64 := NewRegisterInputs(app, 64, 256, cBroadcaster)
	cBroadcaster.addReceiver(input64)

	grid := tview.NewGrid()
	grid.SetRows(1, 1)
	grid.SetColumns(0, 0)
	grid.SetBorders(true)
	grid.AddItem(input64.GetBox(), 0, 0, 1, 1, 0, 0, true)
	grid.AddItem(output64.GetBox(), 0, 1, 1, 1, 0, 0, false)

	// Setup the application with the components defined above
	app.SetRoot(grid, true)
	app.SetFocus(grid)
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

package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Run() {
	app := tview.NewApplication()

	cBroadcaster := newChangeBroadcaster(256)

	input64 := NewRegisterInputs(app, 64, 256, 16, cBroadcaster)
	output64 := NewRegisterOutputs(app, 64, 256, 16)

	input32 := NewRegisterInputs(app, 32, 256, 16, cBroadcaster)
	output32 := NewRegisterOutputs(app, 32, 256, 16)

	input16 := NewRegisterInputs(app, 16, 256, 16, cBroadcaster)
	output16 := NewRegisterOutputs(app, 16, 256, 16)

	input8 := NewRegisterInputs(app, 8, 256, 16, cBroadcaster)
	output8 := NewRegisterOutputs(app, 8, 256, 16)

	// Add update receivers, now that all initialisation updates have completed
	cBroadcaster.addReceiver(input64)
	cBroadcaster.addReceiver(output64)

	cBroadcaster.addReceiver(input32)
	cBroadcaster.addReceiver(output32)

	cBroadcaster.addReceiver(input16)
	cBroadcaster.addReceiver(output16)

	cBroadcaster.addReceiver(input8)
	cBroadcaster.addReceiver(output8)

	cBroadcaster.broadcastZeros()

	grid := tview.NewGrid()
	grid.SetRows(1, 1, 1, 1)
	grid.SetColumns(0, 0)
	grid.SetBorders(true)

	grid.AddItem(input64.GetBox(), 0, 0, 1, 1, 0, 0, true)
	grid.AddItem(output64.GetBox(), 0, 1, 1, 1, 0, 0, false)

	grid.AddItem(input32.GetBox(), 1, 0, 1, 1, 0, 0, false)
	grid.AddItem(output32.GetBox(), 1, 1, 1, 1, 0, 0, false)

	grid.AddItem(input16.GetBox(), 2, 0, 1, 1, 0, 0, false)
	grid.AddItem(output16.GetBox(), 2, 1, 1, 1, 0, 0, false)

	grid.AddItem(input8.GetBox(), 3, 0, 1, 1, 0, 0, false)
	grid.AddItem(output8.GetBox(), 3, 1, 1, 1, 0, 0, false)

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

package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Run() {
	app := tview.NewApplication()

	/*
		// Create text display
		text := tview.NewTextView().
			SetChangedFunc(func() { app.Draw() })

		// Create number input
		inputField := tview.NewInputField()
		inputField.SetLabel("Enter a number: ").
			SetFieldWidth(10).
			SetAcceptanceFunc(tview.InputFieldInteger).
			SetDoneFunc(func(key tcell.Key) { text.SetText(inputField.GetText()) })

		// Create flex to hold the input and text
		flex := tview.NewFlex().
			AddItem(inputField, 0, 1, true).
			AddItem(text, 0, 1, false)
	*/

	cBroadcaster := newChangeBroadcaster()

	output := NewRegisterOutputs(app, 64, 256)
	cBroadcaster.addReceiver(output)

	input := NewRegisterInputs(app, 64, 256, cBroadcaster)
	cBroadcaster.addReceiver(input)

	// Create flex to hold the input and text
	flex := tview.NewFlex().
		AddItem(input.GetBox(), 0, 1, true).
		AddItem(output.GetBox(), 0, 1, false)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyESC:
			app.Stop()
		}
		return event
	})

	// Setup the application with the components defined above
	app.SetRoot(flex, true).
		SetFocus(flex)

	if err := app.Run(); err != nil {
		panic(err)
	}
}

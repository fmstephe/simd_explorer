package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type input4X64 struct {
	app *tview.Application

	focus          int
	inputsForFocus []*tview.InputField

	input  *tview.Flex
	output *tview.Flex
}

func build4x64bitInputs(app *tview.Application) (input *input4X64) {
	// Create text display
	text1 := tview.NewTextView().
		SetChangedFunc(func() { app.Draw() })
	text2 := tview.NewTextView().
		SetChangedFunc(func() { app.Draw() })
	text3 := tview.NewTextView().
		SetChangedFunc(func() { app.Draw() })
	text4 := tview.NewTextView().
		SetChangedFunc(func() { app.Draw() })

	outputFlex := tview.NewFlex()
	outputFlex.AddItem(text1, 0, 1, false)
	outputFlex.AddItem(text2, 0, 1, false)
	outputFlex.AddItem(text3, 0, 1, false)
	outputFlex.AddItem(text4, 0, 1, false)

	// Create number input
	input1 := tview.NewInputField()
	input1.SetFieldWidth(20)
	input1.SetAcceptanceFunc(tview.InputFieldInteger)
	input1.SetChangedFunc(func(text string) { text1.SetText(text) })
	input1.SetDoneFunc(func(key tcell.Key) { text1.SetText(input1.GetText()) })

	// Create number input
	input2 := tview.NewInputField()
	input2.SetFieldWidth(20)
	input2.SetAcceptanceFunc(tview.InputFieldInteger)
	input2.SetChangedFunc(func(text string) { text2.SetText(text) })
	input2.SetDoneFunc(func(key tcell.Key) { text1.SetText(input1.GetText()) })

	// Create number input
	input3 := tview.NewInputField()
	input3.SetFieldWidth(20)
	input3.SetAcceptanceFunc(tview.InputFieldInteger)
	input3.SetChangedFunc(func(text string) { text3.SetText(text) })
	input3.SetDoneFunc(func(key tcell.Key) { text1.SetText(input1.GetText()) })

	// Create number input
	input4 := tview.NewInputField()
	input4.SetFieldWidth(20)
	input4.SetAcceptanceFunc(tview.InputFieldInteger)
	input4.SetChangedFunc(func(text string) { text4.SetText(text) })
	input4.SetDoneFunc(func(key tcell.Key) { text1.SetText(input1.GetText()) })

	inputFlex := tview.NewFlex()
	inputFlex.AddItem(input1, 0, 1, true)
	inputFlex.AddItem(input2, 0, 1, true)
	inputFlex.AddItem(input3, 0, 1, true)
	inputFlex.AddItem(input4, 0, 1, true)

	input = &input4X64{
		app:            app,
		focus:          0,
		inputsForFocus: []*tview.InputField{input1, input2, input3, input4},
		input:          inputFlex,
		output:         outputFlex,
	}
	input.InitCapture()
	return input
}

func (i *input4X64) GetInputOutput() (input, output *tview.Flex) {
	return i.input, i.output
}

func (i *input4X64) InitCapture() {
	i.input.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			i.cycleFocus(1)
		case tcell.KeyBacktab:
			i.cycleFocus(-1)
		}

		return event
	})
}

func (i *input4X64) cycleFocus(move int) {
	i.focus += move
	idx := i.focus % 4
	if idx < 0 {
		idx = 4 + idx
	}
	i.app.SetFocus(i.inputsForFocus[idx])
}

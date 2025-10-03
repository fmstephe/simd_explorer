package uiio

import (
	"github.com/fmstephe/simd_explorer/pkg/assembly"
	"github.com/fmstephe/simd_explorer/pkg/ui/stackapp"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UIInstruction struct {
	instruction       assembly.Instruction
	inputUIParameters []*UIParameterParts
	outputUIParameter *UIParameterParts
	box               *tview.Grid

	focus      int
	selectable []tview.Primitive
	app        *stackapp.StackApp
}

func NewUIInstruction(app *stackapp.StackApp, instruction assembly.Instruction) *UIInstruction {
	// UIRegister is required for callbacks in register input components.
	// When the input components are changed they callback into the
	// UIRegister to indicate that a value has been changed and
	// inputs/outputs reprocessed and broadcast.
	//
	// TODO this design _feels_ awkward, so we should have a think about
	// this in the future
	uiInst := &UIInstruction{
		instruction: instruction,
		focus:       0,
		selectable:  []tview.Primitive{},
		app:         app,
	}

	inputs := []*UIParameterParts{}
	for _, param := range instruction.Inputs() {
		input := NewUIParameterInputs(app, param, uiInst)
		uiInst.selectable = append(uiInst.selectable, input.selectablePrimitives()...)
		inputs = append(inputs, input)
	}

	output := NewUIParameterOutputs(app, instruction.Output(), uiInst)

	gridLeft := tview.NewGrid()
	gridLeft.SetBorder(true)
	gridLeft.SetTitle("Inputs")

	gridRight := tview.NewGrid()
	gridRight.SetBorder(true)
	gridRight.SetTitle("Outputs Base %d")

	for i, input := range inputs {
		gridLeft.AddItem(input.GetBox(), i, 0, 1, 1, 0, 0, true)
	}

	gridRight.AddItem(output.GetBox(), 0, 0, 1, 1, 0, 0, false)

	grid := tview.NewGrid()
	grid.AddItem(gridLeft, 0, 0, 1, 1, 0, 0, true)
	grid.AddItem(gridRight, 0, 1, 1, 1, 0, 0, false)

	// Fill out the fields for the UIRegister
	uiInst.inputUIParameters = inputs
	uiInst.outputUIParameter = output
	uiInst.box = grid

	// Setup the tab focus cycling (is there a better way to approach
	// this?)
	uiInst.initFocusCycling()

	// Set output fields from zeroed inputs
	uiInst.inputsChanged()

	return uiInst
}

func (r *UIInstruction) GetPrimitive() tview.Primitive {
	return r.box
}

func (r *UIInstruction) inputsChanged() {
	inputs := [][]byte{}
	for _, input := range r.inputUIParameters {
		inputs = append(inputs, input.getData())
	}
	output := r.instruction.Run(inputs)
	r.outputUIParameter.setData(output)
}

func (r *UIInstruction) initFocusCycling() {
	r.box.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			r.cycleFocus(1)
		case tcell.KeyBacktab:
			r.cycleFocus(-1)
		}

		return event
	})
}

func (r *UIInstruction) cycleFocus(move int) {
	r.focus += move
	idx := r.focus % len(r.selectable)
	if idx < 0 {
		idx = len(r.selectable) + idx
	}
	r.app.SetFocus(r.selectable[idx])
}

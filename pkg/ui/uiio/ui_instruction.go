package uiio

import (
	"fmt"

	"github.com/fmstephe/simd_explorer/pkg/assembly"
	"github.com/fmstephe/simd_explorer/pkg/ui/stackapp"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// This type is now obsolete, because we don't allow changing the base system
// for inputs/outputs as we did previously This decision was made as it became
// clearer that many (most, all?) inputs/outputs make the most sense with a
// specific base system If you are adding two sets of floats, no other base
// system is likely useful for representing the contents of the result. We'll
// leave the struct here for a short while, in case some new observation
// reverses everything. But most likely it will go away in the near future
type UIParametersSet struct {
	inst   assembly.Instruction
	Base2  *UIInstruction
	Base10 *UIInstruction
	Base16 *UIInstruction
}

func NewUIParametersSet(app *stackapp.StackApp, inst assembly.Instruction) *UIParametersSet {
	rs := &UIParametersSet{
		Base2:  NewUIParameters(app, inst),
		Base10: NewUIParameters(app, inst),
		Base16: NewUIParameters(app, inst),
	}

	return rs
}

type UIInstruction struct {
	instruction       assembly.Instruction
	inputUIParameters []*UIParameterParts
	outputUIParameter *UIParameterParts
	box               *tview.Grid

	focus      int
	selectable []tview.Primitive
	app        *stackapp.StackApp
}

func NewUIParameters(app *stackapp.StackApp, instruction assembly.Instruction) *UIInstruction {
	// UIRegister is required for callbacks in register input components.
	// When the input components are changed they callback into the
	// UIRegister to indicate that a value has been changed and
	// inputs/outputs reprocessed and broadcast.
	//
	// TODO this design _feels_ awkward, so we should have a think about
	// this in the future
	r := &UIInstruction{
		instruction: instruction,
		focus:       0,
		selectable:  []tview.Primitive{},
		app:         app,
	}

	inputs := []*UIParameterParts{}
	for _, param := range instruction.Inputs() {
		input := NewUIParameterInputs(app, param, r)
		r.selectable = append(r.selectable, input.selectablePrimitives()...)
		inputs = append(inputs, input)
	}

	output := NewUIParameterOutputs(app, instruction.Output(), r)

	gridLeft := tview.NewGrid()
	gridLeft.SetBorder(true)
	// TODO that's very fragile, need a better way to capture the base, or don't display it in this part of the UI?
	gridLeft.SetTitle(fmt.Sprintf("Inputs Base %d", instruction.Inputs()[0].Base()))

	gridRight := tview.NewGrid()
	gridRight.SetBorder(true)
	gridRight.SetTitle(fmt.Sprintf("Outputs Base %d", instruction.Output().Base()))

	for i, input := range inputs {
		gridLeft.AddItem(input.GetBox(), i, 0, 1, 1, 0, 0, true)
	}

	gridRight.AddItem(output.GetBox(), 0, 0, 1, 1, 0, 0, false)

	grid := tview.NewGrid()
	grid.AddItem(gridLeft, 0, 0, 1, 1, 0, 0, true)
	grid.AddItem(gridRight, 0, 1, 1, 1, 0, 0, false)

	// Fill out the fields for the UIRegister
	r.inputUIParameters = inputs
	r.outputUIParameter = output
	r.box = grid

	// Setup the tab focus cycling (is there a better way to approach
	// this?)
	r.initFocusCycling()

	// Set output fields from zeroed inputs
	r.inputsChanged()

	return r
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

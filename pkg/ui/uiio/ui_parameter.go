package uiio

import (
	"fmt"

	"github.com/fmstephe/simd_explorer/pkg/assembly"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
	"github.com/fmstephe/simd_explorer/pkg/ui/stackapp"
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
	Base2  *UIParameters
	Base10 *UIParameters
	Base16 *UIParameters
}

func NewUIParametersSet(app *stackapp.StackApp, inst assembly.Instruction) *UIParametersSet {
	cBroadcaster := newChangeBroadcaster(inst)

	rs := &UIParametersSet{
		Base2:  NewUIParameters(app, inst.Inputs(), inst.Output(), cBroadcaster),
		Base10: NewUIParameters(app, inst.Inputs(), inst.Output(), cBroadcaster),
		Base16: NewUIParameters(app, inst.Inputs(), inst.Output(), cBroadcaster),
	}

	// Set all parts to have 0 values
	cBroadcaster.broadcastZeros()

	return rs
}

type UIParameters struct {
	inputUIParameters []*UIParameterParts
	outputUIParameter *UIParameterParts
	cBroadcaster      *changeBroadcaster
	box               tview.Primitive
}

func NewUIParameters(app *stackapp.StackApp, inputParameters []*number.Parameter, outputParameter *number.Parameter, cBroadcaster *changeBroadcaster) *UIParameters {
	// UIRegister is required for callbacks in register input components.
	// When the input components are changed they callback into the
	// UIRegister to indicate that a value has been changed and
	// inputs/outputs reprocessed and broadcast.
	//
	// TODO this design _feels_ awkward, so we should have a think about
	// this in the future
	r := &UIParameters{}

	inputs := []*UIParameterParts{}
	for _, param := range inputParameters {
		input := NewUIParameterInputs(app, param, r)
		inputs = append(inputs, input)
	}

	output := NewUIParameterOutputs(app, outputParameter, r)

	gridLeft := tview.NewGrid()
	gridLeft.SetBorder(true)
	// TODO that's very fragile, need a better way to capture the base, or don't display it in this part of the UI?
	gridLeft.SetTitle(fmt.Sprintf("Inputs Base %d", inputParameters[0].Base()))

	gridRight := tview.NewGrid()
	gridRight.SetBorder(true)
	gridRight.SetTitle(fmt.Sprintf("Outputs Base %d", outputParameter.Base()))

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
	r.cBroadcaster = cBroadcaster
	r.box = grid

	// Add this UIRegister to the change broadcaster
	cBroadcaster.addReceiver(r)

	return r
}

func (r *UIParameters) GetPrimitive() tview.Primitive {
	return r.box
}

func (r *UIParameters) setData(inputs [][]byte, output []byte) {
	for i, input := range r.inputUIParameters {
		input.setData(inputs[i])
	}
	r.outputUIParameter.setData(output)
}

func (r *UIParameters) inputsChanged() {
	inputs := [][]byte{}
	for _, input := range r.inputUIParameters {
		inputs = append(inputs, input.getData())
	}
	r.cBroadcaster.broadcastChange(inputs)
}

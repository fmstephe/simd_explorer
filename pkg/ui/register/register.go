package register

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
type UIRegisterSet struct {
	inst   assembly.Instruction
	Base2  *UIRegister
	Base10 *UIRegister
	Base16 *UIRegister
}

func NewUIRegisterSet(app *stackapp.StackApp, inst assembly.Instruction) *UIRegisterSet {
	cBroadcaster := newChangeBroadcaster(inst)

	rs := &UIRegisterSet{
		Base2:  NewUIRegister(app, inst.Inputs(), inst.Output(), cBroadcaster),
		Base10: NewUIRegister(app, inst.Inputs(), inst.Output(), cBroadcaster),
		Base16: NewUIRegister(app, inst.Inputs(), inst.Output(), cBroadcaster),
	}

	// Set all parts to have 0 values
	cBroadcaster.broadcastZeros()

	return rs
}

type UIRegister struct {
	inputRegisters []*RegisterParts
	outputRegister *RegisterParts
	cBroadcaster   *changeBroadcaster
	box            tview.Primitive
}

func NewUIRegister(app *stackapp.StackApp, inputParameters []*number.Parameter, outputParameters *number.Parameter, cBroadcaster *changeBroadcaster) *UIRegister {
	// UIRegister is required for callbacks in register input components.
	// When the input components are changed they callback into the
	// UIRegister to indicate that a value has been changed and
	// inputs/outputs reprocessed and broadcast.
	//
	// TODO this design _feels_ awkward, so we should have a think about
	// this in the future
	r := &UIRegister{}

	inputs := []*RegisterParts{}
	for _, param := range inputParameters {
		input := NewRegisterInputs(app, param, r)
		inputs = append(inputs, input)
	}

	output := NewRegisterOutputs(app, outputParameters, r)

	gridLeft := tview.NewGrid()
	gridLeft.SetBorder(true)
	// TODO that's very fragile, need a better way to capture the base, or don't display it in this part of the UI?
	gridLeft.SetTitle(fmt.Sprintf("Inputs Base %d", inputParameters[0].Base()))

	gridRight := tview.NewGrid()
	gridRight.SetBorder(true)
	gridRight.SetTitle(fmt.Sprintf("Outputs Base %d", outputParameters.Base()))

	for i, input := range inputs {
		gridLeft.AddItem(input.GetBox(), i, 0, 1, 1, 0, 0, true)
	}

	gridRight.AddItem(output.GetBox(), 0, 0, 1, 1, 0, 0, false)

	grid := tview.NewGrid()
	grid.AddItem(gridLeft, 0, 0, 1, 1, 0, 0, true)
	grid.AddItem(gridRight, 0, 1, 1, 1, 0, 0, false)

	// Fill out the fields for the UIRegister
	r.inputRegisters = inputs
	r.outputRegister = output
	r.cBroadcaster = cBroadcaster
	r.box = grid

	// Add this UIRegister to the change broadcaster
	cBroadcaster.addReceiver(r)

	return r
}

func (r *UIRegister) GetPrimitive() tview.Primitive {
	return r.box
}

func (r *UIRegister) setData(inputs [][]byte, output []byte) {
	for i, input := range r.inputRegisters {
		input.setData(inputs[i])
	}
	r.outputRegister.setData(output)
}

func (r *UIRegister) inputsChanged() {
	inputs := [][]byte{}
	for _, input := range r.inputRegisters {
		inputs = append(inputs, input.getData())
	}
	r.cBroadcaster.broadcastChange(inputs)
}

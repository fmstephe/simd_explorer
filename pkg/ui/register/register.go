package register

import (
	"fmt"

	"github.com/fmstephe/simd_explorer/pkg/assembly"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
	"github.com/fmstephe/simd_explorer/pkg/ui/stackapp"
	"github.com/rivo/tview"
)

type UIRegisterSet struct {
	inst   assembly.Instruction
	Base2  *UIRegister
	Base10 *UIRegister
	Base16 *UIRegister
}

func NewUIRegisterSet(app *stackapp.StackApp, inst assembly.Instruction) *UIRegisterSet {
	cBroadcaster := newChangeBroadcaster(inst)

	rs := &UIRegisterSet{
		Base2:  NewUIRegister(app, 2, inst.InputSizes(), inst.OutputSize(), cBroadcaster),
		Base10: NewUIRegister(app, 10, inst.InputSizes(), inst.OutputSize(), cBroadcaster),
		Base16: NewUIRegister(app, 16, inst.InputSizes(), inst.OutputSize(), cBroadcaster),
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

func NewUIRegister(app *stackapp.StackApp, base int, inputSizes []number.Converter, outputSize number.Converter, cBroadcaster *changeBroadcaster) *UIRegister {
	// UIRegister is required for callbacks in register input components.
	// When the input components are changed they callback into the
	// UIRegister to indicate that a value has been changed and
	// inputs/outputs reprocessed and broadcast.
	//
	// TODO this design _feels_ awkward, so we should have a think about
	// this in the future
	r := &UIRegister{}

	inputs := []*RegisterParts{}
	for _, inputSize := range inputSizes {
		mustValidInputOutputSize(inputSize.GetBitWidth())
		inputPartSize := getPartSize(inputSize.GetBitWidth())
		input := NewRegisterInputs(app, inputPartSize, inputSize, r)
		inputs = append(inputs, input)
	}

	mustValidInputOutputSize(outputSize.GetBitWidth())
	outputPartSize := getPartSize(outputSize.GetBitWidth())

	output := NewRegisterOutputs(app, outputPartSize, outputSize, r)

	gridLeft := tview.NewGrid()
	gridLeft.SetBorder(true)
	gridLeft.SetTitle(fmt.Sprintf("Inputs Base %d", base))

	gridRight := tview.NewGrid()
	gridRight.SetBorder(true)
	gridRight.SetTitle(fmt.Sprintf("Outputs Base %d", base))

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

// TODO this likely isn't the finaly approach we will take, but for now we just
// display 64 bit parts for large input/outputs and 'fitted' size for smaller
// input/outputs. We will probably want to make this more flexible in the
// future, and allow for a range of different part sizes. But for now we are
// simple and fix the part size.
func getPartSize(totalSize int) int {
	mustValidInputOutputSize(totalSize)
	switch totalSize {
	case 512, 256, 128:
		return 64
	default:
		return totalSize
	}
}

package register

import (
	"fmt"

	"github.com/fmstephe/simd_explorer/pkg/assembly"
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
		Base2:  NewUIRegister(app, inst.InputSize(), inst.OutputSize(), 2, cBroadcaster),
		Base10: NewUIRegister(app, inst.InputSize(), inst.OutputSize(), 10, cBroadcaster),
		Base16: NewUIRegister(app, inst.InputSize(), inst.OutputSize(), 16, cBroadcaster),
	}

	// Set all parts to have 0 values
	cBroadcaster.broadcastZeros()

	return rs
}

type UIRegister struct {
	box tview.Primitive
}

func NewUIRegister(app *stackapp.StackApp, inputSize, outputSize, base int, cBroadcaster *changeBroadcaster) *UIRegister {
	mustValidInputOutputSize(inputSize)
	mustValidInputOutputSize(outputSize)

	inputPartSize := getPartSize(inputSize)
	outputPartSize := getPartSize(outputSize)

	input := NewRegisterInputs(app, inputPartSize, inputSize, base, cBroadcaster)
	output := NewRegisterOutputs(app, outputPartSize, outputSize, base, cBroadcaster)

	// Add update receivers, now that all initialisation updates have completed
	cBroadcaster.addInputReceiver(input)
	cBroadcaster.addOutputReceiver(output)

	gridLeft := tview.NewGrid()
	gridLeft.SetBorder(true)
	gridLeft.SetTitle(fmt.Sprintf("Inputs Base %d", base))

	gridRight := tview.NewGrid()
	gridRight.SetBorder(true)
	gridRight.SetTitle(fmt.Sprintf("Outputs Base %d", base))

	gridLeft.AddItem(input.GetBox(), 0, 0, 1, 1, 0, 0, true)
	gridRight.AddItem(output.GetBox(), 0, 0, 1, 1, 0, 0, false)

	grid := tview.NewGrid()
	grid.AddItem(gridLeft, 0, 0, 1, 1, 0, 0, true)
	grid.AddItem(gridRight, 0, 1, 1, 1, 0, 0, false)

	return &UIRegister{
		box: grid,
	}
}

func (r *UIRegister) GetPrimitive() tview.Primitive {
	return r.box
}

// TODO this likely isn't the finaly approach we will take, but for now we just
// display 64 bit parts for large input/outputs and 'fitted' size for smaller
// input/outputs. We will probably want to make this more flexible in the
// future, and allow for a range of different part sizes. But for now we are
// simple and fix the part size.
func getPartSize(totalSize int) int {
	mustValidInputOutputSize(totalSize)
	switch totalSize {
	case 512, 256, 128, 64:
		return 64
	default:
		return totalSize
	}
}

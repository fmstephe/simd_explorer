package register

import (
	"fmt"

	"github.com/rivo/tview"
)

type UIRegisterSet struct {
	Base2  *UIRegister
	Base10 *UIRegister
	Base16 *UIRegister
}

func NewUIRegisterSet(app *tview.Application, simdsize int) *UIRegisterSet {
	cBroadcaster := newChangeBroadcaster(simdsize)

	rs := &UIRegisterSet{
		Base2:  NewUIRegister(app, simdsize, 2, cBroadcaster),
		Base10: NewUIRegister(app, simdsize, 10, cBroadcaster),
		Base16: NewUIRegister(app, simdsize, 16, cBroadcaster),
	}

	// Set all parts to have 0 values
	cBroadcaster.broadcastZeros()

	return rs
}

type UIRegister struct {
	simdsize int
	base     int
	box      tview.Primitive
}

func NewUIRegister(app *tview.Application, simdsize, base int, cBroadcaster *changeBroadcaster) *UIRegister {
	mustValidSimdsize(simdsize)

	input64 := NewRegisterInputs(app, 64, simdsize, base, cBroadcaster)
	output64 := NewRegisterOutputs(app, 64, simdsize, base, cBroadcaster)

	input32 := NewRegisterInputs(app, 32, simdsize, base, cBroadcaster)
	output32 := NewRegisterOutputs(app, 32, simdsize, base, cBroadcaster)

	input16 := NewRegisterInputs(app, 16, simdsize, base, cBroadcaster)
	output16 := NewRegisterOutputs(app, 16, simdsize, base, cBroadcaster)

	input8 := NewRegisterInputs(app, 8, simdsize, base, cBroadcaster)
	output8 := NewRegisterOutputs(app, 8, simdsize, base, cBroadcaster)

	// Add update receivers, now that all initialisation updates have completed
	cBroadcaster.addReceiver(input64)
	cBroadcaster.addReceiver(output64)

	cBroadcaster.addReceiver(input32)
	cBroadcaster.addReceiver(output32)

	cBroadcaster.addReceiver(input16)
	cBroadcaster.addReceiver(output16)

	cBroadcaster.addReceiver(input8)
	cBroadcaster.addReceiver(output8)

	gridLeft := tview.NewGrid()
	gridLeft.SetBorder(true)
	gridLeft.SetTitle(fmt.Sprintf("Inputs Base %d", base))

	gridRight := tview.NewGrid()
	gridRight.SetBorder(true)
	gridRight.SetTitle(fmt.Sprintf("Outputs Base %d", base))

	gridLeft.AddItem(input64.GetBox(), 0, 0, 1, 1, 0, 0, true)
	gridRight.AddItem(output64.GetBox(), 0, 0, 1, 1, 0, 0, false)

	gridLeft.AddItem(input32.GetBox(), 1, 0, 1, 1, 0, 0, false)
	gridRight.AddItem(output32.GetBox(), 1, 0, 1, 1, 0, 0, false)

	gridLeft.AddItem(input16.GetBox(), 2, 0, 1, 1, 0, 0, false)
	gridRight.AddItem(output16.GetBox(), 2, 0, 1, 1, 0, 0, false)

	gridLeft.AddItem(input8.GetBox(), 3, 0, 1, 1, 0, 0, false)
	gridRight.AddItem(output8.GetBox(), 3, 0, 1, 1, 0, 0, false)

	grid := tview.NewGrid()
	grid.AddItem(gridLeft, 0, 0, 1, 1, 0, 0, true)
	grid.AddItem(gridRight, 0, 1, 1, 1, 0, 0, false)

	return &UIRegister{
		simdsize: simdsize,
		base:     base,
		box:      grid,
	}
}

func (r *UIRegister) GetPrimitive() tview.Primitive {
	return r.box
}

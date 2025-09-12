package ui

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/rivo/tview"
)

type RegisterOutputs struct {
	id uuid.UUID

	app        *tview.Application
	allOutputs []*tview.TextView
	box        *tview.Flex

	bitsize  int
	simdsize int
	base     int

	converter *valueConverter
}

func NewRegisterOutputs(app *tview.Application, bitsize, simdsize, base int) *RegisterOutputs {
	textWidth := textWidthForBitsize(bitsize)
	outputsCount := partsForBitsize(bitsize, simdsize)

	rOutputs := &RegisterOutputs{
		id: uuid.New(),

		app:        app,
		allOutputs: make([]*tview.TextView, outputsCount),
		box:        tview.NewFlex(),

		bitsize:  bitsize,
		simdsize: simdsize,
		base:     base,

		converter: newValueConverter(bitsize, 16),
	}

	for i := range rOutputs.allOutputs {
		output := tview.NewTextView()
		output.SetText("0")
		output.SetSize(1, textWidth)
		output.SetBorderPadding(0, 0, 0, 0)

		rOutputs.allOutputs[i] = output
		rOutputs.box.AddItem(output, 0, 1, false)
	}

	// Initialise outputs to all zeros
	rOutputs.dstDataChanged(make([]byte, simdsize/8))

	return rOutputs
}

func (out *RegisterOutputs) receiverId() uuid.UUID {
	return out.id
}

func (out *RegisterOutputs) GetBox() *tview.Flex {
	return out.box
}

func (out *RegisterOutputs) dstDataChanged(bytes []byte) {
	if (len(bytes) * 8) != out.simdsize {
		panic(fmt.Errorf("Bad data update, received %d bits, but need %d", len(bytes)*8, out.simdsize))
	}
	fmt.Printf("%s received data change %0.8b\n\n", out.describe(), bytes)

	bytesPer := out.bitsize / 8

	for i, output := range out.allOutputs {
		idx := i * bytesPer
		txt := out.converter.bytesToString(bytes[idx:])

		if output.GetText(false) != txt {
			fmt.Printf("%s[%d] changing text from %q to %q\n", out.describe(), i, output.GetText(false), txt)
			output.SetText(txt)
		}
	}
}

func (out *RegisterOutputs) describe() string {
	return fmt.Sprintf("%q-%d-%d-output", out.id.String()[:6], out.bitsize, out.simdsize)
}

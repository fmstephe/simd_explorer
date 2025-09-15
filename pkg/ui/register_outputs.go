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

	return rOutputs
}

func (out *RegisterOutputs) receiverId() uuid.UUID {
	return out.id
}

func (out *RegisterOutputs) getBox() *tview.Flex {
	return out.box
}

func (out *RegisterOutputs) dataParts() int {
	return len(out.allOutputs)
}

func (out *RegisterOutputs) setPart(i int, chunk []byte) {
	txt := out.converter.bytesToString(chunk)
	output := out.allOutputs[i]
	if output.GetText(false) != txt {
		// Only set the output text if the new value is
		// different from the old, this reduces noise in the logs
		fmt.Printf("%s[%d] changing text from %q to %q using %0.8b\n", out.describe(), i, output.GetText(false), txt, chunk)
		out.allOutputs[i].SetText(txt)
	}
}

func (out *RegisterOutputs) describe() string {
	return fmt.Sprintf("%q-%d-%d-output", out.id.String()[:6], out.bitsize, out.simdsize)
}

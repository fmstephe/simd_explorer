package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
	"github.com/rivo/tview"
)

type RegisterInputs struct {
	id uuid.UUID

	app       *tview.Application
	allInputs []*tview.InputField
	box       *tview.Flex

	focus int

	bitsize  int
	simdsize int
	base     int

	// Callback when the data in this set of inputs is changed
	cBroadcaster *changeBroadcaster
	converter    *valueConverter
}

func NewRegisterInputs(app *tview.Application, bitsize, simdsize, base int, cBroadcaster *changeBroadcaster) *RegisterInputs {
	textWidth := textWidthForBitsize(bitsize)
	inputsCount := partsForBitsize(bitsize, simdsize)

	rInputs := &RegisterInputs{
		id: uuid.New(),

		app:       app,
		allInputs: make([]*tview.InputField, inputsCount),
		box:       tview.NewFlex(),

		focus: 0,

		bitsize:  bitsize,
		simdsize: simdsize,
		base:     base,

		cBroadcaster: cBroadcaster,
		converter:    newValueConverter(bitsize, base),
	}

	for i := range inputsCount {
		input := tview.NewInputField()
		input.SetFieldWidth(textWidth)
		input.SetBorderPadding(0, 0, 0, 0)
		input.SetAcceptanceFunc(rInputs.converter.inputAcceptor())

		rInputs.allInputs[i] = input
		rInputs.box.AddItem(input, 0, 1, true)
	}

	rInputs.initFocusCycling()

	// We delay setting the changed-func to reduce initialisation noise in
	// the logs
	for _, input := range rInputs.allInputs {
		input.SetChangedFunc(func(txt string) {
			// Broadcast the change to all data-changed
			// receivers
			rInputs.srcDataChanged()
		})
	}

	return rInputs
}

func (in *RegisterInputs) receiverId() uuid.UUID {
	return in.id
}

func (in *RegisterInputs) getBox() *tview.Flex {
	return in.box
}

func (in *RegisterInputs) initFocusCycling() {
	in.box.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			in.cycleFocus(1)
		case tcell.KeyBacktab:
			in.cycleFocus(-1)
		}

		return event
	})
}

func (in *RegisterInputs) cycleFocus(move int) {
	in.focus += move
	idx := in.focus % len(in.allInputs)
	if idx < 0 {
		idx = len(in.allInputs) + idx
	}
	in.app.SetFocus(in.allInputs[idx])
}

func (in *RegisterInputs) srcDataChanged() {
	bytes := make([]byte, 0, in.simdsize)

	for _, input := range in.allInputs {
		txt := input.GetText()
		bytes = append(bytes, in.converter.stringToBytes(txt)...)
	}

	fmt.Printf("%s broadcasting data change %0.8b\n\n", in.describe(), bytes)

	in.cBroadcaster.broadcastChange(bytes)
}

func (in *RegisterInputs) dataParts() int {
	return len(in.allInputs)
}

func (in *RegisterInputs) setPart(i int, chunk []byte) {
	txt := in.converter.bytesToString(chunk)
	input := in.allInputs[i]
	if input.GetText() != txt {
		// Only set the output text if the new value is
		// different from the old, this reduces noise in the logs
		fmt.Printf("%s[%d] changing text from %q to %q using %0.8b\n", in.describe(), i, input.GetText(), txt, chunk)
		in.allInputs[i].SetText(txt)
	}
}

func (in *RegisterInputs) describe() string {
	return fmt.Sprintf("%q-%d-%d--input", in.id.String()[:6], in.bitsize, in.simdsize)
}

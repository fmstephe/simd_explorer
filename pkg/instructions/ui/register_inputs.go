package ui

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
	"github.com/rivo/tview"
)

type RegisterInputs struct {
	id uuid.UUID

	app *tview.Application

	bitsize     int
	simdsize    int
	inputsCount int

	focus     int
	allInputs []*tview.InputField

	box *tview.Flex

	// Callback when the data in this set of inputs is changed
	cBroadcaster *changeBroadcaster

	converter *valueConverter
}

func NewRegisterInputs(app *tview.Application, bitsize, simdsize int, cBroadcaster *changeBroadcaster) *RegisterInputs {
	textWidth := textWidthForBitsize(bitsize)
	inputsCount := partsForBitsize(bitsize, simdsize)

	rInputs := &RegisterInputs{
		id:           uuid.New(),
		app:          app,
		bitsize:      bitsize,
		simdsize:     simdsize,
		inputsCount:  inputsCount,
		focus:        0,
		allInputs:    make([]*tview.InputField, inputsCount),
		box:          tview.NewFlex(),
		cBroadcaster: cBroadcaster,
		converter:    newValueConverter(bitsize, 16),
	}

	for i := range inputsCount {
		input := tview.NewInputField()
		input.SetFieldWidth(textWidth)
		input.SetBorderPadding(0, 0, 0, 0)
		input.SetAcceptanceFunc(rInputs.converter.inputAcceptor())

		rInputs.allInputs[i] = input
		rInputs.box.AddItem(input, 0, 1, true)
	}

	rInputs.InitFocusCycling()

	// Initialise inputs to zeros
	rInputs.dstDataChanged(make([]byte, simdsize/8))

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

func (in *RegisterInputs) GetBox() *tview.Flex {
	return in.box
}

func (in *RegisterInputs) InitFocusCycling() {
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
	idx := in.focus % in.inputsCount
	if idx < 0 {
		idx = in.inputsCount + idx
	}
	in.app.SetFocus(in.allInputs[idx])
}

func (in *RegisterInputs) srcDataChanged() {
	bytes := make([]byte, 0, in.simdsize)

	for _, input := range in.allInputs {
		txt := input.GetText()
		bytes = append(bytes, in.converter.stringToBytes(txt)...)
	}

	fmt.Printf("%s broadcasting data change %q\n\n", in.describe(), bytes)

	in.cBroadcaster.broadcastChange(bytes, in.id)
}

func (in *RegisterInputs) dstDataChanged(bytes []byte) {
	if (len(bytes) * 8) != in.simdsize {
		panic(fmt.Errorf("Bad data update, received %d bits, but need %d", len(bytes)*8, in.simdsize))
	}
	fmt.Printf("%s received data change %q\n\n", in.describe(), bytes)

	bytesPer := in.bitsize / 8

	for i, input := range in.allInputs {
		idx := i * bytesPer
		txt := in.converter.bytesToString(bytes[idx:])
		if input.GetText() != txt {
			// Only set the input text if the new value is
			// different from the old The prevents the
			// data-changechange->update-event->data-change loop
			// from running indefinitely.
			fmt.Printf("%s[%d] changing text from %q to %q\n", in.describe(), i, input.GetText(), txt)
			input.SetText(txt)
		}
	}
}

// InputFieldInteger accepts unsigned integers.
func InputFieldUint(bitsize int) func(string, rune) bool {
	return func(txt string, _ rune) bool {
		_, err := strconv.ParseUint(txt, 10, bitsize)
		fmt.Printf("Accepting(%b) %s", err == nil, txt)
		return err == nil
	}
}

func (in *RegisterInputs) describe() string {
	return fmt.Sprintf("%q-%d-%d--input", in.id.String()[:6], in.bitsize, in.simdsize)
}

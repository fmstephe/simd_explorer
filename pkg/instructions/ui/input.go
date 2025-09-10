package ui

import (
	"encoding/binary"
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type RegisterInputs struct {
	app *tview.Application

	bitsize     int
	simdsize    int
	inputsCount int

	focus     int
	allInputs []*tview.InputField

	box *tview.Flex

	// Callback when the data in this set of inputs is changed
	dataChanged func([]byte)
}

func NewRegisterInputs(app *tview.Application, bitsize, simdsize int, dataChanged func([]byte)) *RegisterInputs {
	textWidth := textWidthForBitsize(bitsize)
	inputsCount := inputsForBitsize(bitsize, simdsize)

	rInputs := &RegisterInputs{
		app:         app,
		bitsize:     bitsize,
		simdsize:    simdsize,
		inputsCount: inputsCount,
		focus:       0,
		allInputs:   make([]*tview.InputField, inputsCount),
		box:         tview.NewFlex(),
		dataChanged: dataChanged,
	}

	for i := range inputsCount {
		input := tview.NewInputField()
		input.SetFieldWidth(textWidth)
		input.SetBorderPadding(0, 0, 0, 0)
		input.SetAcceptanceFunc(tview.InputFieldInteger)
		input.SetChangedFunc(func(text string) {
			// When _any_ piece of data changes we reprocess the data from all inputs
			rInputs.sourceDataChanged()
		})

		rInputs.allInputs[i] = input
		rInputs.box.AddItem(input, 0, 1, true)
	}

	rInputs.InitCapture()

	return rInputs
}

func (in *RegisterInputs) GetBox() *tview.Flex {
	return in.box
}

func (in *RegisterInputs) InitCapture() {
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

func (in *RegisterInputs) sourceDataChanged() {
	endian := binary.LittleEndian
	bytes := make([]byte, 0, in.simdsize)

	for _, input := range in.allInputs {
		txt := input.GetText()
		if txt == "" {
			// TODO is this the right way to deal with this?
			// shouldn't all of the values be set somehow?
			txt = "0"
		}
		val, err := strconv.ParseUint(txt, 10, in.bitsize)
		if err != nil {
			panic(fmt.Errorf("Unexpected value %q found in register input, expecting decimal with bitsize %d", input.GetText(), in.bitsize))
		}
		switch in.bitsize {
		case 8:
			bytes = append(bytes, byte(val))
		case 16:
			bytes = endian.AppendUint16(bytes, uint16(val))
		case 32:
			bytes = endian.AppendUint32(bytes, uint32(val))
		case 64:
			bytes = endian.AppendUint64(bytes, val)
		}
	}

	in.dataChanged(bytes)
}

func (in *RegisterInputs) destinationDataChange(bytes []byte) {
	endian := binary.LittleEndian

	if (len(bytes) * 8) != in.simdsize {
		panic(fmt.Errorf("Bad data update, received %d bits, but need %d", len(bytes)*8, in.simdsize))
	}

	bytesPer := in.bitsize / 8

	for i, input := range in.allInputs {
		idx := i * bytesPer
		val := uint64(0)
		switch in.bitsize {
		case 8:
			val = uint64(bytes[idx])
		case 16:
			val = uint64(endian.Uint16(bytes[idx:]))
		case 32:
			val = uint64(endian.Uint32(bytes[idx:]))
		case 64:
			val = endian.Uint64(bytes[idx:])
		}

		input.SetText(strconv.FormatUint(val, 10))
	}
}

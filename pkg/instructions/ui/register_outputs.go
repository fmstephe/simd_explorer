package ui

import (
	"encoding/binary"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/rivo/tview"
)

type RegisterOutputs struct {
	id  uuid.UUID
	app *tview.Application

	bitsize      int
	simdsize     int
	outputsCount int

	outputs []*tview.TextView

	box *tview.Flex
}

func NewRegisterOutputs(app *tview.Application, bitsize, simdsize int) *RegisterOutputs {
	textWidth := textWidthForBitsize(bitsize)
	inputsCount := inputsForBitsize(bitsize, simdsize)

	rOutputs := &RegisterOutputs{
		id:           uuid.New(),
		app:          app,
		bitsize:      bitsize,
		simdsize:     simdsize,
		outputsCount: inputsCount,
		outputs:      make([]*tview.TextView, inputsCount),
		box:          tview.NewFlex(),
	}

	for i := range inputsCount {
		output := tview.NewTextView()
		output.SetSize(1, textWidth)
		output.SetBorderPadding(0, 0, 0, 0)

		rOutputs.outputs[i] = output
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
	endian := binary.LittleEndian

	if (len(bytes) * 8) != out.simdsize {
		panic(fmt.Errorf("Bad data update, received %d bits, but need %d", len(bytes)*8, out.simdsize))
	}

	bytesPer := out.bitsize / 8

	for i, output := range out.outputs {
		idx := i * bytesPer
		val := uint64(0)
		switch out.bitsize {
		case 8:
			val = uint64(bytes[idx])
		case 16:
			val = uint64(endian.Uint16(bytes[idx:]))
		case 32:
			val = uint64(endian.Uint32(bytes[idx:]))
		case 64:
			val = endian.Uint64(bytes[idx:])
		}

		output.SetText(strconv.FormatUint(val, 10))
	}
}

package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
	"github.com/rivo/tview"
)

type RegisterParts struct {
	id uuid.UUID
	// Indicates if this is input or output for logging
	kind string

	app      *tview.Application
	allParts []uiPart
	box      *tview.Grid

	focus int

	bitsize  int
	simdsize int
	base     int

	// Callback when the data in this set of parts is changed
	cBroadcaster *changeBroadcaster
	converter    *valueConverter
}

func NewRegisterInputs(app *tview.Application, bitsize, simdsize, base int, cBroadcaster *changeBroadcaster) *RegisterParts {
	return NewRegisterParts(app, bitsize, simdsize, base, &inputPartBuilder{}, cBroadcaster)
}

func NewRegisterOutputs(app *tview.Application, bitsize, simdsize, base int, cBroadcaster *changeBroadcaster) *RegisterParts {
	return NewRegisterParts(app, bitsize, simdsize, base, &textViewPartBuilder{}, cBroadcaster)
}

func NewRegisterParts(app *tview.Application, bitsize, simdsize, base int, partsBuilder uiPartBuilder, cBroadcaster *changeBroadcaster) *RegisterParts {
	textWidth := textWidthForBitsize(bitsize)
	partsCount := partsForBitsize(bitsize, simdsize)

	grid := tview.NewGrid()
	// We always have a maximum of 8 columns per row
	grid.SetRows(3, 3, 3, 3, 3, 3, 3, 3)
	grid.SetBorder(false)

	rParts := &RegisterParts{
		id:   uuid.New(),
		kind: partsBuilder.kind(),

		app:      app,
		allParts: make([]uiPart, partsCount),
		box:      grid,

		focus: 0,

		bitsize:  bitsize,
		simdsize: simdsize,
		base:     base,

		cBroadcaster: cBroadcaster,
		converter:    newValueConverter(bitsize, base),
	}

	for i := range partsCount {
		part := partsBuilder.build()
		part.setTitle(fmt.Sprintf("%d:%d", i*bitsize, (i+1)*bitsize))
		part.setBorder(true)
		part.setFieldWidth(textWidth)
		part.setAcceptanceFunc(rParts.converter.inputAcceptor())

		rParts.allParts[i] = part

		column := i % 8
		row := i / 8
		grid.AddItem(part.primitive(), row, column, 1, 1, 1, 1, true)
	}

	rParts.initFocusCycling()

	// We delay setting the changed-func to reduce initialisation noise in
	// the logs
	for _, part := range rParts.allParts {
		part.setChangedFunc(func(txt string) {
			// Broadcast the change to all data-changed
			// receivers
			rParts.srcDataChanged()
		})
	}

	return rParts
}

func (in *RegisterParts) receiverId() uuid.UUID {
	return in.id
}

func (in *RegisterParts) getBox() *tview.Grid {
	return in.box
}

func (in *RegisterParts) initFocusCycling() {
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

func (in *RegisterParts) cycleFocus(move int) {
	in.focus += move
	idx := in.focus % len(in.allParts)
	if idx < 0 {
		idx = len(in.allParts) + idx
	}
	in.app.SetFocus(in.allParts[idx].primitive())
}

func (in *RegisterParts) srcDataChanged() {
	bytes := make([]byte, 0, in.simdsize)

	for _, part := range in.allParts {
		txt := part.getText()
		bytes = append(bytes, in.converter.stringToBytes(txt)...)
	}

	fmt.Printf("%s broadcasting data change %0.8b\n\n", in.describe(), bytes)

	in.cBroadcaster.broadcastChange(bytes)
}

func (in *RegisterParts) dataParts() int {
	return len(in.allParts)
}

func (in *RegisterParts) setPart(i int, chunk []byte) {
	txt := in.converter.bytesToString(chunk)
	part := in.allParts[i]
	if part.getText() != txt {
		// Only set the output text if the new value is
		// different from the old, this reduces noise in the logs
		fmt.Printf("%s[%d] changing text from %q to %q using %0.8b\n", in.describe(), i, part.getText(), txt, chunk)
		in.allParts[i].setText(txt)
	}
}

func (in *RegisterParts) describe() string {
	return fmt.Sprintf("%q-%d-%d--%s", in.id.String()[:6], in.bitsize, in.simdsize, in.kind)
}

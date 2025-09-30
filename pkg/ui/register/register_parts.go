package register

import (
	"fmt"

	"github.com/fmstephe/simd_explorer/pkg/ui/number"
	"github.com/fmstephe/simd_explorer/pkg/ui/stackapp"
	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
	"github.com/rivo/tview"
)

type RegisterParts struct {
	id uuid.UUID
	// Indicates if this is input or output for logging
	kind string

	app      *stackapp.StackApp
	allParts []uiPart
	box      *tview.Grid

	focus int

	// Callback when the data in this set of parts is changed
	uiRegister *UIRegister
	parameter  *number.Parameter
}

func NewRegisterInputs(app *stackapp.StackApp, parameter *number.Parameter, uiRegister *UIRegister) *RegisterParts {
	return NewRegisterParts(app, parameter, &inputPartBuilder{}, uiRegister)
}

func NewRegisterOutputs(app *stackapp.StackApp, parameter *number.Parameter, uiRegister *UIRegister) *RegisterParts {
	return NewRegisterParts(app, parameter, &textViewPartBuilder{}, uiRegister)
}

func NewRegisterParts(app *stackapp.StackApp, parameter *number.Parameter, partsBuilder uiPartBuilder, uiRegister *UIRegister) *RegisterParts {
	grid := tview.NewGrid()
	// We always have a maximum of 8 columns per row
	grid.SetRows(3, 3, 3, 3, 3, 3, 3, 3)
	grid.SetBorder(true)
	grid.SetTitle(fmt.Sprintf("%d Bit %s", parameter.PartBitWidth(), partsBuilder.kind()))

	rParts := &RegisterParts{
		id:   uuid.New(),
		kind: partsBuilder.kind(),

		app:      app,
		allParts: make([]uiPart, parameter.Parts()),
		box:      grid,

		focus: 0,

		uiRegister: uiRegister,
		parameter:  parameter,
	}

	parts := parameter.Parts()
	partBitWidth := parameter.PartBitWidth()
	partsPerLine := calcPartsPerLine(parameter)
	partConverter := parameter.Converter()

	// TODO we are going to have to build more consideration into this for 512 bit registers
	// on my monitor right now I can't display the 64 bit parts of the 512 bit register on a single line.
	for i := range parts {
		part := partsBuilder.build()
		part.setTitle(fmt.Sprintf("%d:%d", i*partBitWidth, (i+1)*partBitWidth))
		part.setBorder(true)
		part.setFieldWidth(partConverter.GetTextWidth())
		part.setAcceptanceFunc(partConverter.InputAcceptor())

		rParts.allParts[i] = part

		column := i % partsPerLine
		row := i / partsPerLine
		grid.AddItem(part.primitive(), row, column, 1, 1, 1, 1, true)
	}

	rParts.initFocusCycling()

	// We delay setting the changed-func to reduce initialisation noise in
	// the logs
	for _, part := range rParts.allParts {
		part.setChangedFunc(func(txt string) {
			// Notify the uiRegister that some input data has changed
			uiRegister.inputsChanged()
		})
	}

	return rParts
}

func (in *RegisterParts) receiverId() uuid.UUID {
	return in.id
}

func (in *RegisterParts) GetBox() *tview.Grid {
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

func (in *RegisterParts) getData() []byte {
	bytes := make([]byte, 0, in.parameter.TotalBitWidth())

	for _, part := range in.allParts {
		txt := part.getText()
		bytes = append(bytes, in.parameter.Converter().StringToBytes(txt)...)
	}

	fmt.Printf("%s broadcasting data change %0.8b\n\n", in.describe(), bytes)

	return bytes
}

func (in *RegisterParts) setData(bytes []byte) {
	if (len(bytes) * 8) != in.parameter.TotalBitWidth() {
		panic(fmt.Errorf("Bad data update, received %d bits, but need %d", len(bytes)*8, in.parameter.TotalBitWidth()))
	}
	partsCount := len(in.allParts)
	fmt.Printf("%s received data change %0.8b\n\n", in.describe(), bytes)

	if len(bytes)%partsCount != 0 {
		panic(fmt.Errorf("%s update with %d bytes, not cleanly divisible by %d parts", in.describe(), len(bytes), partsCount))
	}

	bytesPer := len(bytes) / partsCount
	for i, part := range in.allParts {
		idx := i * bytesPer
		chunk := bytes[idx : idx+bytesPer]
		txt := in.parameter.Converter().BytesToString(chunk)
		if part.getText() != txt {
			// Only set the output text if the new value is
			// different from the old, this reduces noise in the logs
			fmt.Printf("%s[%d] changing text from %q to %q using %0.8b\n", in.describe(), i, part.getText(), txt, chunk)
			part.setText(txt)
		}
	}
}

func (in *RegisterParts) describe() string {
	return fmt.Sprintf("%q-%d-%d--%s", in.id.String()[:6], in.parameter.PartBitWidth(), in.parameter.TotalBitWidth(), in.kind)
}

func calcPartsPerLine(parameter *number.Parameter) int {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}
	defer screen.Fini()
	width, _ := screen.Size()
	// We subtract 4 here to heuristically account for the input/output borders
	half := (width - 4) / 2
	// NB: The +2 heuristally allows for a border
	partWidth := parameter.Converter().GetTextWidth() + 2
	// For parts which are smaller than 10, we force the size to 10 to
	// allow the title to display correctly
	partWidth = max(partWidth, 10)
	perLine := half / partWidth
	return perLine
}

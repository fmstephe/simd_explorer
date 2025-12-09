package uiio

import (
	"fmt"
	"log"
	"strconv"

	"github.com/fmstephe/simd_explorer/pkg/ui/number"
	"github.com/fmstephe/simd_explorer/pkg/ui/stackapp"
	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
	"github.com/rivo/tview"
)

type UIParameterParts struct {
	id uuid.UUID
	// Indicates if this is input or output for logging
	kind string

	app      *stackapp.StackApp
	allParts []uiParameterPart
	box      *tview.Grid

	// Callback when the data in this set of parts is changed
	uiParameters *UIInstruction
	parameter    *number.Parameter
}

func NewUIParameterInputs(app *stackapp.StackApp, parameter *number.Parameter, uiRegister *UIInstruction) *UIParameterParts {
	return NewUIParameterParts(app, parameter, &inputPartBuilder{}, uiRegister)
}

func NewUIParameterOutputs(app *stackapp.StackApp, parameter *number.Parameter, uiRegister *UIInstruction) *UIParameterParts {
	return NewUIParameterParts(app, parameter, &textViewPartBuilder{}, uiRegister)
}

func NewUIParameterParts(app *stackapp.StackApp, parameter *number.Parameter, partsBuilder uiPartBuilder, uiParameters *UIInstruction) *UIParameterParts {
	grid := tview.NewGrid()
	// We always have a maximum of 8 columns per row
	grid.SetRows(3, 3, 3, 3, 3, 3, 3, 3)
	grid.SetBorder(true)
	grid.SetTitle(fmt.Sprintf(" '%s' %s (%d bits) ", parameter.Name(), parameter.Kind(), parameter.TotalBitWidth()))

	pParts := &UIParameterParts{
		id:   uuid.New(),
		kind: partsBuilder.kind(),

		app:      app,
		allParts: make([]uiParameterPart, parameter.Parts()),
		box:      grid,

		uiParameters: uiParameters,
		parameter:    parameter,
	}

	parts := parameter.Parts()
	partBitWidth := parameter.PartBitWidth()
	partsPerLine := calcPartsPerLine(parameter)

	for i := range parts {
		part := partsBuilder.build()
		part.setTitle(fmt.Sprintf("%d:%d", i*partBitWidth, (i+1)*partBitWidth))
		part.setBorder(true)
		part.setFieldWidth(parameter.GetTextWidth())
		part.setFieldWidth(parameter.Base())
		part.setAcceptanceFunc(parameter.InputAcceptor())

		pParts.allParts[i] = part

		column := i % partsPerLine
		row := i / partsPerLine
		grid.AddItem(part.primitive(), row, column, 1, 1, 1, 1, true)
	}

	// We delay setting the changed-func to reduce initialisation noise in
	// the logs
	for _, part := range pParts.allParts {
		part.setChangedFunc(func(txt string) {
			if normalised, changed := parameter.Normalised(txt); changed {
				// If normalisation changed the text value -
				// set it again with the normalised value.
				// Carefully avoid any of the other processing
				// here.
				part.setText(normalised)
				return
			}
			// If the part's txt is an unstable value, then warn the user by setting background to yellow.
			// This won't help colourblind users, but we can resolve that later if needed.
			if !parameter.IsStable(txt) {
				part.setBackgroundColor(tcell.ColorRed)
			} else {
				part.setBackgroundColor(tview.Styles.ContrastBackgroundColor)
			}
			// Notify the uiParameters that some input data has changed
			uiParameters.inputsChanged()
		})
	}

	return pParts
}

func (in *UIParameterParts) GetBox() *tview.Grid {
	return in.box
}

func (in *UIParameterParts) SetDefaults(start byte) (end byte) {
	end = start
	for _, part := range in.allParts {
		// This is a bit of a hack to set default values for all inputs
		// We restrict it to string representation of a single byte value, as this is our smallest representable value.
		// TODO let's see how this works in practice.
		part.setText(strconv.FormatInt(int64(end), in.parameter.Base()))
		end++
	}
	return end
}

func (in *UIParameterParts) syncToParameter() {
	txts := make([]string, len(in.allParts))

	for i, part := range in.allParts {
		txt := part.getText()
		txts[i] = txt
	}
	log.Printf("%s synced to parameter: %v", in.describe(), txts)

	in.parameter.DataFromStrings(txts)
}

func (in *UIParameterParts) syncFromParameter() {
	txts := in.parameter.DataToStrings()
	log.Printf("%s sync from parameter: %v", in.describe(), txts)

	for i, txt := range txts {
		part := in.allParts[i]
		if part.getText() != txt {
			// Only set the output text if the new value is
			// different from the old, this reduces noise in the logs
			log.Printf("%s[%d] changing text from %q to %q", in.describe(), i, part.getText(), txt)
			part.setText(txt)
		}
	}
}

func (in *UIParameterParts) selectablePrimitives() []tview.Primitive {
	selectables := []tview.Primitive{}
	for _, part := range in.allParts {
		selectables = append(selectables, part.primitive())
	}
	return selectables
}

func (in *UIParameterParts) describe() string {
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
	partWidth := parameter.GetTextWidth() + 2
	// For parts which are smaller than 10, we force the size to 10 to
	// allow the title to display correctly
	partWidth = max(partWidth, 10)
	perLine := half / partWidth
	return min(perLine, 8)
}

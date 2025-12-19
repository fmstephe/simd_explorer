package uiio

import (
	"slices"

	"github.com/fmstephe/simd_explorer/pkg/assembly"
	"github.com/fmstephe/simd_explorer/pkg/ui/stackapp"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UIInstruction struct {
	instruction       assembly.Instruction
	inputUIParameters []*UIParameterParts
	outputUIParameter *UIParameterParts
	box               *tview.Grid
	source            *tview.Box

	focus      int
	selectable []tview.Primitive
	app        *stackapp.StackApp
}

func NewUIInstruction(app *stackapp.StackApp, instruction assembly.Instruction) *UIInstruction {
	// We allocate the struct early to get the inputsChanged callback function
	uiInst := &UIInstruction{}

	inputs, selectable := buildInputs(app, instruction, uiInst.inputsChanged)
	output := NewUIParameterOutputs(app, instruction.Output())

	uiGrid := buildUIGrid(instruction, inputs, output)

	assemblyView := tview.NewTextView()
	assemblyView.SetBorder(true)
	assemblyView.SetTitle("Go Assembly")
	assemblyView.SetText(instruction.Assembly())

	// Fill in the struct fields
	*uiInst = UIInstruction{
		instruction:       instruction,
		inputUIParameters: inputs,
		outputUIParameter: output,
		box:               uiGrid,
		source:            assemblyView.Box,

		focus:      0,
		selectable: selectable,
		app:        app,
	}

	uiInst.setInputDefaults()
	uiInst.initInputCapture()

	return uiInst
}

func (r *UIInstruction) GetPrimitive() tview.Primitive {
	return r.box
}

func (r *UIInstruction) setInputDefaults() {
	vals := r.makeDefaultInputs()
	idx := 0
	for _, input := range r.inputUIParameters {
		input.SetValues(vals[idx : idx+input.Parts()])
		idx += input.Parts()
	}
}

func (r *UIInstruction) setInputDefaultsReverse() {
	vals := r.makeDefaultInputs()
	slices.Reverse(vals)
	idx := 0
	for _, input := range r.inputUIParameters {
		input.SetValues(vals[idx : idx+input.Parts()])
		idx += input.Parts()
	}
}

func (r *UIInstruction) setInputsZero() {
	for _, input := range r.inputUIParameters {
		input.SetValues(make([]int64, input.Parts()))
	}
}

func (r *UIInstruction) makeDefaultInputs() []int64 {
	totalVals := 0
	for _, input := range r.inputUIParameters {
		totalVals += input.Parts()
	}

	vals := make([]int64, totalVals)
	for i := range totalVals {
		vals[i] = int64(i)
	}

	return vals
}

func (r *UIInstruction) inputsChanged() {
	// Sync all ui input fields to the instruction Input() parameters
	for _, input := range r.inputUIParameters {
		input.syncToParameter()
	}

	// NB: Running the instruction syncs the Output() parameter
	r.instruction.Run()

	// Sync ui output fields from instruction Output() parameter
	r.outputUIParameter.syncFromParameter()
}

func (r *UIInstruction) initInputCapture() {
	// Establish keyboard shortcuts for the buttons
	r.box.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Handle non-character keys first
		switch event.Key() {
		case tcell.KeyTab:
			r.cycleFocus(1)
		case tcell.KeyBacktab:
			r.cycleFocus(-1)
			// TODO add arrow keys to this
		}

		// Handle character keys
		switch event.Rune() {
		case rune('z'), rune('Z'):
			r.setInputsZero()
		case rune('u'), rune('U'):
			r.setInputDefaults()
		case rune('r'), rune('R'):
			r.setInputDefaultsReverse()
		case rune('s'), rune('S'):
			// TODO display assembly source code here
		}

		// Allow the event to propagate
		return event
	})

}

func (r *UIInstruction) cycleFocus(move int) {
	r.focus += move
	idx := r.focus % len(r.selectable)
	if idx < 0 {
		idx = len(r.selectable) + idx
	}
	r.app.SetFocus(r.selectable[idx])
}

func buildInputs(app *stackapp.StackApp, instruction assembly.Instruction, inputsChanged func()) (inputs []*UIParameterParts, selectable []tview.Primitive) {
	for _, param := range instruction.Inputs() {
		input := NewUIParameterInputs(app, param, inputsChanged)
		inputs = append(inputs, input)
		selectable = append(selectable, input.selectablePrimitives()...)
	}

	return inputs, selectable
}

func buildUIGrid(instruction assembly.Instruction, inputs []*UIParameterParts, output *UIParameterParts) *tview.Grid {
	gridInputs := tview.NewGrid()
	for i, input := range inputs {
		gridInputs.AddItem(input.GetBox(), i, 0, 1, 1, 0, 0, true)
	}

	gridOutput := tview.NewGrid()
	gridOutput.AddItem(output.GetBox(), 0, 0, 1, 1, 0, 0, false)

	gridInputOutput := tview.NewGrid()
	gridInputOutput.AddItem(gridInputs, 0, 0, 1, 1, 0, 0, true)
	gridInputOutput.AddItem(gridOutput, 0, 1, 1, 1, 0, 0, false)

	gridButtons := buildDefaultButtons()

	gridOuter := tview.NewGrid()
	gridOuter.SetBorder(true)
	gridOuter.SetTitle(instruction.Name())
	gridOuter.SetRows(8, 0)

	gridOuter.AddItem(gridInputOutput, 1, 0, 1, 1, 0, 0, true)
	gridOuter.AddItem(gridButtons, 0, 0, 1, 1, 0, 0, false)

	return gridOuter
}

// NB: These buttons don't actually do anything right now - they just display the keyboard shortcuts for the default value setting functions. They really should either become real buttons (fix the mouse interaction problem) or we should display these shortcuts some other way.
func buildDefaultButtons() *tview.Grid {
	buttonClear := tview.NewButton(tview.Escape(`[Z]ero`))
	buttonFill := tview.NewButton(tview.Escape(`A[u]tofill`))
	buttonFillRev := tview.NewButton(tview.Escape(`Autofill [R]everse`))
	buttonSource := tview.NewButton(tview.Escape(`[S]how Source`))

	gridButtons := tview.NewGrid()
	gridButtons.AddItem(buttonClear, 0, 0, 1, 1, 0, 0, false)
	gridButtons.AddItem(buttonFill, 0, 1, 1, 1, 0, 0, false)
	gridButtons.AddItem(buttonFillRev, 0, 2, 1, 1, 0, 0, false)
	gridButtons.AddItem(buttonSource, 0, 3, 1, 1, 0, 0, false)

	return gridButtons
}

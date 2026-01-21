package uiio

import (
	"slices"

	"github.com/fmstephe/simd_explorer/pkg/assembly"
	"github.com/fmstephe/simd_explorer/pkg/ui/stackapp"
	"github.com/gdamore/tcell/v2"
	"github.com/fmstephe/tview"
)

type UIInstruction struct {
	instruction       assembly.Instruction
	inputUIParameters []*UIParameterParts
	outputUIParameter *UIParameterParts
	box               *tview.Grid
	source            *tview.Grid

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
	sourceGrid := buildSourceGrid(instruction)

	// Fill in the struct fields
	*uiInst = UIInstruction{
		instruction:       instruction,
		inputUIParameters: inputs,
		outputUIParameter: output,
		box:               uiGrid,
		source:            sourceGrid,

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
		// case tcell.KeyF1:
		// KeyF1 behaviour is currently managed in stackapp package
		case tcell.KeyF2:
			r.app.Push(r.source)
		case tcell.KeyF3:
			r.setInputsZero()
		case tcell.KeyF4:
			r.setInputDefaults()
		case tcell.KeyF5:
			r.setInputDefaultsReverse()
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
	gridButtons := buildInstructionButtons()

	gridInputs := tview.NewGrid()
	for i, input := range inputs {
		gridInputs.AddItem(input.GetBox(), i, 0, 1, 1, 0, 0, true)
	}

	gridOutput := tview.NewGrid()
	gridOutput.AddItem(output.GetBox(), 0, 0, 1, 1, 0, 0, false)

	gridInputOutput := tview.NewGrid()
	gridInputOutput.AddItem(gridInputs, 0, 0, 1, 1, 0, 0, true)
	gridInputOutput.AddItem(gridOutput, 0, 1, 1, 1, 0, 0, false)

	gridOuter := tview.NewGrid()
	gridOuter.SetBorder(true)
	gridOuter.SetTitle(instruction.Name())
	gridOuter.SetRows(3, 0)

	gridOuter.AddItem(gridButtons, 0, 0, 1, 1, 0, 0, false)
	gridOuter.AddItem(gridInputOutput, 1, 0, 1, 1, 0, 0, true)

	return gridOuter
}

func buildSourceGrid(instruction assembly.Instruction) *tview.Grid {
	gridButtons := buildSourceButtons()

	assemblyView := tview.NewTextView()
	assemblyView.SetBorder(true)
	assemblyView.SetTitle("Go Assembly")
	assemblyView.SetText(instruction.Assembly())

	gridOuter := tview.NewGrid()
	gridOuter.SetBorder(true)
	gridOuter.SetTitle(instruction.Name())
	gridOuter.SetRows(3, 0)

	gridOuter.AddItem(gridButtons, 0, 0, 1, 1, 0, 0, false)
	gridOuter.AddItem(assemblyView, 1, 0, 1, 1, 0, 0, true)

	return gridOuter
}

// NB: These buttons don't actually do anything right now - they just display the keyboard shortcuts for the default value setting functions. They really should either become real buttons (fix the mouse interaction problem) or we should display these shortcuts some other way.
func buildInstructionButtons() *tview.Grid {
	return buildButtonPanel(false)
}

// NB: These buttons don't actually do anything right now - they just display the keyboard shortcuts for the default value setting functions. They really should either become real buttons (fix the mouse interaction problem) or we should display these shortcuts some other way.
func buildSourceButtons() *tview.Grid {
	return buildButtonPanel(true)
}

// NB: These buttons don't actually do anything right now - they just display the keyboard shortcuts for the default value setting functions. They really should either become real buttons (fix the mouse interaction problem) or we should display these shortcuts some other way.
func buildButtonPanel(onlyBack bool) *tview.Grid {
	buttonBack := tview.NewButton(tview.Escape(`[F1] Back`))
	buttonSource := tview.NewButton(tview.Escape(`[F2] Show Source`)).SetDisabled(onlyBack)
	buttonZero := tview.NewButton(tview.Escape(`[F3] Zero`)).SetDisabled(onlyBack)
	buttonFill := tview.NewButton(tview.Escape(`[F4] Autofill`)).SetDisabled(onlyBack)
	buttonFillRev := tview.NewButton(tview.Escape(`[F5] Autofill Reverse`)).SetDisabled(onlyBack)

	gridButtons := tview.NewGrid()
	gridButtons.AddItem(buttonBack, 0, 0, 1, 1, 0, 0, false)
	gridButtons.AddItem(buttonSource, 0, 1, 1, 1, 0, 0, false)
	gridButtons.AddItem(buttonZero, 0, 2, 1, 1, 0, 0, false)
	gridButtons.AddItem(buttonFill, 0, 3, 1, 1, 0, 0, false)
	gridButtons.AddItem(buttonFillRev, 0, 4, 1, 1, 0, 0, false)

	return gridButtons
}

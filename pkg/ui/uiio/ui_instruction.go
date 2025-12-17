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
	uiInst := &UIInstruction{
		instruction: instruction,
		focus:       0,
		selectable:  []tview.Primitive{},
		app:         app,
	}

	inputs := []*UIParameterParts{}
	for _, param := range instruction.Inputs() {
		input := NewUIParameterInputs(app, param, uiInst)
		uiInst.selectable = append(uiInst.selectable, input.selectablePrimitives()...)
		inputs = append(inputs, input)
	}

	output := NewUIParameterOutputs(app, instruction.Output(), uiInst)

	gridInputs := tview.NewGrid()

	gridOutputs := tview.NewGrid()

	for i, input := range inputs {
		gridInputs.AddItem(input.GetBox(), i, 0, 1, 1, 0, 0, true)
	}

	gridOutputs.AddItem(output.GetBox(), 0, 0, 1, 1, 0, 0, false)

	gridInputOutput := tview.NewGrid()
	gridInputOutput.AddItem(gridInputs, 0, 0, 1, 1, 0, 0, true)
	gridInputOutput.AddItem(gridOutputs, 0, 1, 1, 1, 0, 0, false)

	buttonClear := tview.NewButton(tview.Escape(`[Z]ero`))
	buttonFill := tview.NewButton(tview.Escape(`A[u]tofill`))
	buttonFillRev := tview.NewButton(tview.Escape(`Autofill [R]everse`))
	buttonSource := tview.NewButton(tview.Escape(`[S]how Source`))

	gridButtons := tview.NewGrid()
	gridButtons.AddItem(buttonClear, 0, 0, 1, 1, 0, 0, false)
	gridButtons.AddItem(buttonFill, 0, 1, 1, 1, 0, 0, false)
	gridButtons.AddItem(buttonFillRev, 0, 2, 1, 1, 0, 0, false)
	gridButtons.AddItem(buttonSource, 0, 3, 1, 1, 0, 0, false)

	gridOuter := tview.NewGrid()
	gridOuter.SetBorder(true)
	gridOuter.SetTitle(instruction.Name())
	gridOuter.SetRows(8, 0)

	gridOuter.AddItem(gridInputOutput, 1, 0, 1, 1, 0, 0, true)
	gridOuter.AddItem(gridButtons, 0, 0, 1, 1, 0, 0, false)

	// Fill out the fields for the UIRegister
	uiInst.inputUIParameters = inputs
	uiInst.outputUIParameter = output
	uiInst.box = gridOuter

	// Set some default values for the inputs Often we don't even need to
	// set the inputs manually, just having unique values in each input
	// will often characterise the instruction well enough
	uiInst.setInputDefaults()

	assemblyView := tview.NewTextView()
	assemblyView.SetBorder(true)
	assemblyView.SetTitle("Go Assembly")
	assemblyView.SetText(instruction.Assembly())

	// Establish keyboard shortcuts for the buttons
	gridOuter.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Handle non-character keys first
		switch event.Key() {
		case tcell.KeyTab:
			uiInst.cycleFocus(1)
		case tcell.KeyBacktab:
			uiInst.cycleFocus(-1)
		}

		// Handle character keys
		switch event.Rune() {
		case rune('z'), rune('Z'):
			uiInst.setInputsZero()
		case rune('u'), rune('U'):
			uiInst.setInputDefaults()
		case rune('r'), rune('R'):
			uiInst.setInputDefaultsReverse()
		case rune('s'), rune('S'):
		}

		// Allow the event to propagate
		return event
	})

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

func (r *UIInstruction) initFocusCycling() {
	r.box.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			r.cycleFocus(1)
		case tcell.KeyBacktab:
			r.cycleFocus(-1)
		}

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

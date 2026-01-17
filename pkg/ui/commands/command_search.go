package commands

import (
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly"
	"github.com/fmstephe/simd_explorer/pkg/ui/stackapp"
	"github.com/fmstephe/simd_explorer/pkg/ui/uiio"
	"github.com/gdamore/tcell/v2"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/rivo/tview"
)

type CommandSearch struct {
	list  *tview.List
	input *tview.InputField
	flex  *tview.Flex
	cache map[string]*uiio.UIInstruction
}

func NewCommandSearch(instructions []assembly.Instruction, app *stackapp.StackApp) *CommandSearch {
	instMap, instNames := buildInstructionMap(instructions)

	assemblyView := tview.NewTextView()
	assemblyView.SetBorder(true)
	assemblyView.SetTitle("Go Assembly")

	list := tview.NewList()
	list.SetBorder(true)
	list.SetTitle("Choose An Instruction")
	list.ShowSecondaryText(true)

	cache := map[string]*uiio.UIInstruction{}
	// Build list of instructions
	for _, name := range instNames {
		inst := instMap[name]
		list.AddItem(name, inst.Description(), 0, buildInstructionSelectedFunc(app, inst, cache))
	}

	// When a new list item is selected, update the assembly view to
	// display the assembly for that instruction
	list.SetChangedFunc(func(index int, mainText string, _ string, _ rune) {
		assemblyView.SetText(instMap[mainText].Assembly())
	})

	// Set the assembly view from the first instruction in the list
	if len(instNames) > 0 {
		assemblyView.SetText(instMap[instNames[0]].Assembly())
	}

	input := tview.NewInputField()
	input.SetBorder(true)
	input.SetTitle("Instruction Filter")

	input.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			listIdx := list.GetCurrentItem()
			log.Printf("Enter Pressed %d\n", listIdx)
			list.GetItemSelectedFunc(listIdx)()
			// Short-circuit default handling of the enter key
			return nil
		case tcell.KeyUp, tcell.KeyBacktab:
			listIdx := list.GetCurrentItem()
			list.SetCurrentItem(listIdx - 1)
			// Short-circuit default handling of the up/shift+tab keys
			return nil
		case tcell.KeyDown, tcell.KeyTAB:
			listIdx := list.GetCurrentItem()
			list.SetCurrentItem(listIdx + 1)
			// Short-circuit default handling of the down/tab keys
			return nil
		}
		return event
	})

	input.SetChangedFunc(func(txt string) {
		// When the input-field is updated, update the list with
		// filtered instructions, using case-insensitive search
		found := fuzzy.FindFold(txt, instNames)
		list.Clear()
		for _, name := range found {
			inst := instMap[name]
			list.AddItem(name, inst.Description(), 0, buildInstructionSelectedFunc(app, inst, cache))
		}
	})

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)

	flex.AddItem(input, 3, 1, true)
	flex.AddItem(list, 0, 2, false)
	flex.AddItem(assemblyView, 0, 3, false)
	return &CommandSearch{
		list:  list,
		input: input,
		flex:  flex,
		cache: cache,
	}
}

func (s *CommandSearch) GetBox() *tview.Flex {
	return s.flex
}

func buildInstructionMap(instructions []assembly.Instruction) (instMap map[string]assembly.Instruction, instNames []string) {
	instMap = map[string]assembly.Instruction{}
	instNames = []string{}
	for _, inst := range instructions {
		name := inst.Name()
		if !inst.Supported() {
			name = name + " (Not Supported!)"
		}
		instMap[name] = inst
		instNames = append(instNames, name)
	}

	return instMap, instNames
}

func buildInstructionSelectedFunc(app *stackapp.StackApp, inst assembly.Instruction, cache map[string]*uiio.UIInstruction) func() {
	return func() {
		log.Printf("Chosen %s", inst.Name())
		if !inst.Supported() {
			// If the instruction is not supported, do nothing
			return
		}

		uiInst, ok := cache[inst.Name()]
		if !ok {
			// If the ui-instruction is not in the cache, create it
			// and cache it
			uiInst = uiio.NewUIInstruction(app, inst)
			cache[inst.Name()] = uiInst
		}

		app.Push(uiInst.GetPrimitive())
	}
}

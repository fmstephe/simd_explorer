package commands

import (
	"github.com/gdamore/tcell/v2"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/rivo/tview"
)

type CommandSearch struct {
	list     *tview.List
	input    *tview.InputField
	grid     *tview.Grid
	commands []string
}

func NewCommandSearch(commands []string, app *tview.Application) *CommandSearch {
	list := tview.NewList()
	list.SetBorder(true)
	list.SetTitle("Choose An Instruction")
	list.ShowSecondaryText(false)
	for _, cmd := range commands {
		list.AddItem(cmd, "", 0, func() {})
	}

	input := tview.NewInputField()
	input.SetBorder(true)
	input.SetTitle("Instruction Filter")
	input.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		// TAB or SHIFT+TAB cycles focus to the list
		case tcell.KeyTAB, tcell.KeyBacktab:
			app.SetFocus(list)
		}
		return event
	})
	input.SetChangedFunc(func(txt string) {
		// When the input-field is updated, update the list with
		// filtered instructions
		found := fuzzy.Find(txt, commands)
		list.Clear()
		for _, cmd := range found {
			list.AddItem(cmd, "", 0, func() {})
		}
	})

	// Text input directed to the list is appended to the inputfield This
	// allows text input to continue to work even if the list itself has
	// been tabbed to or clicked on
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		// Deletions cause a deletion from the input-field
		case tcell.KeyDelete, tcell.KeyDEL:
			input.SetText(deleteFrom(input.GetText()))
		// Text input is appended to the input-field
		case tcell.KeyRune:
			input.SetText(input.GetText() + string(event.Rune()))
		// TAB or SHIFT+TAB cycles focus to the input-field
		case tcell.KeyTAB, tcell.KeyBacktab:
			app.SetFocus(input)
			// We interrupt the event processing here, otherwise
			// the list will interpret TAB as navigating the list
			return nil
		}

		return event
	})

	grid := tview.NewGrid()
	grid.AddItem(list, 0, 0, 1, 1, 0, 0, false)
	grid.AddItem(input, 1, 0, 1, 1, 0, 0, true)
	return &CommandSearch{
		list:     list,
		input:    input,
		grid:     grid,
		commands: commands,
	}
}

func (s *CommandSearch) GetBox() *tview.Grid {
	return s.grid
}

func deleteFrom(txt string) string {
	if txt == "" {
		return ""
	}
	return txt[:len(txt)-1]
}

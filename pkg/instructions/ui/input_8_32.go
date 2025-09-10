package ui

import "github.com/rivo/tview"

type input8X32 struct {
	app *tview.Application

	focus          int
	inputsForFocus []*tview.InputField

	input  *tview.Flex
	output *tview.Flex
}

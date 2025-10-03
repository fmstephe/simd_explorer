package uiio

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type uiParameterPart interface {
	getText() string
	setText(string)
	setTitle(string)
	setBorder(bool)
	setFieldWidth(int)
	setAcceptanceFunc(func(string, rune) bool)
	setChangedFunc(func(string))
	setBackgroundColor(tcell.Color)

	primitive() tview.Primitive
}

type uiPartBuilder interface {
	build() uiParameterPart
	kind() string
}

// builder for input-field parts
type inputPartBuilder struct {
}

func (i *inputPartBuilder) build() uiParameterPart {
	return newInputPart()
}

func (i *inputPartBuilder) kind() string {
	return "input"
}

// uiPart for InputField
type inputPart struct {
	input *tview.InputField
}

func newInputPart() *inputPart {
	return &inputPart{
		input: tview.NewInputField(),
	}
}

func (i *inputPart) getText() string {
	return i.input.GetText()
}

func (i *inputPart) setText(txt string) {
	i.input.SetText(txt)
}

func (i *inputPart) setTitle(title string) {
	i.input.SetTitle(title)
}

func (i *inputPart) setBorder(showBorder bool) {
	i.input.SetBorder(showBorder)
}

func (i *inputPart) setFieldWidth(width int) {
	i.input.SetFieldWidth(width)
}

func (i *inputPart) setAcceptanceFunc(af func(string, rune) bool) {
	i.input.SetAcceptanceFunc(af)
}

func (i *inputPart) setChangedFunc(cf func(string)) {
	i.input.SetChangedFunc(cf)
}

func (i *inputPart) setBackgroundColor(color tcell.Color) {
	i.input.SetBackgroundColor(color)
}

func (i *inputPart) primitive() tview.Primitive {
	return i.input
}

// builder for text-view parts
type textViewPartBuilder struct {
}

func (t *textViewPartBuilder) build() uiParameterPart {
	return newTextViewPart()
}

func (t *textViewPartBuilder) kind() string {
	return "text"
}

// uiPart for TextView
type textViewPart struct {
	textView *tview.TextView
}

func newTextViewPart() *textViewPart {
	return &textViewPart{
		textView: tview.NewTextView(),
	}
}

func (i *textViewPart) getText() string {
	return i.textView.GetText(false)
}

func (i *textViewPart) setText(txt string) {
	i.textView.SetText(txt)
}

func (i *textViewPart) setTitle(title string) {
	i.textView.SetTitle(title)
}

func (i *textViewPart) setBorder(showBorder bool) {
	i.textView.SetBorder(showBorder)
}

func (i *textViewPart) setFieldWidth(width int) {
	i.textView.SetSize(1, width)
}

func (i *textViewPart) setAcceptanceFunc(af func(string, rune) bool) {
	// No op
}

func (i *textViewPart) setChangedFunc(cf func(string)) {
	// No op
}

func (i *textViewPart) setBackgroundColor(color tcell.Color) {
	i.textView.SetBackgroundColor(color)
}

func (i *textViewPart) primitive() tview.Primitive {
	return i.textView
}

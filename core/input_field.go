package core

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type InputField struct {
	inputField *tview.InputField
}

func NewInputField(label string, width int, handler func(textToCheck string, lastChar rune) bool) *InputField {
	i := &InputField{
		inputField: tview.NewInputField(),
	}
	i.inputField.SetLabel(label)
	i.inputField.SetFieldWidth(width)
	i.inputField.SetAcceptanceFunc(handler)
	return i
}
func (i *InputField) GetInputField() *tview.InputField {
	return i.inputField
}
func (i *InputField) SetBackgroundColor(color tcell.Color) {
	i.inputField.SetBackgroundColor(color)
}
func (i *InputField) SetDoneFunc(handler func(key tcell.Key)) {
	i.inputField.SetDoneFunc(handler)
}
func (i *InputField) SetAutocompleteFunc(callback func(currentText string) (entries []string)) {
	i.inputField.SetAutocompleteFunc(callback)
}
func (i *InputField) SetAutocompletedFunc(autocompleted func(text string, index int, source int) bool) {
	i.inputField.SetAutocompletedFunc(autocompleted)
}
func (i *InputField) SetText(text string) {
	i.inputField.SetText(text)
}
func (i *InputField) GetText() string {
	return i.inputField.GetText()
}

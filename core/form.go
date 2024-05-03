package core

import "github.com/rivo/tview"

type Form struct {
	form       *tview.Form
	border     bool
	title      string
	titleAlign int
}

func NewForm(border bool, title string, titleAlign int) *Form {
	f := &Form{
		form:       tview.NewForm(),
		border:     border,
		title:      title,
		titleAlign: titleAlign,
	}
	f.form.SetBorder(f.border).SetTitle(f.title).SetTitleAlign(f.titleAlign)
	return f
}
func (f *Form) GetForm() *tview.Form {
	return f.form
}
func (f *Form) AddDropDown(label string, options []string, initialOption int, selected func(option string, optionIndex int)) {
	f.form.AddDropDown(label, options, initialOption, selected)
}
func (f *Form) AddInputField(label string, value string, fieldWidth int, accept func(textToCheck string, lastChar rune) bool, changed func(text string)) {
	f.form.AddInputField(label, value, fieldWidth, accept, changed)
}
func (f *Form) AddTextArea(label string, text string, fieldWidth int, fieldHeight int, maxLength int, changed func(text string)) {
	f.form.AddTextArea(label, text, fieldWidth, fieldHeight, maxLength, changed)
}
func (f *Form) AddTextView(label string, text string, fieldWidth int, fieldHeight int, dynamicColors bool, scrollable bool) {
	f.form.AddTextView(label, text, fieldWidth, fieldHeight, dynamicColors, scrollable)
}
func (f *Form) AddCheckbox(label string, checked bool, changed func(checked bool)) {
	f.form.AddCheckbox(label, checked, changed)
}
func (f *Form) AddPasswordField(label string, value string, fieldWidth int, mask rune, changed func(text string)) {
	f.form.AddPasswordField(label, value, fieldWidth, mask, changed)
}
func (f *Form) AddButton(label string, selected func()) {
	f.form.AddButton(label, selected)
}
func (f *Form) GetFormItemDropDown(label string) *tview.DropDown {
	return f.form.GetFormItemByLabel(label).(*tview.DropDown)
}
func (f *Form) GetFormItemInputField(label string) *tview.InputField {
	return f.form.GetFormItemByLabel(label).(*tview.InputField)
}

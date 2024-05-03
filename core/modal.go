package core

import "github.com/rivo/tview"

type Modal struct {
	modal *tview.Modal
}

func NewModal() *Modal {
	p := &Modal{
		modal: tview.NewModal(),
	}
	return p
}
func (m *Modal) GetModal() *tview.Modal {
	return m.modal
}
func (m *Modal) SetText(text string) {
	m.modal.SetText(text)
}
func (m *Modal) AddButtons(labels []string) {
	m.modal.AddButtons(labels)
}
func (m *Modal) SetDoneFunc(handler func(buttonIndex int, buttonLabel string)) {
	m.modal.SetDoneFunc(handler)
}

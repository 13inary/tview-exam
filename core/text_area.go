package core

import "github.com/rivo/tview"

type TextArea struct {
	textArea    *tview.TextArea
	placeholder string
	title       string
	border      bool
}

func NewTextArea(placeholder string, title string, border bool) *TextArea {
	t := &TextArea{
		textArea:    tview.NewTextArea(),
		placeholder: placeholder,
		title:       title,
		border:      border,
	}
	t.textArea.SetPlaceholder(t.placeholder)
	t.textArea.SetTitle(t.title)
	t.textArea.SetBorder(t.border)
	return t
}
func (t *TextArea) GetTextArea() *tview.TextArea {
	return t.textArea
}
func (t *TextArea) SetMovedFunc(moveFunc func()) {
	t.textArea.SetMovedFunc(moveFunc)
}
func (t *TextArea) GetCursor() (int, int, int, int) {
	return t.textArea.GetCursor()
}

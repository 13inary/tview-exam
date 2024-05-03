package core

import (
	"github.com/rivo/tview"
)

type Checkbox struct {
	checkbox *tview.Checkbox
}

func NewCheckbox(label string) *Checkbox {
	c := &Checkbox{
		checkbox: tview.NewCheckbox(),
	}
	c.checkbox.SetLabel(label)
	return c
}
func (b *Checkbox) GetCheckbox() *tview.Checkbox {
	return b.checkbox
}

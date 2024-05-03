package core

import (
	"github.com/rivo/tview"
)

type DropDown struct {
	dropDown *tview.DropDown
}

func NewDropDown(label string) *DropDown {
	d := &DropDown{
		dropDown: tview.NewDropDown(),
	}
	d.dropDown.SetLabel(label)
	return d
}
func (d *DropDown) GetDropDown() *tview.DropDown {
	return d.dropDown
}

func (d *DropDown) SetOptions(texts []string, selected func(text string, index int)) {
	d.dropDown.SetOptions(texts, selected)
}
func (d *DropDown) GetCurrentOption() (int, string) {
	return d.dropDown.GetCurrentOption()
}

package core

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Box struct {
	box *tview.Box
}

func NewBox(title string, border bool) *Box {
	b := &Box{
		box: tview.NewBox(),
	}
	b.box.SetBorder(border)
	b.box.SetTitle(title)
	return b
}
func (b *Box) GetBox() *tview.Box {
	return b.box
}
func (b *Box) SetBackgroundColor(color tcell.Color) {
	b.box.SetBackgroundColor(color)
}

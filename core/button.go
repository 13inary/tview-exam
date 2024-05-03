package core

import (
	"github.com/rivo/tview"
)

type Button struct {
	button *tview.Button
	x      int
	y      int
	width  int
	height int
}

func NewButton(label string, border bool, x int, y int, width int, height int) *Button {
	b := &Button{
		button: tview.NewButton(label),
		x:      x,
		y:      y,
		width:  width,
		height: height,
	}
	b.button.SetBorder(border)
	b.button.SetRect(b.x, b.y, b.width, b.height)
	return b
}
func (b *Button) GetButton() *tview.Button {
	return b.button
}
func (b *Button) SetSelectedFunc(handler func()) {
	b.button.SetSelectedFunc(handler)
}
func (b *Button) SetBorder(show bool) {
	b.button.SetBorder(show)
}

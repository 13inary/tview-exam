package core

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Frame struct {
	frame  *tview.Frame
	title  string
	border bool
	top    int
	bottom int
	header int
	footer int
	left   int
	right  int
}

func NewFrame(p tview.Primitive, title string, border bool, top int, bottom int, header int, footer int, left int, right int) *Frame {
	f := &Frame{
		frame:  tview.NewFrame(p),
		title:  title,
		border: border,
		top:    top,
		bottom: bottom,
		header: header,
		footer: footer,
		left:   left,
		right:  right,
	}
	f.frame.SetBorders(f.top, f.bottom, f.header, f.footer, f.left, f.right)
	f.frame.SetBorder(f.border)
	f.frame.SetTitle(f.title)
	return f
}
func (f *Frame) GetFrame() *tview.Frame {
	return f.frame
}
func (f *Frame) SetInputCapture(inputCapture func(event *tcell.EventKey) *tcell.EventKey) {
	f.frame.SetInputCapture(inputCapture)
}
func (f *Frame) GetPrimitive() tview.Primitive {
	return f.frame.GetPrimitive()
}
func (f *Frame) SetPrimitive(p tview.Primitive) {
	f.frame.SetPrimitive(p)
}
func (f *Frame) AddText(text string, header bool, align int, color tcell.Color) {
	f.frame.AddText(text, header, align, color)
}

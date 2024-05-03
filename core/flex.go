package core

import (
	"github.com/rivo/tview"
)

type Flex struct {
	flex *tview.Flex
}

func NewFlex() *Flex {
	f := &Flex{
		flex: tview.NewFlex(),
	}
	return f
}
func (f *Flex) GetFlex() *tview.Flex {
	return f.flex
}
func (f *Flex) AddItem(item tview.Primitive, fixedSize int, proportion int, focus bool) {
	f.flex.AddItem(item, fixedSize, proportion, focus)
}
func (f *Flex) SetDirection(direction int) {
	f.flex.SetDirection(direction)
}

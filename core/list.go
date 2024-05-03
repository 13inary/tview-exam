package core

import "github.com/rivo/tview"

type List struct {
	list *tview.List
}

func NewList() *List {
	l := &List{
		list: tview.NewList(),
	}
	return l
}

func (l *List) AddItem(mainText string, secondaryText string, shortcut rune, selected func()) {
	l.list.AddItem(mainText, secondaryText, shortcut, selected)
}

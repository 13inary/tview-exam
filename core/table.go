package core

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Table struct {
	table  *tview.Table
	border bool
}

func NewTable(border bool) *Table {
	t := &Table{
		table:  tview.NewTable(),
		border: border,
	}
	t.table.SetBorders(t.border)
	return t
}
func (t *Table) GetTable() *tview.Table {
	return t.table
}
func (t *Table) SetCell(row int, column int, text string, color tcell.Color, align int) {
	cell := tview.NewTableCell(text)
	cell.SetTextColor(color)
	cell.SetAlign(align)
	t.table.SetCell(row, column, cell)
}
func (t *Table) Select(row int, column int) {
	t.table.Select(row, column)
}
func (t *Table) SetFixed(rows int, columns int) {
	t.table.SetFixed(rows, columns)
}
func (t *Table) SetDoneFunc(handler func(key tcell.Key)) {
	t.table.SetDoneFunc(handler)
}
func (t *Table) SetSelectedFunc(handler func(row, column int)) {
	t.table.SetSelectedFunc(handler)
}
func (t *Table) SetTextColor(row int, column int, color tcell.Color) {
	t.table.GetCell(row, column).SetTextColor(color)
}
func (t *Table) SetSelectable(rows bool, columns bool) {
	t.table.SetSelectable(rows, columns)
}

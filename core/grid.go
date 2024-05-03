package core

import "github.com/rivo/tview"

type Grid struct {
	grid *tview.Grid
}

func NewGrid() *Grid {
	g := &Grid{
		grid: tview.NewGrid(),
	}
	return g
}
func (g *Grid) GetGrid() *tview.Grid {
	return g.grid
}
func (g *Grid) SetBorders(show bool) {
	g.grid.SetBorders(show)
}
func (g *Grid) SetColumns(columns ...int) {
	g.grid.SetColumns(columns...)
}
func (g *Grid) SetRows(rows ...int) {
	g.grid.SetRows(rows...)
}
func (g *Grid) AddItem(p tview.Primitive, row int, column int, rowSpan int, colSpan int, minGridHeight int, minGridWidth int, focus bool) {
	g.grid.AddItem(p, row, column, rowSpan, colSpan, minGridHeight, minGridWidth, focus)
}

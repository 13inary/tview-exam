package modules

import (
	"tview-exam/core"

	"github.com/rivo/tview"
)

func BuildGrid(app *core.App) *core.Grid {
	menu := core.NewTextView(true, true, false, tview.AlignCenter)
	menu.SetText("menu")
	mai := core.NewTextView(true, true, false, tview.AlignCenter)
	mai.SetText("main")
	side := core.NewTextView(true, true, false, tview.AlignCenter)
	side.SetText("side")
	head := core.NewTextView(true, true, false, tview.AlignCenter)
	head.SetText("head")
	foot := core.NewTextView(true, true, false, tview.AlignCenter)
	foot.SetText("foot")

	grid := core.NewGrid()
	grid.SetRows(3, 0, 3)
	grid.SetColumns(30, 0, 30)
	grid.SetBorders(true)
	grid.AddItem(head.GetTextView(), 0, 0, 1, 3, 0, 0, false)
	grid.AddItem(foot.GetTextView(), 2, 0, 1, 3, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(menu.GetTextView(), 0, 0, 0, 0, 0, 0, false)
	grid.AddItem(mai.GetTextView(), 1, 0, 1, 3, 0, 0, false)
	grid.AddItem(side.GetTextView(), 0, 0, 0, 0, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(menu.GetTextView(), 1, 0, 1, 1, 0, 100, false)
	grid.AddItem(mai.GetTextView(), 1, 1, 1, 1, 0, 100, false)
	grid.AddItem(side.GetTextView(), 1, 2, 1, 1, 0, 100, false)

	return grid
}

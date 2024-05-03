package modules

import (
	"strings"
	"tview-exam/core"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func BuildModal2(app *core.App) *core.Pages {
	// Returns a new primitive which puts the provided primitive in the center and
	// sets its size to the given width and height.
	modal2 := func(p tview.Primitive, width, height int) tview.Primitive {
		m2 := core.NewFlex()
		m2.SetDirection(tview.FlexRow)
		m2.AddItem(nil, 0, 1, false)
		m2.AddItem(p, height, 1, true)
		m2.AddItem(nil, 0, 1, false)

		m := core.NewFlex()
		m.AddItem(nil, 0, 1, false)
		m.AddItem(m2.GetFlex(), width, 1, true)
		m.AddItem(nil, 0, 1, false)
		return m.GetFlex()
	}

	// Returns a new primitive which puts the provided primitive in the center and
	// sets its size to the given width and height.
	//modal3 := func(p tview.Primitive, width, height int) tview.Primitive {
	//	return tview.NewGrid().
	//		SetColumns(0, width, 0).
	//		SetRows(0, height, 0).
	//		AddItem(p, 1, 1, 1, 1, 0, 0, true)
	//}

	background := core.NewTextView(true, true, true, tview.AlignCenter)
	background.SetTextColor(tcell.ColorBlue)
	background.SetText(strings.Repeat("background ", 1000))

	box := core.NewBox("Centered Box", true)

	pages := core.NewPages()
	pages.AddPage("background", background.GetTextView(), true, true)
	pages.AddPage("modal", modal2(box.GetBox(), 40, 10), true, true)
	return pages
}

// alert shows a confirmation dialog.
func alert(pages *tview.Pages, id string, message string) *tview.Pages {
	return pages.AddPage(
		id,
		tview.NewModal().
			SetText(message).
			AddButtons([]string{"确定"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				pages.HidePage(id).RemovePage(id)
			}),
		false,
		true,
	)
}

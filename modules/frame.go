package modules

import (
	"tview-exam/core"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func BuildFrame(app *core.App) *core.Frame {
	plug := core.NewBox("", false)
	plug.SetBackgroundColor(tcell.ColorBlue)

	frame := core.NewFrame(plug.GetBox(), "", false, 2, 2, 2, 2, 4, 4)
	frame.AddText("Header left", true, tview.AlignLeft, tcell.ColorWhite)
	frame.AddText("Header middle", true, tview.AlignCenter, tcell.ColorWhite)
	frame.AddText("Header right", true, tview.AlignRight, tcell.ColorWhite)
	frame.AddText("Header second middle", true, tview.AlignCenter, tcell.ColorRed)
	frame.AddText("Footer middle", false, tview.AlignCenter, tcell.ColorGreen)
	frame.AddText("Footer second middle", false, tview.AlignCenter, tcell.ColorGreen)
	return frame
}

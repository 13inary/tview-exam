package modules

import (
	"tview-exam/core"
)

func BuildList(app *core.App) *core.List {
	list := core.NewList()
	list.AddItem("List item 1", "Some explanatory text", 'a', nil)
	list.AddItem("List item 2", "Some explanatory text", 'b', nil)
	list.AddItem("List item 3", "Some explanatory text", 'c', nil)
	list.AddItem("List item 4", "Some explanatory text", 'd', nil)
	list.AddItem("Quit", "Press to exit", 'q', func() {
		app.GetApp().Stop()
	})
	return list
}

package modules

import (
	"tview-exam/core"

	"github.com/rivo/tview"
)

func BuildForm(app *core.App) *core.Form {
	form := core.NewForm(true, "title", tview.AlignLeft)
	form.AddDropDown("Title", []string{"Mr.", "Ms.", "Mrs.", "Dr.", "Prof."}, 0, nil)
	form.AddInputField("First name", "", 20, nil, nil)
	form.AddInputField("Last name", "", 20, nil, nil)
	form.AddTextArea("Address", "", 40, 0, 0, nil)
	form.AddTextView("Notes", "This is just a demo.\nYou can enter whatever you wish.", 40, 2, true, false)
	form.AddCheckbox("Age 18+", false, nil)
	form.AddPasswordField("Password", "", 10, '*', nil)
	form.AddButton("Save", nil)
	form.AddButton("Quit", func() {
		app.GetApp().Stop()
	})
	return form
}

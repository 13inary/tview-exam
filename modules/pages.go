package modules

import (
	"fmt"
	"tview-exam/core"
)

func BuildPages(app *core.App) *core.Pages {
	pages := core.NewPages()
	pageCount := 5

	for page := 0; page < pageCount; page++ {
		modal := core.NewModal()
		modal.SetText(fmt.Sprintf("This is page %d. Choose where to go next.", page+1))
		modal.AddButtons([]string{"Next", "Quit"})

		func(page int) {
			modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonIndex == 0 {
					pages.SwitchToPage(fmt.Sprintf("page-%d", (page+1)%pageCount))
				} else {
					app.GetApp().Stop()
				}
			})
			pages.AddPage(fmt.Sprintf("page-%d", page), modal.GetModal(), false, page == 0)
		}(page)

	}

	return pages
}

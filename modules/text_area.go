package modules

import (
	"fmt"
	"tview-exam/core"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func BuildTextArea(app *core.App) *core.Pages {
	textArea := core.NewTextArea("Enter text here...", "Text Area", true)

	helpInfo := core.NewTextView(true, true, false, tview.AlignLeft)
	helpInfo.SetText(" Press F1 for help, press Ctrl-C to exit")

	position := core.NewTextView(true, true, false, tview.AlignRight)

	moveFunc := func() {
		fromRow, fromColumn, toRow, toColumn := textArea.GetCursor()
		if fromRow == toRow && fromColumn == toColumn {
			position.SetText(fmt.Sprintf("Row: [yellow]%d[white], Column: [yellow]%d ", fromRow, fromColumn))
		} else {
			position.SetText(fmt.Sprintf("[red]From[white] Row: [yellow]%d[white], Column: [yellow]%d[white] - [red]To[white] Row: [yellow]%d[white], To Column: [yellow]%d ", fromRow, fromColumn, toRow, toColumn))
		}
	}
	textArea.SetMovedFunc(moveFunc)
	moveFunc()

	mainView := core.NewGrid()
	mainView.SetRows(0, 1)
	mainView.AddItem(textArea.GetTextArea(), 0, 0, 1, 2, 0, 0, true)
	mainView.AddItem(helpInfo.GetTextView(), 1, 0, 1, 1, 0, 0, false)
	mainView.AddItem(position.GetTextView(), 1, 1, 1, 1, 0, 0, false)

	pages := core.NewPages()

	help1 := core.NewTextView(true, true, false, tview.AlignLeft)
	help1.SetText(`[green]Navigation

[yellow]Left arrow[white]: Move left.
[yellow]Right arrow[white]: Move right.
[yellow]Down arrow[white]: Move down.
[yellow]Up arrow[white]: Move up.
[yellow]Ctrl-A, Home[white]: Move to the beginning of the current line.
[yellow]Ctrl-E, End[white]: Move to the end of the current line.
[yellow]Ctrl-F, page down[white]: Move down by one page.
[yellow]Ctrl-B, page up[white]: Move up by one page.
[yellow]Alt-Up arrow[white]: Scroll the page up.
[yellow]Alt-Down arrow[white]: Scroll the page down.
[yellow]Alt-Left arrow[white]: Scroll the page to the left.
[yellow]Alt-Right arrow[white]:  Scroll the page to the right.
[yellow]Alt-B, Ctrl-Left arrow[white]: Move back by one word.
[yellow]Alt-F, Ctrl-Right arrow[white]: Move forward by one word.

[blue]Press Enter for more help, press Escape to return.`)

	help2 := core.NewTextView(true, true, false, tview.AlignLeft)
	help2.SetText(`[green]Editing[white]

Type to enter text.
[yellow]Ctrl-H, Backspace[white]: Delete the left character.
[yellow]Ctrl-D, Delete[white]: Delete the right character.
[yellow]Ctrl-K[white]: Delete until the end of the line.
[yellow]Ctrl-W[white]: Delete the rest of the word.
[yellow]Ctrl-U[white]: Delete the current line.

[blue]Press Enter for more help, press Escape to return.`)

	help3 := core.NewTextView(true, true, false, tview.AlignLeft)
	help3.SetText(`[green]Selecting Text[white]

Move while holding Shift or drag the mouse.
Double-click to select a word.

[green]Clipboard

[yellow]Ctrl-Q[white]: Copy.
[yellow]Ctrl-X[white]: Cut.
[yellow]Ctrl-V[white]: Paste.

[green]Undo

[yellow]Ctrl-Z[white]: Undo.
[yellow]Ctrl-Y[white]: Redo.

[blue]Press Enter for more help, press Escape to return.`)

	help := core.NewFrame(help1.GetTextView(), "Help", true, 1, 1, 0, 0, 2, 2)
	help.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			pages.SwitchToPage("main")
			return nil
		} else if event.Key() == tcell.KeyEnter {
			switch {
			case help.GetPrimitive() == help1.GetTextView():
				help.SetPrimitive(help2.GetTextView())
			case help.GetPrimitive() == help2.GetTextView():
				help.SetPrimitive(help3.GetTextView())
			case help.GetPrimitive() == help3.GetTextView():
				help.SetPrimitive(help1.GetTextView())
			}
			return nil
		}
		return event
	})

	tmp := core.NewGrid()
	tmp.SetColumns(0, 64, 0)
	tmp.SetRows(0, 22, 0)
	tmp.AddItem(help.GetFrame(), 1, 1, 1, 1, 0, 0, true)
	pages.AddAndSwitchToPage("main", mainView.GetGrid(), true)
	pages.AddPage("help", tmp.GetGrid(), true, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyF1 {
			pages.ShowPage("help") //TODO: Check when clicking outside help window with the mouse. Then clicking help again.
			return nil
		}
		return event
	})

	return pages
}

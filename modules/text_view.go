package modules

import (
	"fmt"
	"strings"
	"time"
	"tview-exam/core"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const corporate = `Leverage agile frameworks to provide a robust synopsis for high level overviews. Iterative approaches to corporate strategy foster collaborative thinking to further the overall value proposition. Organically grow the holistic world view of disruptive innovation via workplace diversity and empowerment.

Bring to the table win-win survival strategies to ensure proactive domination. At the end of the day, going forward, a new normal that has evolved from generation X is on the runway heading towards a streamlined cloud solution. User generated content in real-time will have multiple touchpoints for offshoring.

Capitalize on low hanging fruit to identify a ballpark value added activity to beta test. Override the digital divide with additional clickthroughs from DevOps. Nanotechnology immersion along the information highway will close the loop on focusing solely on the bottom line.

[yellow]Press Enter, then Tab/Backtab for word selections
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
aa
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
`

func BuildTextView(app *core.App) *core.TextView {
	textView := core.NewTextView(true, true, true, tview.AlignLeft)
	textView.SetChangedFunc(func() {
		app.GetApp().Draw()
	})
	// show 1
	textView.ShowText(true, func(t *tview.TextView, numSelections *int) {
		for _, word := range strings.Split(corporate, " ") {
			if word == "the" {
				word = "[red]the[white]"
			}
			if word == "to" {
				word = fmt.Sprintf(`["%d"]to[""]`, *numSelections)
				*numSelections++
			}
			fmt.Fprintf(t, "%s ", word)
			time.Sleep(200 * time.Millisecond)
		}
	})
	// show 2
	//go func() {
	//	w := tview.ANSIWriter(textView.GetTextView())
	//	if _, err := io.Copy(w, os.Stdin); err != nil {
	//		panic(err)
	//	}
	//}()
	textView.SetDoneFunc(func(key tcell.Key) {
		switch {
		case key == tcell.KeyEnter:
			if textView.IsHighlight() {
				textView.CloseHighlight()
			} else {
				textView.ScrollToHighlight(0)
			}
		case key == tcell.KeyTab:
			if textView.IsHighlight() {
				textView.ScrollToHighlightNext()
			}
		case key == tcell.KeyBacktab:
			if textView.IsHighlight() {
				textView.ScrollToHighlightPrefix()
			}
		default:
		}
	})
	return textView
}

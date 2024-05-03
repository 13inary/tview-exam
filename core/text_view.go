package core

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TextView struct {
	textView      *tview.TextView
	dynamicColors bool
	regions       bool
	border        bool
	numSelections *int
	showDone      chan struct{}
}

func NewTextView(dynamicColors bool, regions bool, border bool, textAlign int) *TextView {
	numSelections := 0
	t := &TextView{
		textView:      tview.NewTextView(),
		dynamicColors: dynamicColors,
		regions:       regions,
		border:        border,
		numSelections: &numSelections,
		showDone:      make(chan struct{}, 1),
	}
	t.textView.SetDynamicColors(t.dynamicColors).SetRegions(t.regions).SetBorder(t.border)
	t.textView.SetTextAlign(textAlign)
	return t
}
func (t *TextView) GetTextView() *tview.TextView {
	return t.textView
}
func (t *TextView) SetChangedFunc(handler func()) {
	t.textView.SetChangedFunc(handler)
}
func (t *TextView) SetText(text string) {
	t.textView.SetText(text)
}
func (t *TextView) ShowText(async bool, printText func(t *tview.TextView, numSelections *int)) {
	show := func() {
		printText(t.textView, t.numSelections)
		close(t.showDone)
	}
	if async {
		go show()
		return
	}
	show()
}
func (t *TextView) SetDoneFunc(doneFuc func(key tcell.Key)) {
	t.textView.SetDoneFunc(doneFuc)
}
func (t *TextView) IsHighlight() bool {
	currentSelection := t.textView.GetHighlights()
	return len(currentSelection) > 0
}
func (t *TextView) CloseHighlight() {
	t.textView.Highlight()
}
func (t *TextView) ScrollToHighlight(index int) {
	t.textView.Highlight(strconv.Itoa(index)).ScrollToHighlight()
}
func (t *TextView) ScrollToHighlightPrefix() {
	currentSelection := t.textView.GetHighlights()
	if len(currentSelection) <= 0 {
		t.ScrollToHighlight(0)
	}
	currentIndex, _ := strconv.Atoi(currentSelection[0])
	newIndex := (currentIndex - 1 + *t.numSelections) % *t.numSelections
	t.ScrollToHighlight(newIndex)
}
func (t *TextView) ScrollToHighlightNext() {
	currentSelection := t.textView.GetHighlights()
	if len(currentSelection) <= 0 {
		t.ScrollToHighlight(0)
	}
	currentIndex, _ := strconv.Atoi(currentSelection[0])
	newIndex := (currentIndex + 1) % *t.numSelections
	t.ScrollToHighlight(newIndex)
}
func (t *TextView) SetTextColor(color tcell.Color) {
	t.textView.SetTextColor(color)
}

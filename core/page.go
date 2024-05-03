package core

import "github.com/rivo/tview"

type Pages struct {
	pages *tview.Pages
}

func NewPages() *Pages {
	p := &Pages{
		pages: tview.NewPages(),
	}
	return p
}
func (p *Pages) GetPages() *tview.Pages {
	return p.pages
}
func (p *Pages) AddAndSwitchToPage(name string, item tview.Primitive, resize bool) {
	p.pages.AddAndSwitchToPage(name, item, resize)
}
func (p *Pages) AddPage(name string, item tview.Primitive, resize bool, visible bool) {
	p.pages.AddPage(name, item, resize, visible)
}
func (p *Pages) SwitchToPage(name string) {
	p.pages.SwitchToPage(name)
}
func (p *Pages) ShowPage(name string) {
	p.pages.ShowPage(name)
}
func (p *Pages) HidePage(name string) {
	p.pages.HidePage(name)
}
func (p *Pages) RemovePage(name string) {
	p.pages.RemovePage(name)
}

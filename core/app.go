package core

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type App struct {
	app         *tview.Application
	fullScreen  bool
	enableMouse bool
}

func NewApp(fullScreen bool, enableMouse bool) *App {
	a := &App{
		app:         tview.NewApplication(),
		fullScreen:  fullScreen,
		enableMouse: enableMouse,
	}
	return a
}
func (a *App) GetApp() *tview.Application {
	return a.app
}
func (a *App) SetInputCapture(inputCapture func(event *tcell.EventKey) *tcell.EventKey) {
	a.app.SetInputCapture(inputCapture)
}
func (a *App) Run(root tview.Primitive, focus tview.Primitive) error {
	a.app.SetRoot(root, a.fullScreen)
	a.app.EnableMouse(a.enableMouse)
	a.app.SetFocus(focus)
	return a.app.Run()
}

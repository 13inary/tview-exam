package main

import (
	"tview-exam/core"
	"tview-exam/modules"
)

func main() {
	app := core.NewApp(true, true)
	//box := modules.NewBox("title", true)
	//list := modules.NewList(app)
	//form := core.BuildForm(app)
	textView := modules.BuildTextView(app)
	//area := core.BuildTextArea(app)
	//table := core.BuildTable(app)
	//treeView := core.BuildTreeView(app)
	//grid := core.BuildGrid(app)
	//flex := core.BuildFlex(app)
	//pages := core.BuildPages(app)
	//frame := core.BuildFrame(app)
	//button := modules.BuildButton(app)
	//checkbox := modules.BuildCheckbox(app)
	//dropDown := modules.BuildDropDown(app)
	//inputFiled := modules.BuildInputField(app)
	//modal := modules.BuildModal(app)
	if err := app.Run(textView.GetTextView(), textView.GetTextView()); err != nil {
		panic(err)
	}
}

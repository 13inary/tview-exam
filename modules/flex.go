package modules

import (
	"tview-exam/core"

	"github.com/rivo/tview"
)

func BuildFlex(app *core.App) *core.Flex {
	flex := core.NewFlex()

	left := core.NewBox("Left (1/2 x width of Top)", true)
	flex.AddItem(left.GetBox(), 0, 1, false)

	middleTop := core.NewBox("Top", true)
	middleMid := core.NewBox("Middle (3 x height of Top)", true)
	middleBottom := core.NewBox("Bottom (5 rows)", true)
	middle := core.NewFlex()
	middle.SetDirection(tview.FlexRow)
	middle.AddItem(middleTop.GetBox(), 0, 1, false)
	middle.AddItem(middleMid.GetBox(), 0, 3, false)
	middle.AddItem(middleBottom.GetBox(), 5, 1, false)
	flex.AddItem(middle.GetFlex(), 0, 2, false)

	right := core.NewBox("Right (20 cols)", true)
	flex.AddItem(right.GetBox(), 20, 1, false)
	return flex
}

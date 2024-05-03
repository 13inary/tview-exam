package core

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TreeNode struct {
	treeNode *tview.TreeNode
}

func NewTreeNode(text string) *TreeNode {
	t := &TreeNode{
		treeNode: tview.NewTreeNode(text),
	}
	return t
}
func (t *TreeNode) GetTreeNode() *tview.TreeNode {
	return t.treeNode
}
func (t *TreeNode) SetColor(color tcell.Color) {
	t.treeNode.SetColor(color)
}
func (t *TreeNode) SetSelectable(selectable bool) {
	t.treeNode.SetSelectable(selectable)
}
func (t *TreeNode) AddChild(node *TreeNode) {
	t.treeNode.AddChild(node.GetTreeNode())
}
func (t *TreeNode) GetChildren() []*tview.TreeNode {
	return t.treeNode.GetChildren()
}
func (t *TreeNode) ExistChildren() bool {
	return len(t.treeNode.GetChildren()) > 0
}
func (t *TreeNode) IsExpanded() bool {
	return t.treeNode.IsExpanded()
}
func (t *TreeNode) SetExpanded(expanded bool) {
	t.treeNode.SetExpanded(expanded)
}
func (t *TreeNode) SwitchExpand() {
	t.treeNode.SetExpanded(!t.treeNode.IsExpanded())
}
func (t *TreeNode) SetReference(reference interface{}) {
	t.treeNode.SetReference(reference)
}
func (t *TreeNode) GetReference() interface{} {
	return t.treeNode.GetReference()
}

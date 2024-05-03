package core

import (
	"github.com/rivo/tview"
)

type TreeView struct {
	treeView    *tview.TreeView
	root        *TreeNode
	currentNode *TreeNode
}

func NewTreeView(root *TreeNode, currentNode *TreeNode) *TreeView {
	t := &TreeView{
		treeView:    tview.NewTreeView(),
		root:        root,
		currentNode: currentNode,
	}
	t.treeView.SetRoot(t.root.GetTreeNode())
	t.treeView.SetCurrentNode(t.currentNode.GetTreeNode())
	return t
}
func (t *TreeView) GetTreeView() *tview.TreeView {
	return t.treeView
}
func (t *TreeView) SetSelectedFunc(handler func(node *tview.TreeNode)) {
	t.treeView.SetSelectedFunc(handler)
}

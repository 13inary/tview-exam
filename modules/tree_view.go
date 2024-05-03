package modules

import (
	"io/ioutil"
	"path/filepath"
	"tview-exam/core"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func BuildTreeView(app *core.App) *core.TreeView {
	root := core.NewTreeNode(".")
	root.SetColor(tcell.ColorRed)

	tree := core.NewTreeView(root, root)

	addNode := func(target *tview.TreeNode, path string) {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			node := core.NewTreeNode(file.Name())
			node.SetReference(filepath.Join(path, file.Name()))
			node.SetSelectable(file.IsDir())
			if file.IsDir() {
				node.SetColor(tcell.ColorGreen)
			}
			target.AddChild(node.GetTreeNode())
		}
	}

	addNode(root.GetTreeNode(), ".")

	// If a directory was selected, open it.
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return // Selecting the root node does nothing.
		}

		children := node.GetChildren()
		if len(children) == 0 {
			// Load and show files in this directory.
			path := reference.(string)
			addNode(node, path)
		} else {
			// Collapse if visible, expand if collapsed.
			node.SetExpanded(!node.IsExpanded())
		}
	})
	return tree
}

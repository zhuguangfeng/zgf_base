package main

import (
	"fmt"
)

type TreeNode struct {
	Id       string
	Name     string
	Children []*TreeNode
}

func addNode(parent *TreeNode, path []string) *TreeNode {
	if len(path) == 0 {
		return parent
	}
	nodeMap := make(map[string]*TreeNode)
	name := path[0]
	var child *TreeNode
	for _, c := range parent.Children {
		if c.Name == name {
			child = c
			break
		}
	}
	if child == nil {
		id := fmt.Sprintf("%d", len(parent.Children)+1)
		child = &TreeNode{
			Id:       id,
			Name:     name,
			Children: []*TreeNode{},
		}
		parent.Children = append(parent.Children, child)
		nodeMap[name] = child
	}
	return addNode(child, path[1:])
}
func genCascadeTree(req [][]string) *TreeNode {
	root := &TreeNode{}

	for _, group := range req {
		addNode(root, group)
	}

	return root
}

func main() {
	var sli = [][]string{
		{"骁龙", "骁龙S1", "APQ8064"},
		{"骁龙", "骁龙S1", "MSM8960T"},
		{"骁龙", "骁龙200", "骁龙212(MSM8909v2)"},
		{"骁龙", "骁龙200", "骁龙200(MSM8x12)"},
		{"三星(Exynos)", "Exynos 2000系列", "Exynos 2200"},
		{"三星(Exynos)", "Exynos 3000|4000|5000系列", "Exynos 3475"},
	}
	tree := genCascadeTree(sli)
	printTree(tree, 0)
}

func printTree(node *TreeNode, level int) {
	for i := 0; i < level; i++ {
		fmt.Printf("  ")
	}
	fmt.Printf("%s (%s)\n", node.Name, node.Id)
	for _, child := range node.Children {
		printTree(child, level+1)
	}
}

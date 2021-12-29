/**
 * File: tree_node.go
 * Author: Eddie Ho
 * Date: 2021-12-27
 * Project: BST
 * Purpose: Declare and implement TreeNode struct
 */

package tree_node

import (
	"fmt"
)

func max(x int, y int) int {
	if x > y {
		return x
	}

	return y
}

type TreeNode struct {
	item   int
	left   *TreeNode
	right  *TreeNode
	height int
}

func NewTreeNode(item int, left *TreeNode, right *TreeNode) *TreeNode {
	p := TreeNode{item: item, left: left, right: right, height: 0}
	p.UpdateHeight()
	return &p
}

func (t *TreeNode) Display() {
	fmt.Printf("|%+v|", t.item)
}

func (t *TreeNode) Height() int {
	leftHeight := -1
	if t.left != nil {
		leftHeight = t.left.height
	}

	rightHeight := -1
	if t.right != nil {
		rightHeight = t.right.height
	}

	return 1 + max(leftHeight, rightHeight)
}

func (t *TreeNode) BalanceFactor() int {
	leftHeight := -1
	if t.left != nil {
		leftHeight = t.left.height
	}

	rightHeight := -1
	if t.right != nil {
		rightHeight = t.right.height
	}

	return leftHeight - rightHeight
}

func (t *TreeNode) UpdateHeight() int {
	t.height = t.Height()
	return t.height
}

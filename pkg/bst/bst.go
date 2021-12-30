package bst

import (
	. "github.com/noodleslove/bst/pkg/tree_node"
)

type Bst struct {
	root *TreeNode
}

func NewBst() *Bst {
	p := Bst{root: nil}
	return &p
}

func NewBstFromList(sortedList []int) *Bst {
	p := Bst{root: FromSortedList(sortedList)}
	return &p
}

func (b *Bst) Insert(insertMe int) {
	TreeInsert(&b.root, insertMe)
}

func (b *Bst) Erase(target int) {
	TreeErase(&b.root, target)
}

func (b *Bst) Contains(target int) bool {
	return TreeSearch(b.root, target) != nil
}

func (b *Bst) ClearAll() {
	TreeClear(&b.root)
}

func (b *Bst) Empty() bool {
	return b.root == nil
}

func (b *Bst) InOrder() string {
	return InOrderString(b.root)
}

func (b *Bst) PreOrder() string {
	return PreOrderString(b.root)
}

func (b *Bst) PostOrder() string {
	return PostOrderString(b.root)
}

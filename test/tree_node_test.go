package test

import (
	"fmt"
	"math/rand"
	"testing"

	. "github.com/noodleslove/bst/pkg/tree_node"
)

const debug bool = true

func TestTreeNode(t *testing.T) {
	n := NewTreeNode(0, nil, nil)
	if n == nil || n.Height() != 0 {
		t.Errorf("NewTreeNode incorrect")
	}
}

func TestTreeSearch(t *testing.T) {
	var tree *TreeNode
	var res *TreeNode
	a := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		a[i] = i + 1
	}

	tree = FromSortedList(a)
	res = TreeSearch(tree, 0)
	if res != nil {
		t.Errorf("TreeSearch incorrect")
	}

	for i := 1; i <= 1000; i++ {
		res = TreeSearch(tree, i)
		if res == nil {
			t.Errorf("TreeSearch incorrect")
		}
	}

	res = TreeSearch(tree, 9999)
	if res != nil {
		t.Errorf("TreeSearch incorrect")
	}
}

func TestTreeInsert(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := []int{5, 2, 7, 1, 3, 6, 8, 9}

	r1 := FromSortedList(a)
	if debug {
		TreePrint(r1, 0)
		fmt.Println()
	}

	var r2 *TreeNode = nil
	for _, e := range b {
		TreeInsert(&r2, e)
	}
	if debug {
		TreePrint(r2, 0)
	}
}

func TestTreeErase(t *testing.T) {
	a, flag := make([]int, 1000), false
	for i := 0; i < 1000; i++ {
		a[i] = i + 1
	}

	r := FromSortedList(a)
	for i := 1000; i > 0; i-- {
		flag = TreeErase(&r, i)
		if flag == false {
			t.Errorf("TreeErase incorrect")
		}
	}

	if r != nil {
		t.Errorf("TreeErase incorrect: expect r = nil")
	}
}

func TestTreeCopy(t *testing.T) {
	a := make([]int, 10)
	for i := range a {
		a[i] = rand.Intn(100)
	}

	r1 := FromSortedList(a)
	r2 := TreeCopy(r1)

	if debug {
		TreePrint(r1, 0)
		fmt.Println()
		TreePrint(r2, 0)
	}
}

func TestTreeAdd(t *testing.T) {
	r1 := FromSortedList([]int{1, 3, 5})
	r2 := FromSortedList([]int{6, 8, 10})

	TreeAdd(&r1, r2)
	if debug {
		TreePrint(r1, 0)
	}
}

func TestTreeClear(t *testing.T) {
	a := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		a[i] = i + 1
	}

	r := FromSortedList(a)
	if r == nil {
		t.Errorf("FromSortedList incorrect")
	}

	TreeClear(&r)
	if r != nil {
		t.Errorf("TreeClear incorrect")
	}
}

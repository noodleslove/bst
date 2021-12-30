package tree_node

import (
	"fmt"
	"os"
)

// Postcondition: Return value is the tree node with target value, if it cannot
//      find any node with matching value then return NULL.
func TreeSearch(root *TreeNode, target int) *TreeNode {
	if root == nil { // Base case
		return root
	}

	if root.item == target { // Found target
		return root
	}

	if target < root.item { // Recursive step
		return TreeSearch(root.left, target) // Left sub-tree
	}

	return TreeSearch(root.right, target) // Right sub-tree
}

// Postcondition: Insert the given value to correct place in BST
func TreeInsert(root **TreeNode, insertMe int) {
	if *root == nil { // Base case
		*root = NewTreeNode(insertMe, nil, nil)
	} else if insertMe < (*root).item { // Move to left sub-tree
		TreeInsert(&(*root).left, insertMe)
	} else { // Move to right sub-tree
		TreeInsert(&(*root).right, insertMe)
	}

	(*root).UpdateHeight() // Update tree height
}

// Postcondition: Return value is a BST copy of the given BST
func TreeCopy(root *TreeNode) *TreeNode {
	if root == nil { // Base case
		return nil
	}

	// Copy BST node by node
	return NewTreeNode(root.item, TreeCopy(root.left), TreeCopy(root.right))
}

// Postcondition: Free tree node pointers and clear BST object
func TreeClear(root **TreeNode) {
	if *root == nil { // Base case
		return
	}

	TreeClear(&(*root).left)  // Clear left sub-tree first
	TreeClear(&(*root).right) // then clear right sub-tree
	*root = nil               // assign pointer to NULL
}

func HeightHelper(root **TreeNode) {
	if *root != nil {
		(*root).UpdateHeight()
	}
}

// Postcondition: Delete the given target from BST object. If cannot find
//      target then return false and do nothing to BST. Otherwise, delete
//      the target tree node, and adjust the BST object accordingly.
func TreeErase(root **TreeNode, target int) bool {
	if *root == nil { // Base case
		return false
	}

	defer HeightHelper(root) // Update tree height

	if target == (*root).item { // Case 4: Found the target tree node
		if (*root).left == nil { // Case 4.a: Root has no left sub-tree,
			*root = (*root).right // bypass root and connect the right child
		} else { // Case 4.b: Target has a left child: Replace target with left child's
			TreeRemoveMax(&(*root).left, &(*root).item) // rightmost child. Eliminate the
		} // rightmost child from left sub-tree.

		return true
	} else if target < (*root).item { // Recursive step: Search left sub-tree
		return TreeErase(&(*root).left, target)
	}

	return TreeErase(&(*root).right, target) // Recursive step: Search right sub-tree
}

// Postcondition: Replace target with its left child's rightmost child. Also
//      eliminate the rightmost child from left sub-tree
func ReplaceMax(root **TreeNode, target **TreeNode) {
	if (*root).right == nil { // Base case
		(*target).item = (*root).item // Replace target with max value
		(*root) = (*root).right
		return
	}

	ReplaceMax(&(*root).right, target) // Recursive step: Find max value

	(*root).UpdateHeight() // Update tree height afterward
}

// Postcondition: Erase rightmost node from the tree -> max_value
func TreeRemoveMax(root **TreeNode, maxVal *int) {
	if (*root).right == nil { // Base case
		*maxVal = (*root).item // Assign max value to destination
		(*root) = (*root).left // Connect its left sub-tree
		return
	}

	// Recursive step: Search right child
	TreeRemoveMax(&(*root).right, maxVal)

	(*root).UpdateHeight() // Update tree height afterward
}

// Postcondition: Insert all values in second BST object to first BST object.
func TreeAdd(dest **TreeNode, src *TreeNode) {
	if src == nil { // Base case
		return
	}

	TreeInsert(dest, src.item) // Insert current node to first BST
	// Recursive step
	TreeAdd(dest, src.left)  // Insert left sub-tree
	TreeAdd(dest, src.right) // Insert right sub-tree

	(*dest).UpdateHeight() // Update tree height
}

// Postcondition: Construct a BST object from a sorted list
func FromSortedList(arr []int) *TreeNode {
	size := len(arr)
	if size == 0 { // Base case
		return nil
	}
	// Find mid element and create node
	mid := size / 2
	root := NewTreeNode(arr[mid], nil, nil)
	left, right := arr[:mid], arr[mid+1:]
	// Recursive step: Repeat process for left and right sub-tree
	root.left = FromSortedList(left)
	root.right = FromSortedList(right)

	return root
}

// Postcondition: Print BST
func TreePrint(root *TreeNode, depth int) {
	if root != nil {
		TreePrint(root.right, depth+1)                  // Print right sub-tree
		fmt.Printf("%*s[%d]\n", 4*depth, "", root.item) // Print current root
		TreePrint(root.left, depth+1)                   // Print left sub-tree
	}
}

// =================================
// Traversal
// =================================

func InOrder(root *TreeNode, out *os.File) {
	if root == nil {
		return
	}

	InOrder(root.left, out)
	_, err := out.Write([]byte(fmt.Sprintf("%4d", root.item)))
	if err != nil {
		panic(err)
	}
	InOrder(root.right, out)
}

func PreOrder(root *TreeNode, out *os.File) {
	if root == nil {
		return
	}

	_, err := out.Write([]byte(fmt.Sprintf("%4d", root.item)))
	if err != nil {
		panic(err)
	}
	PreOrder(root.left, out)
	PreOrder(root.right, out)
}

func PostOrder(root *TreeNode, out *os.File) {
	if root == nil {
		return
	}

	PostOrder(root.left, out)
	PostOrder(root.right, out)
	_, err := out.Write([]byte(fmt.Sprintf("%4d", root.item)))
	if err != nil {
		panic(err)
	}
}

func InOrderString(root *TreeNode) string {
	if root == nil {
		return ""
	}

	return InOrderString(root.left) + fmt.Sprintf("[%d]", root.item) + InOrderString(root.right)
}

func PreOrderString(root *TreeNode) string {
	if root == nil {
		return ""
	}

	return fmt.Sprintf("[%d]", root.item) + PreOrderString(root.left) + PreOrderString(root.right)
}

func PostOrderString(root *TreeNode) string {
	if root == nil {
		return ""
	}

	return PostOrderString(root.left) + PostOrderString(root.right) + fmt.Sprintf("[%d]", root.item)
}

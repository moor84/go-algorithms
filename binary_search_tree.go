package algorithms

import "errors"

// BinarySearchTreeNode node
type BinarySearchTreeNode struct {
	key   int
	value string
	left  *BinarySearchTreeNode
	right *BinarySearchTreeNode
}

// BinarySearchTree implementation
type BinarySearchTree struct {
	root *BinarySearchTreeNode
}

// Find a value by key
func (bst *BinarySearchTree) Find(key int) (string, error) {
	node := bst.findNode(bst.root, key)
	if node != nil {
		return node.value, nil
	}
	return "", errors.New("Not found")
}

func (bst *BinarySearchTree) findNode(node *BinarySearchTreeNode, key int) *BinarySearchTreeNode {
	if node == nil {
		return nil
	}
	if node.key == key {
		return node
	}
	if key < node.key {
		return bst.findNode(node.left, key)
	}
	if key > node.key {
		return bst.findNode(node.right, key)
	}
	return nil
}

// Insert a value with a key
func (bst *BinarySearchTree) Insert(key int, value string) {
	bst.root = bst.insertNode(bst.root, key, value)
}

func (bst *BinarySearchTree) insertNode(node *BinarySearchTreeNode, key int, value string) *BinarySearchTreeNode {
	if node == nil {
		return &BinarySearchTreeNode{key: key, value: value}
	}

	if key < node.key {
		node.left = bst.insertNode(node.left, key, value)
	} else if key > node.key {
		node.right = bst.insertNode(node.right, key, value)
	}

	return node
}

// Delete given key
func (bst *BinarySearchTree) Delete(key int) {
	bst.root = bst.deleteNode(bst.root, key)
}

func (bst *BinarySearchTree) findMin(node *BinarySearchTreeNode) *BinarySearchTreeNode {
	if node.left == nil {
		return node
	}
	return bst.findMin(node.left)
}

func (bst *BinarySearchTree) deleteNode(node *BinarySearchTreeNode, key int) *BinarySearchTreeNode {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = bst.deleteNode(node.left, key)
		return node
	}
	if key > node.key {
		node.right = bst.deleteNode(node.right, key)
		return node
	}
	// key == node.key:
	// No children
	if node.left == nil && node.right == nil {
		return nil
	}
	// One child
	if node.left == nil {
		return node.right
	} else if node.right == nil {
		return node.left
	}
	// Two children
	minRight := bst.findMin(node.right)
	node.key = minRight.key
	node.value = minRight.value
	node.right = bst.deleteNode(node.right, node.key)
	return node
}

func (bst *BinarySearchTree) inOrderNodes(node *BinarySearchTreeNode, nodes *[]*BinarySearchTreeNode) {
	if node == nil {
		return
	}
	bst.inOrderNodes(node.left, nodes)
	*nodes = append(*nodes, node)
	bst.inOrderNodes(node.right, nodes)
}

func (bst *BinarySearchTree) preOrderNodes(node *BinarySearchTreeNode, nodes *[]*BinarySearchTreeNode) {
	if node == nil {
		return
	}
	*nodes = append(*nodes, node)
	bst.preOrderNodes(node.left, nodes)
	bst.preOrderNodes(node.right, nodes)
}

func (bst *BinarySearchTree) postOrderNodes(node *BinarySearchTreeNode, nodes *[]*BinarySearchTreeNode) {
	if node == nil {
		return
	}
	bst.postOrderNodes(node.left, nodes)
	bst.postOrderNodes(node.right, nodes)
	*nodes = append(*nodes, node)
}

func (bst *BinarySearchTree) traverseKeys(f func(*BinarySearchTreeNode, *[]*BinarySearchTreeNode)) []int {
	nodes := make([]*BinarySearchTreeNode, 0)
	keys := make([]int, 0)
	f(bst.root, &nodes)
	for _, node := range nodes {
		keys = append(keys, node.key)
	}
	return keys
}

// InOrder traversal
func (bst *BinarySearchTree) InOrder() []int {
	return bst.traverseKeys(bst.inOrderNodes)
}

// PreOrder traversal
func (bst *BinarySearchTree) PreOrder() []int {
	return bst.traverseKeys(bst.preOrderNodes)
}

// PostOrder traversal
func (bst *BinarySearchTree) PostOrder() []int {
	return bst.traverseKeys(bst.postOrderNodes)
}

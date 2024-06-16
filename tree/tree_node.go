// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tree

// Node of tree
type Node[T any] struct {
	element T
	left    *Node[T]
	right   *Node[T]
	parent  *Node[T]
}

// IsLeaf current node is leaf node
func (n *Node[T]) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

// HashTwoChildren current node has two children
func (n *Node[T]) HashTwoChildren() bool {
	return n.left != nil && n.right != nil
}

// IsLeftChild current node is left child node
func (n *Node[T]) IsLeftChild() bool {
	return n.parent != nil && n == n.parent.left
}

// IsRightChild current node is right child node
func (n *Node[T]) IsRightChild() bool {
	return n.parent != nil && n == n.parent.right
}

// sibling return current node sibling node
func (n *Node[T]) sibling() *Node[T] {
	if n.IsLeftChild() {
		return n.parent.right
	}
	if n.IsRightChild() {
		return n.parent.left
	}
	// has no sibling
	return nil
}

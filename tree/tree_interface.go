// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tree

import (
	"github.com/fengyuan-liang/GoKit/collection"
)

// ITree represents the interface for a binary search tree
type ITree[E comparable] interface {
	// Size returns the number of tree nodes
	Size() uint64
	// IsEmpty determines if the tree is empty
	IsEmpty() bool
	// Clear clears the tree
	Clear()
	// Height returns the height of the tree
	Height() uint64
	// HeightWithNode returns the height of the specified node
	HeightWithNode(node *Node[E]) uint64
	// HeightByLevelOrder returns the height of the tree using a non-recursive level-order traversal
	HeightByLevelOrder() uint64
	// Compare compare two element
	Compare(e1 E, e2 E) int
}

// BaseTree is a base struct for the binary search tree
type BaseTree[E any] struct {
	// comparator is used to compare elements during search
	comparator collection.CompareFunc[E]
	// size represents the current number of nodes in the tree
	size uint64
	// root is the root node of the tree
	root *Node[E]
}

// Visit is a function type used during tree traversal
// It returns true to stop traversing and false to continue traversing
type Visit[E comparable] func(element E) bool

// Visitor represents a visitor during tree traversal
type Visitor[E comparable] struct {
	// Stop indicates whether to stop traversing
	Stop bool
	// Visit is the visit function
	Visit Visit[E]
}

// ITraversal represents the interface for tree traversal
type ITraversal[E comparable] interface {
	// preorderTraversal performs a PreOrder traversal of the tree
	preorderTraversal(visitor *Visitor[E])
	// inorderTraversal performs an InOrder traversal of the tree
	inorderTraversal(visitor *Visitor[E])
	// postorderTraversal performs a PostOrder traversal of the tree
	postorderTraversal(visitor *Visitor[E])
	// leverOrderTraversal performs a LevelOrder traversal of the tree
	leverOrderTraversal(visitor *Visitor[E])
}

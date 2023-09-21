// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tree

import "GoKit/collection"

type ITree[E comparable] interface {
	// Size the number of tree nodes
	Size() uint64
	// IsEmpty determine if the tree is empty
	IsEmpty() bool
	// Clear tree
	Clear()
	// Height returns the height of the tree
	Height() uint64
	// HeightWithNode returns the height of the specified node
	HeightWithNode(node *Node[E]) uint64
	// HeightByLevelOrder Non-recursive return of tree height
	HeightByLevelOrder() uint64
}

type BaseTree[E comparable] struct {
	// comparator compare with search
	comparator collection.Compare[E]
	// size current number of nodes
	size uint64
	// root the root node of the tree
	root *Node[E]
}

// Visit returns true to stop traversing
//
// Returns false to continue traversing without stopping
type Visit[E comparable] func(element E) bool

type Visitor[E comparable] struct {
	// Stop tag stop traversing
	Stop bool
	// Visit
	Visit Visit[E]
}

// ITraversal traversal the tree
// @see Visitor
type ITraversal[E comparable] interface {
	// preorderTraversal PreOrder the tree
	preorderTraversal(visitor *Visitor[E])
	// inorderTraversal
	inorderTraversal(visitor *Visitor[E])
	// postorderTraversal
	postorderTraversal(visitor *Visitor[E])
	// leverOrderTraversal
	leverOrderTraversal(visitor *Visitor[E])
}

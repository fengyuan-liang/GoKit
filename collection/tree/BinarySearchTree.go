// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tree

import "math"

func NewBstTree[E comparable]() ITree[E] {
	return &BinarySearchTree[E]{}
}

type BinarySearchTree[E comparable] struct {
	BaseTree[E]
}

func (b *BinarySearchTree[E]) Size() uint64 {
	return b.size
}

func (b *BinarySearchTree[E]) IsEmpty() bool {
	return b.size == 0
}

func (b *BinarySearchTree[E]) Clear() {
	b.root = nil
	b.size = 0
}

func (b *BinarySearchTree[E]) Height() uint64 {
	return b.HeightWithNode(b.root)
}

func (b *BinarySearchTree[E]) HeightWithNode(node *Node[E]) uint64 {
	if node == nil {
		return 0
	}
	return uint64(1 + math.Max(float64(b.HeightWithNode(node.left)), float64(b.HeightWithNode(node.right))))
}

func (b *BinarySearchTree[E]) HeightByLevelOrder() uint64 {
	//TODO implement me
	panic("implement me")
}

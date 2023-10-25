// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package lists

import (
	"fmt"
	"github.com/fengyuan-liang/GoKit/collection"
	"github.com/fengyuan-liang/GoKit/collection/stream"
)

type BaseList[E any] struct {
	// list size
	size int
}

// Size returns the number of elements in the ArrayList.
func (b *BaseList[E]) Size() int {
	return b.size
}

// IsEmpty checks if the ArrayList is empty.
func (b *BaseList[E]) IsEmpty() bool {
	return b.size == 0
}

func (b *BaseList[E]) rangeCheckForAdd(index int) {
	// Check if the index is out of bounds for adding an element.
	// If the index is less than 0 or greater than the size, it is considered illegal.
	if index < 0 || index > b.size {
		panic(fmt.Sprintf("illegal index, index=%d, actual size=%d", index, b.size))
	}
}

func (b *BaseList[E]) rangeCheck(index int) {
	// Check if the index is out of bounds for adding an element.
	// If the index is less than 0 or greater than the size, it is considered illegal.
	if index < 0 || index > b.size-1 {
		panic(fmt.Sprintf("illegal index, index=%d, actual size=%d", index, b.size))
	}
}

// ================= abstract ====================

func (b *BaseList[E]) Stream() stream.IStream[E, E] {
	//TODO implement me
	panic("implement me")
}

func (b *BaseList[E]) Contains(element E) bool {
	//TODO implement me
	panic("implement me")
}

func (b *BaseList[E]) Add(element E) {
	//TODO implement me
	panic("implement me")
}

func (b *BaseList[E]) AddAll(elements []E) {
	//TODO implement me
	panic("implement me")
}

func (b *BaseList[E]) AddAtIndex(index int, element E) {
	//TODO implement me
	panic("implement me")
}

func (b *BaseList[E]) Get(index int) E {
	//TODO implement me
	panic("implement me")
}

func (b *BaseList[E]) Set(index int, element E) E {
	//TODO implement me
	panic("implement me")
}

func (b *BaseList[E]) Remove(index int) (oldElement E) {
	//TODO implement me
	panic("implement me")
}

func (b *BaseList[E]) IndexOf(element E) int {
	//TODO implement me
	panic("implement me")
}

func (b *BaseList[E]) ToSlice() []E {
	//TODO implement me
	panic("implement me")
}

func (b *BaseList[E]) Sort(compareFunc ...collection.CompareFunc[E]) {
	//TODO implement me
	panic("implement me")
}

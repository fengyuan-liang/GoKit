// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package lists

import (
	"fmt"
	"github.com/fengyuan-liang/GoKit/collection"
	"github.com/fengyuan-liang/GoKit/collection/stream"
	"strings"
)

// NewArrayList creates a new ArrayList with the default capacity.
func NewArrayList[E any]() IList[E] {
	return NewArrayListWithCapacity[E](DefaultCapacity)
}

// NewArrayListWithCapacity creates a new ArrayList with the specified capacity.
func NewArrayListWithCapacity[E any](capacity int) IList[E] {
	return &ArrayList[E]{
		elements: make([]E, capacity),
	}
}

// ArrayList represents a dynamic array implementation of the IList interface.
type ArrayList[E any] struct {
	BaseList[E]
	// elements
	elements []E
}

// Contains checks if the ArrayList contains the specified element.
func (a *ArrayList[E]) Contains(element E) bool {
	return a.IndexOf(element) != ElementNotFound
}

// Add appends the element to the end of the ArrayList.
func (a *ArrayList[E]) Add(element E) {
	a.AddAtIndex(a.size, element)
}

func (a *ArrayList[E]) AddAll(elements []E) {
	a.elements = append(a.elements, elements...)
	a.size += len(elements)
}

// AddAtIndex inserts the element at the specified index in the ArrayList.
func (a *ArrayList[E]) AddAtIndex(index int, element E) {
	a.rangeCheckForAdd(index)
	a.ensureCapacity(a.size + 1)

	// Shift elements to the right to make space for the new element.
	for i := a.size; i > index; i-- {
		a.elements[i] = a.elements[i-1]
	}

	a.elements[index] = element
	a.size++
}

// Get returns the element at the specified index in the ArrayList.
func (a *ArrayList[E]) Get(index int) E {
	a.rangeCheck(index)
	return a.elements[index]
}

// Set updates the element at the specified index in the ArrayList and returns the previous element.
func (a *ArrayList[E]) Set(index int, element E) E {
	a.rangeCheck(index)
	oldElement := a.elements[index]
	a.elements[index] = element
	return oldElement
}

// Remove removes the element at the specified index in the ArrayList and returns the removed element.
func (a *ArrayList[E]) Remove(index int) (oldElement E) {
	a.rangeCheck(index)
	oldElement = a.elements[index]

	// Shift elements to the left to fill the gap.
	for i := index; i < a.Size(); i++ {
		a.elements[i] = a.elements[i+1]
	}

	a.size--
	return
}

// Clear removes all elements from the ArrayList.
func (a *ArrayList[E]) Clear() {
	a.elements = make([]E, DefaultCapacity)
	a.size = 0
}

// IndexOf returns the index of the first occurrence of the specified element in the ArrayList.
// If the element is not found, it returns ElementNotFound.
func (a *ArrayList[E]) IndexOf(element E) int {
	if (any)(element) == nil {
		panic(fmt.Sprint("Element cannot be nil"))
	}

	for i := 0; i < a.size; i++ {
		if collection.Equal(a.elements[i], element) {
			return i
		}
	}

	return ElementNotFound
}

// ensureCapacity Ensure sufficient array capacity
func (a *ArrayList[E]) ensureCapacity(capacity int) {
	oldCapacity := a.size
	if capacity < oldCapacity {
		return
	}
	// Expand the capacity by 1.5 times
	newCapacity := oldCapacity + (oldCapacity >> 1)
	a.elements = append(a.elements, make([]E, newCapacity-oldCapacity)...)
}

func (a *ArrayList[E]) ToString() string {
	var builder strings.Builder
	builder.WriteString("[")
	for i := 0; i < a.size-1; i++ {
		builder.WriteString(fmt.Sprintf("%v,", a.elements[i]))
	}
	builder.WriteString(fmt.Sprintf("%v]", a.elements[a.size-1]))
	return builder.String()
}

func (a *ArrayList[E]) ToSlice() []E {
	return a.elements[:a.Size()]
}

func (a *ArrayList[E]) Stream() stream.IStream[E, E] {
	return stream.Of[E, E](a.elements)
}

func (a *ArrayList[E]) Sort(compareFunc ...collection.CompareFunc[E]) {
	collection.Sort(a.ToSlice(), compareFunc...)
}

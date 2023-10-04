// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package lists

import "GoKit/collection"

const (
	DefaultCapacity = 10
	ElementNotFound = -1
)

type IList[E any] interface {
	collection.ToString
	// Size Returns the number of elements in the list.
	Size() int
	// IsEmpty Checks if the list is empty.
	IsEmpty() bool
	// Contains Checks if the list contains a specific element.
	Contains(element E) bool
	// Add Adds an element to the end of the list.
	Add(element E)
	// AddAtIndex Inserts an element at the specified index.
	AddAtIndex(index int, element E)
	// Get Returns the element at the specified index.
	Get(index int) E
	// Set Sets the element at the specified index and returns the previous element.
	Set(index int, element E) E
	// Remove Delete values on the specified index
	Remove(index int) (oldElement E)
	// IndexOf Returns the index of the first occurrence of the specified element in the list.
	IndexOf(element E) int
	// Clear Removes all elements from the list.
	Clear()
}

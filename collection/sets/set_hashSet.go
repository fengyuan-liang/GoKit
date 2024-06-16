// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sets

import (
	"fmt"
	"strings"
)

// HashSet holds elements in go's native map
type HashSet[T comparable] struct {
	items map[T]struct{}
}

// EmptyStruct zero memory
var EmptyStruct = struct{}{}

func NewHashSet[T comparable]() ISet[T] {
	return &HashSet[T]{items: make(map[T]struct{})}
}

func (set *HashSet[T]) Add(elements ...T) {
	for _, item := range elements {
		set.items[item] = EmptyStruct
	}
}

func (set *HashSet[T]) Remove(elements ...T) {
	for _, item := range elements {
		delete(set.items, item)
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *HashSet[T]) Contains(elements ...T) bool {
	for _, item := range elements {
		if _, contains := set.items[item]; !contains {
			return false
		}
	}
	return true
}

// IsEmpty returns true if set does not contain any elements.
func (set *HashSet[T]) IsEmpty() bool {
	return set.Size() == 0
}

func (set *HashSet[T]) Size() int {
	return len(set.items)
}

func (set *HashSet[T]) Clear() {
	set.items = make(map[T]struct{})
}

func (set *HashSet[T]) Values() []T {
	values := make([]T, set.Size())
	count := 0
	for item := range set.items {
		values[count] = item
		count++
	}
	return values
}

func (set *HashSet[T]) String() string {
	str := "HashSet\n"
	var items []string
	for k := range set.items {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	return str
}

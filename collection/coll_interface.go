// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package collection

// CompareFunc is a function type used to define a comparison function in GoKit collection.
// It imposes a total ordering on a collection of objects.
// Comparators can be passed to sorting methods to provide precise control over the sort order.
// Comparators can also be used to control the order of certain data structures,
// such as sorted sets or sorted maps, or to provide an ordering for collections of objects
// that don't have a natural ordering.
//
// The comparison function takes two arguments, 'o1' and 'o2', of type 'T' (a comparable type),
// and returns an integer value that represents the comparison result:
// - A negative value if 'o1' is less than 'o2'.
// - Zero if 'o1' is equal to 'o2'.
// - A positive value if 'o1' is greater than 'o2'.
//
//	Example usage:
//	func CompareInt(a, b int) int {
//		if a < b {
//			return -1
//		} else if a > b {
//			return 1
//		}
//			return 0
//		}
//
// var compareFn Compare[int] = CompareInt
//
// Using the comparison function in a sort operation:
// someInts := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
// sort.Slice(someInts, func(i, j int) bool {
// return compareFn(someInts[i], someInts[j]) < 0
// })
type CompareFunc[T any] func(o1 T, o2 T) int

// Comparable is a function type used to define a comparison interface in GoKit collection.
type Comparable[E any] interface {
	// CompareTo define compare rule
	CompareTo(element E) int
}

// Iterable represents a collection that can be iterated over.
type Iterable interface {
	// Iterator returns an iterator for the collection.
	Iterator() Iterator
}

// Iterator represents an iterator over a collection.
type Iterator interface {
	// Next returns true if there are more elements to iterate, or false otherwise.
	Next() bool

	// Value returns the current value of the iterator.
	Value() interface{}
}

type ToString interface {
	ToString() string
}

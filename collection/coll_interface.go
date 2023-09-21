// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package collection

// Comparator A comparison function, which imposes a total ordering on some collection of objects.
// Comparators can be passed to a sort method (such as collections.sort or Arrays.sort) to
// allow precise control over the sort order. Comparators can also be used to control the
// order of certain data structures (such as sorted sets or sorted maps), or to provide
// an ordering for collections of objects that don't have a natural ordering.
type Comparator interface {
	// Compares its two arguments for order. Returns a negative integer, zero, or a
	// positive integer as the first argument is less than, equal to, or greater than the second.
	compare(o1 any, o2 any) int
}

type Compare[T comparable] func(o1 T, o2 T) int

// Comparable This interface imposes a total ordering on the objects of each
// class that implements it. This ordering is referred to as the class's
// natural ordering, and the class's compareTo method is referred to as
// its natural comparison method.
type Comparable interface {
	// Compares this object with the specified object for order.
	// Returns a negative integer, zero, or a positive integer
	// as this object is less than, equal to, or greater than
	// the specified object.
	compareTo(o any) int
}

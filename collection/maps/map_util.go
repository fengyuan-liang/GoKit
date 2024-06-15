// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package maps

import (
	"fmt"
	"math"
)

// checkNonnegative checks if the value is non-negative
func checkNonnegative(value int, name string) {
	if value < 0 {
		panic(fmt.Sprintf("%s cannot be negative: %d", name, value))
	}
}

// capacity calculates a sufficient initial capacity to prevent resizing
// as long as the map does not grow larger than expectedSize and the load factor is ≥ 0.75
func capacity(expectedSize int) int {
	if expectedSize < 3 {
		checkNonnegative(expectedSize, "expectedSize")
		return expectedSize + 1
	}
	if expectedSize < (1 << 30) {
		// This calculation ensures that the load factor does not exceed 0.75
		// when the map grows to the expected size.
		return int(float64(expectedSize)/0.75 + 1)
	}
	return math.MaxInt32 // any large value
}

// NewHashMap creates a new instance of a basic hash map.
// It returns an implementation of the IMap interface using NewEnhancedMap.
func NewHashMap[K comparable, V any]() IMap[K, V] {
	return NewEnhancedMap[K, V]()
}

// NewHashMapWithExpectedSize creates a new instance of a hash map with an expected size.
// It returns an implementation of the IMap interface using NewEnhancedMapWithExpectedSize.
func NewHashMapWithExpectedSize[K comparable, V any](expectedSize int) IMap[K, V] {
	return NewEnhancedMapWithExpectedSize[K, V](expectedSize)
}

// NewConcurrentHashMap creates a new thread-safe hash map.
// It returns an implementation of the IMap interface using a SynchronizedMap that wraps a NewEnhancedMap.
func NewConcurrentHashMap[K comparable, V any]() IMap[K, V] {
	return NewSynchronizedMap[K, V](NewEnhancedMap[K, V]())
}

// NewConcurrentHashMapWithExpectedSize creates a new thread-safe hash map with an expected size.
// It returns an implementation of the IMap interface using a SynchronizedMap that wraps a NewEnhancedMapWithExpectedSize.
func NewConcurrentHashMapWithExpectedSize[K comparable, V any](expectedSize int) IMap[K, V] {
	return NewSynchronizedMap[K, V](NewEnhancedMapWithExpectedSize[K, V](expectedSize))
}

// NewConcurrentLinkedHashMap creates a new thread-safe linked hash map.
// It returns an implementation of the IMap interface using a SynchronizedMap that wraps a NewLinkedHashMap.
func NewConcurrentLinkedHashMap[K comparable, V any]() IMap[K, V] {
	return NewSynchronizedMap[K, V](NewLinkedHashMap[K, V]())
}

// NewConcurrentLinkedHashMapWithExpectedSize creates a new thread-safe linked hash map with an expected size.
// It returns an implementation of the IMap interface using a SynchronizedMap that wraps a NewLinkedHashMapWithExpectedSize.
func NewConcurrentLinkedHashMapWithExpectedSize[K comparable, V any](expectedSize int) IMap[K, V] {
	return NewSynchronizedMap[K, V](NewLinkedHashMapWithExpectedSize[K, V](expectedSize))
}

// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package maps

// IMap represents a generic map interface with keys of type K and values of type V.
// @see TreeMap
// @see EnhancedMap enhance golang map
// @see LinkedHashMap
// @since v1.0.0
// @author eureka
type IMap[K comparable, V any] interface {
	// Put inserts a key-value pair into the map.
	Put(k K, v V)

	PutAll(pairs []*Pair[K, V])

	// Get retrieves the value associated with the specified key from the map.
	// It returns the value and true if the key is found; otherwise, it returns the zero value of V and false.
	Get(k K) (v V, ok bool)

	// MustGet retrieves the value associated with the specified key from the map.
	// It returns the value if the key is found; otherwise, it panics.
	MustGet(k K) (v V)

	// GetOrDefault retrieves the value associated with the specified key from the map.
	// It returns the value if the key is found; otherwise, it returns the defaultValue.
	GetOrDefault(k K, defaultValue V) (v V)

	// Remove removes the key-value pair associated with the specified key from the map.
	Remove(k K)

	// KeySet returns a slice containing all the keys in the map.
	KeySet() []K

	// Values returns a slice containing all the values in the map.
	Values() []V

	// Size returns the number of key-value pairs in the map.
	Size() int

	IsEmpty() bool

	// ForEach applies the specified function to each key-value pair in the map.
	ForEach(f func(k K, v V))

	// ContainsKey checks if the map contains the specified key.
	// It returns true if the key is found; otherwise, it returns false.
	ContainsKey(k K) bool

	// ContainsValue checks if the map contains the specified value.
	// It returns true if the value is found; otherwise, it returns false.
	ContainsValue(value V) bool

	// Clear removes all elements from the map.
	Clear()
}

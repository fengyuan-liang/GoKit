// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package maps

import (
	"reflect"
)

// NewEnhancedMap creates and returns a new instance of the EnhancedMap.
func NewEnhancedMap[K comparable, V any]() IMap[K, V] {
	return &EnhancedMap[K, V]{
		baseMap[K, V]{m: map[K]V{}},
	}
}

// EnhancedMap is a map implementation that satisfies the IMap interface.
type EnhancedMap[K comparable, V any] struct {
	baseMap[K, V]
}

// Put inserts a key-value pair into the map.
func (m *EnhancedMap[K, V]) Put(k K, v V) {
	m.m[k] = v
}

// PutAll put by raw map
func (m *EnhancedMap[K, V]) PutAll(pairs []*Pair[K, V]) {
	for _, p := range pairs {
		m.Put(p.Key, p.Value)
	}
}

// Get retrieves the value associated with the specified key from the map.
// It returns the value and true if the key is found; otherwise, it returns the zero value of V and false.
func (m *EnhancedMap[K, V]) Get(k K) (v V, ok bool) {
	v, ok = m.m[k]
	return v, ok
}

// MustGet retrieves the value associated with the specified key from the map.
// It returns the value if the key is found; otherwise, it panics.
func (m *EnhancedMap[K, V]) MustGet(k K) (v V) {
	v, _ = m.m[k]
	return v
}

// GetOrDefault retrieves the value associated with the specified key from the map.
// It returns the value if the key is found; otherwise, it returns the defaultValue.
func (m *EnhancedMap[K, V]) GetOrDefault(k K, defaultValue V) (v V) {
	if got, ok := m.m[k]; ok {
		return got
	}
	return defaultValue
}

// Remove removes the key-value pair associated with the specified key from the map.
func (m *EnhancedMap[K, V]) Remove(k K) {
	delete(m.m, k)
}

// KeySet returns a slice containing all the keys in the map.
func (m *EnhancedMap[K, V]) KeySet() []K {
	keySet := make([]K, 0)
	m.ForEach(func(k K, v V) {
		keySet = append(keySet, k)
	})
	return keySet
}

// Values returns a slice containing all the values in the map.
func (m *EnhancedMap[K, V]) Values() []V {
	valueSet := make([]V, 0)
	m.ForEach(func(k K, v V) {
		valueSet = append(valueSet, v)
	})
	return valueSet
}

// Size returns the number of key-value pairs in the map.
func (m *EnhancedMap[K, V]) Size() int {
	return len(m.m)
}

// ForEach applies the specified function to each key-value pair in the map.
func (m *EnhancedMap[K, V]) ForEach(f func(K, V)) {
	for k, v := range m.m {
		f(k, v)
	}
}

// ContainsKey checks if the map contains the specified key.
// It returns true if the key is found; otherwise, it returns false.
func (m *EnhancedMap[K, V]) ContainsKey(k K) bool {
	_, ok := m.m[k]
	return ok
}

// ContainsValue checks if the map contains the specified value.
// It returns true if the value is found; otherwise, it returns false.
func (m *EnhancedMap[K, V]) ContainsValue(value V) bool {
	for _, v := range m.m {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

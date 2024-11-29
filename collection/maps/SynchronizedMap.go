// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package maps

import "sync"

// SynchronizedMap is a thread-safe map
type SynchronizedMap[K comparable, V any] struct {
	m  IMap[K, V] // raw map
	mu *sync.RWMutex
}

// NewSynchronizedMap creates a new SynchronizedMap
func NewSynchronizedMap[K comparable, V any](rawMap IMap[K, V]) IMap[K, V] {
	return &SynchronizedMap[K, V]{
		m:  rawMap,
		mu: new(sync.RWMutex),
	}
}

// Put adds a key-value pair to the map
func (s *SynchronizedMap[K, V]) Put(k K, v V) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m.Put(k, v)
}

// PutAll adds multiple key-value pairs to the map
func (s *SynchronizedMap[K, V]) PutAll(pairs []*Pair[K, V]) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m.PutAll(pairs)
}

// Get retrieves the value associated with the key
func (s *SynchronizedMap[K, V]) Get(k K) (v V, ok bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.m.Get(k)
}

// MustGet retrieves the value associated with the key, panics if not found
func (s *SynchronizedMap[K, V]) MustGet(k K) (v V) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.m.MustGet(k)
}

// GetOrDefault retrieves the value associated with the key, or returns defaultValue if not found
func (s *SynchronizedMap[K, V]) GetOrDefault(k K, defaultValue V) (v V) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.m.GetOrDefault(k, defaultValue)
}

// Remove removes the key-value pair from the map
func (s *SynchronizedMap[K, V]) Remove(k K) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m.Remove(k)
}

// KeySet returns a slice of all keys in the map
func (s *SynchronizedMap[K, V]) KeySet() []K {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.m.KeySet()
}

// Values return a slice of all values in the map
func (s *SynchronizedMap[K, V]) Values() []V {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.m.Values()
}

// Size returns the number of key-value pairs in the map
func (s *SynchronizedMap[K, V]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.m.Size()
}

// IsEmpty checks if the map is empty
func (s *SynchronizedMap[K, V]) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.m.IsEmpty()
}

// ForEach applies a function to each key-value pair in the map
func (s *SynchronizedMap[K, V]) ForEach(f func(k K, v V)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m.ForEach(f)
}

// ContainsKey checks if the map contains the given key
func (s *SynchronizedMap[K, V]) ContainsKey(k K) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.m.ContainsKey(k)
}

// ContainsValue checks if the map contains the given value
func (s *SynchronizedMap[K, V]) ContainsValue(value V) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.m.ContainsValue(value)
}

// Clear removes all key-value pairs from the map
func (s *SynchronizedMap[K, V]) Clear() {
	s.mu.Lock() // Use write lock instead of read lock
	defer s.mu.Unlock()
	s.m.Clear()
}

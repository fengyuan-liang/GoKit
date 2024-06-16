// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package maps

func NewLinkedHashMap[K comparable, V any]() IMap[K, V] {
	return &LinkedHashMap[K, V]{
		keys: make([]K, 0),
		baseMap: baseMap[K, V]{
			m: make(map[K]V),
		},
	}
}

func NewLinkedHashMapWithExpectedSize[K comparable, V any](expectedSize int) IMap[K, V] {
	return &LinkedHashMap[K, V]{
		keys: make([]K, 0),
		baseMap: baseMap[K, V]{
			m: make(map[K]V, capacity(expectedSize)),
		},
	}
}

type LinkedHashMap[K comparable, V any] struct {
	baseMap[K, V]
	keys []K
}

func (m *LinkedHashMap[K, V]) Put(k K, v V) {
	if _, ok := m.m[k]; !ok {
		m.keys = append(m.keys, k)
	}
	m.m[k] = v
}

func (m *LinkedHashMap[K, V]) PutAll(pairs []*Pair[K, V]) {
	for _, p := range pairs {
		m.Put(p.Key, p.Value)
	}
}

func (m *LinkedHashMap[K, V]) GetOrDefault(k K, defaultValue V) (v V) {
	if got, ok := m.m[k]; ok {
		return got
	}
	return defaultValue
}

// Remove o(n)
func (m *LinkedHashMap[K, V]) Remove(k K) {
	delete(m.m, k)
	for i, key := range m.keys {
		if key == k {
			// This method of removal is relatively efficient in terms of performance.
			// Since the underlying implementation of slices is based on arrays,
			// using slice operations allows direct modification of the underlying array's pointers
			// and length without the need for moving the entire block of memory. Therefore,
			// the time complexity of this removal operation is O(1), which means it has constant time complexity.
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
			break
		}
	}
}

func (m *LinkedHashMap[K, V]) KeySet() []K {
	return m.keys
}

func (m *LinkedHashMap[K, V]) Values() []V {
	values := make([]V, 0, len(m.keys))
	for _, k := range m.keys {
		values = append(values, m.m[k])
	}
	return values
}

func (m *LinkedHashMap[K, V]) ForEach(f func(K, V)) {
	for _, k := range m.keys {
		f(k, m.m[k])
	}
}

func (m *LinkedHashMap[K, V]) Clear() {
	m.m = make(map[K]V)
	m.keys = make([]K, 0)
}

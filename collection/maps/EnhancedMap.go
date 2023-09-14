// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package maps

import "reflect"

func NewEnhancedMap[K comparable, V any]() IMap[K, V] {
	return &EnhancedMap[K, V]{
		m: map[K]V{},
	}
}

type EnhancedMap[K comparable, V any] struct {
	m map[K]V
}

func (m *EnhancedMap[K, V]) Put(k K, v V) {
	m.m[k] = v
}

func (m *EnhancedMap[K, V]) Get(k K) (v V, ok bool) {
	v, ok = m.m[k]
	return v, ok
}

func (m *EnhancedMap[K, V]) GetOrDefault(k K, defaultValue V) (v V) {
	if got, ok := m.m[k]; ok {
		return got
	}
	return defaultValue
}

func (m *EnhancedMap[K, V]) Remove(k K) {
	delete(m.m, k)
}

func (m *EnhancedMap[K, V]) KeySet() []K {
	keySet := make([]K, 0)
	m.ForEach(func(k K, v V) {
		keySet = append(keySet, k)
	})
	return keySet
}

func (m *EnhancedMap[K, V]) Values() []V {
	valueSet := make([]V, 0)
	m.ForEach(func(k K, v V) {
		valueSet = append(valueSet, v)
	})
	return valueSet
}

func (m *EnhancedMap[K, V]) Size() int {
	return len(m.m)
}

func (m *EnhancedMap[K, V]) ForEach(f func(K, V)) {
	for k, v := range m.m {
		f(k, v)
	}
}

func (m *EnhancedMap[K, V]) ContainsKey(k K) bool {
	_, ok := m.m[k]
	return ok
}

func (m *EnhancedMap[K, V]) ContainsValue(value V) bool {
	for _, v := range m.m {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

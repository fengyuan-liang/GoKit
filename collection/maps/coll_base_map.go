// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package maps

import "reflect"

// baseMap abstract map
type baseMap[K comparable, V any] struct {
	m map[K]V
}

func (m *baseMap[K, V]) Get(k K) (v V, ok bool) {
	v, ok = m.m[k]
	return v, ok
}

func (m *baseMap[K, V]) MustGet(k K) (v V) {
	v, _ = m.m[k]
	return v
}

func (m *baseMap[K, V]) Size() int {
	return len(m.m)
}

func (m *baseMap[K, V]) IsEmpty() bool {
	return m.Size() == 0
}

func (m *baseMap[K, V]) ContainsKey(k K) bool {
	_, ok := m.m[k]
	return ok
}

func (m *baseMap[K, V]) ContainsValue(value V) bool {
	for _, v := range m.m {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

func (m *baseMap[K, V]) Clear() {
	m.m = make(map[K]V)
}

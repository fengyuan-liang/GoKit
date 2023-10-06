// Copyright 2023 QINIU. All rights reserved
// @Description:
// @Version: 1.0.0
// @Date: 2023/10/06 16:02
// @Author: liangfengyuan@qiniu.com

package maps

import "reflect"

// baseMap 抽象map
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

func (m *baseMap[K, V]) RawMap() map[K]V {
	return m.m
}

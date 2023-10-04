// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package maps

import "GoKit/collection"

func NewTreeMap[K comparable, V any]() IMap[K, V] {
	return new(TreeMap[K, V])
}

func NewTreeMapWithComparator[K comparable, V any](comparator collection.CompareFunc[K]) IMap[K, V] {
	return &TreeMap[K, V]{
		comparator: comparator,
	}
}

// TreeMap A Red-Black tree based NavigableMap implementation.
// The map is sorted according to the natural ordering of its
// keys, or by a Comparator provided at map creation time,
// depending on which constructor is used.
type TreeMap[K comparable, V any] struct {
	comparator collection.CompareFunc[K]
}

func (t TreeMap[K, V]) Put(k K, v V) {
	//TODO implement me
	panic("implement me")
}

func (t TreeMap[K, V]) Get(k K) (v V, ok bool) {
	//TODO implement me
	panic("implement me")
}

func (t TreeMap[K, V]) MustGet(k K) (v V) {
	//TODO implement me
	panic("implement me")
}

func (t TreeMap[K, V]) GetOrDefault(k K, defaultValue V) (v V) {
	//TODO implement me
	panic("implement me")
}

func (t TreeMap[K, V]) Remove(k K) {
	//TODO implement me
	panic("implement me")
}

func (t TreeMap[K, V]) KeySet() []K {
	//TODO implement me
	panic("implement me")
}

func (t TreeMap[K, V]) Values() []V {
	//TODO implement me
	panic("implement me")
}

func (t TreeMap[K, V]) Size() int {
	//TODO implement me
	panic("implement me")
}

func (t TreeMap[K, V]) ForEach(f func(K, V)) {
	//TODO implement me
	panic("implement me")
}

func (t TreeMap[K, V]) ContainsKey(k K) bool {
	//TODO implement me
	panic("implement me")
}

func (t TreeMap[K, V]) ContainsValue(value V) bool {
	//TODO implement me
	panic("implement me")
}

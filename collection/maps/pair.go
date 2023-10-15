// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package maps

type Pair[K any, V any] struct {
	Key   K `json:"key"`
	Value V `json:"value"`
}

func MapToPair[K comparable, V any](m IMap[K, V]) []*Pair[K, V] {
	pairs := make([]*Pair[K, V], 0)
	m.ForEach(func(k K, v V) {
		pairs = append(pairs, &Pair[K, V]{k, v})
	})
	return pairs
}

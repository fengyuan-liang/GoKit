// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package maps

// IMap map
// @see TreeMap
// @see EnhancedMap enhance golang map
// @see LinkedHashMap
type IMap[K comparable, V any] interface {
	Put(k K, v V)
	Get(k K) (v V, ok bool)
	GetOrDefault(k K, defaultValue V) (v V)
	Remove(k K)
	KeySet() []K
	Values() []V
	Size() int
	ForEach(f func(K, V))
	ContainsKey(k K) bool
	ContainsValue(value V) bool
}

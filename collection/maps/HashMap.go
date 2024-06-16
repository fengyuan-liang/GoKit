// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package maps

// HashMap Based on red and black trees
// The key of a hashMap is based on a red-black tree at the underlying level and must be comparable.
// If the key is a primitive type, it will be compared according to the primitive type. If it is a
// struct and implements the collection.Comparable[K] interface, it will be compared according to
// the rules of the interface. If it does not implement the interface but a CompareFunc is provided,
// the comparison will follow this function. If neither is implemented, the struct's memory address
// or MD5 value will be used for sorting.
type HashMap[K any, V any] struct {
}

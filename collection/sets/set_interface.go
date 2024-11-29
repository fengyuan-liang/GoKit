// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sets

import "github.com/fengyuan-liang/GoKit/collection"

// ISet interface that all sets implement
type ISet[T any] interface {
	collection.JSONSerializer
	collection.JSONDeserializer

	Add(elements ...T)
	Remove(elements ...T)
	Contains(elements ...T) bool
	IsEmpty() bool
	Size() int
	Clear()
	Values() []T
	String() string
}

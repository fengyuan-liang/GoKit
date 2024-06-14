// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package heaps

type IHeap[E any] interface {
	Size() int
	IsEmpty() bool
	Clear()
	Add(element E)
}

// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package stream

import "github.com/fengyuan-liang/GoKit/collection"

type IStream[In any, Out any] interface {
	Filter(filterFunc func(ele In) bool) IStream[In, Out]
	Map(m func(ele In) Out) IStream[In, Out]
	CollectToSlice() []Out
	ForEach(action func(ele Out))
	Sort(compareFunc ...collection.CompareFunc[Out]) IStream[In, Out]
	Limit(cnt int) IStream[In, Out]
	Skip(cnt int) IStream[In, Out]
}

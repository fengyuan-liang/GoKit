// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package lists

func AsList[T any](arr ...T) IList[T] {
	list := NewArrayList[T]()
	for _, t := range arr {
		list.Add(t)
	}
	return list
}

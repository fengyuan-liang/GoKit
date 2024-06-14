// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package lists

import "fmt"

func AsList[T any](arr ...T) IList[T] {
	list := NewArrayList[T]()
	for _, t := range arr {
		list.Add(t)
	}
	return list
}

func checkNonnegative(value int, name string) int {
	if value < 0 {
		panic(fmt.Sprintf("%s cannot be negative but was:  %d", name, value))
	}
	return value
}

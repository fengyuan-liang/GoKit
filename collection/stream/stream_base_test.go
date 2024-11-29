// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package stream

import (
	"fmt"
	"testing"
)

func TestStream_Map(t *testing.T) {
	list := Of[int, int]([]int{1, 2, 3, 4, 5, 6, 7, 8}).
		Filter(func(element int) bool { return element%2 == 0 }).
		Skip(1).
		Limit(10).
		Map(func(element int) int { return element * 2 }).
		Sort(func(o1 int, o2 int) int { return -1 }).
		CollectToSlice()
	fmt.Printf("%v\n", list)
}

// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package stream

import (
	"fmt"
	"testing"
)

func TestStream_Map(t *testing.T) {
	list := Of[int, int]([]int{1, 2, 3, 4}).Filter(func(element int) bool {
		return element%2 == 0
	}).Filter(func(element int) bool {
		return element%2 == 0
	}).Map(func(element int) int {
		return element * 2
	}).CollectToSlice()
	fmt.Printf("%v\n", list)
}

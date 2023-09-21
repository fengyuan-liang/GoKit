// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package maps

import (
	"strings"
	"testing"
)

func TestNewTreeMap(t *testing.T) {
	treeMap := NewTreeMapWithComparator[string, int](func(o1 string, o2 string) int {
		return strings.Compare(o1, o2)
	})

	treeMap.Put("1", 1)
	treeMap.Put("2", 2)
	treeMap.Put("4", 4)
	treeMap.Put("3", 3)

	treeMap.ForEach(func(s string, i int) {
		t.Logf("k[%s], v[%d]", s, i)
	})
}

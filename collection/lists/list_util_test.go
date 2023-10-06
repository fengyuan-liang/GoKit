// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package lists

import "testing"

func TestAsList(t *testing.T) {
	list := AsList(1, 2, 3, 4, 5)
	t.Logf("%v", list.ToSlice())
}

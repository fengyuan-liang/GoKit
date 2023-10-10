// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import "testing"

func TestMax(t *testing.T) {
	t.Logf("The larger one in 1 and 2 is %d", Max(1, 2))
	t.Logf("The larger one in 1 and 2 is %s", Max("1", "2"))
	t.Logf("The larger one in 1.0 and 2.0 is %f", Max(1.0, 2.0))
}

func TestMin(t *testing.T) {
	t.Logf("The smaller one in 1 and 2 is %d", Min(1, 2))
	t.Logf("The smaller one in 1 and 2 is %s", Min("1", "2"))
	t.Logf("The smaller one in 1.0 and 2.0 is %f", Min(1.0, 2.0))
}

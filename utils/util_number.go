// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"GoKit/collection"
)

func Max[T any](o1, o2 T) T {
	if collection.Compare(o1, o2) > 0 {
		return o1
	}
	return o2
}

func Min[T any](o1, o2 T) T {
	if collection.Compare(o1, o2) < 0 {
		return o1
	}
	return o2
}

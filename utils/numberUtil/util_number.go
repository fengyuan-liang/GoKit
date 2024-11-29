// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package numberUtil

import (
	"github.com/fengyuan-liang/GoKit/collection"
	"regexp"
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

var numberRegex = regexp.MustCompile(`^[-+]?\d*\.?\d+$`)

func IsNumber(s string) bool {
	return numberRegex.MatchString(s)
}

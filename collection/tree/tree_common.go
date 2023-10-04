// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tree

import (
	"GoKit/collection"
	"fmt"
	"strings"
)

// Compare First, use comparator for comparison. If not, use Compare again. If not, use Golang's default sorting rules
func (b *BaseTree[E]) Compare(e1 E, e2 E) int {
	if b.comparator != nil {
		return b.comparator(e1, e2)
	}
	if b.CanCompare(e1, e2) {
		return (any)(e1).(collection.Comparable[E]).CompareTo(e2)
	}
	// If the object to be compared neither implements the Comparable interface nor passes in a comparator,
	// the default comparison rules of the standard type are used
	switch typeE1 := (any)(e1).(type) {
	case int:
		if typeE2, ok := (any)(e2).(int); ok {
			return typeE1 - typeE2
		}
	case int8:
		if typeE2, ok := (any)(e2).(int8); ok {
			return int(typeE1 - typeE2)
		}
	case int16:
		if typeE2, ok := (any)(e2).(int16); ok {
			return int(typeE1 - typeE2)
		}
	case int32:
		if typeE2, ok := (any)(e2).(int32); ok {
			return int(typeE1 - typeE2)
		}
	case int64:
		if typeE2, ok := (any)(e2).(int64); ok {
			return int(typeE1 - typeE2)
		}
	case uint:
		if typeE2, ok := (any)(e2).(uint); ok {
			return int(typeE1 - typeE2)
		}
	case uint16:
		if typeE2, ok := (any)(e2).(uint16); ok {
			return int(typeE1 - typeE2)
		}
	case uint32:
		if typeE2, ok := (any)(e2).(uint32); ok {
			return int(typeE1 - typeE2)
		}
	case uint64:
		if typeE2, ok := (any)(e2).(uint64); ok {
			return int(typeE1 - typeE2)
		}
	case float32:
		if typeE2, ok := (any)(e2).(float32); ok {
			return int(typeE1 - typeE2)
		}
	case float64:
		if typeE2, ok := (any)(e2).(float64); ok {
			return int(typeE1 - typeE2)
		}
	case string:
		if typeE2, ok := (any)(e2).(string); ok {
			return strings.Compare(typeE1, typeE2)
		}
	case bool:
		if typeE2, ok := (any)(e2).(bool); ok {
			if typeE1 == typeE2 {
				return 0
			} else if typeE1 {
				return +1
			} else {
				return -1
			}
		}
	}
	panic(fmt.Sprintf("element[%v][%v] cannot equal", e1, e2))
}

func (b *BaseTree[E]) CanCompare(e1 any, e2 any) bool {
	if _, ok := e1.(collection.Comparable[E]); !ok {
		return false
	}
	if _, ok := e2.(collection.Comparable[E]); !ok {
		return false
	}
	return true
}

// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package collection

import (
	"fmt"
	"strings"
)

func Equal[E any](e1 E, e2 E) bool {
	return Compare(e1, e2) == 0
}

// Compare First, use comparator for comparison. If not, use Compare again. If not, use Golang's default sorting rules
func Compare[E any](e1 E, e2 E) int {
	if CanCompare(e1, e2) {
		return (any)(e1).(Comparable[E]).CompareTo(e2)
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
	panic(fmt.Sprintf("element[%v] and [%v] cannot equal", e1, e2))
}

func CanCompare[E any](e1 E, e2 E) bool {
	if _, ok := (any)(e1).(Comparable[E]); !ok {
		return false
	}
	if _, ok := (any)(e1).(Comparable[E]); !ok {
		return false
	}
	return true
}

func Sort[T any](slice []T, compareFunc ...CompareFunc[T]) {
	if len(compareFunc) != 0 {
		QuickSort(slice, compareFunc[0])
		return
	}
	QuickSort(slice, func(o1, o2 T) int {
		return Compare(o1, o2)
	})
}

// QuickSort quick sort
func QuickSort[T any](slice []T, compareFunc func(o1, o2 T) int) {
	quickSort(slice, compareFunc, 0, len(slice)-1)
}

func quickSort[T any](slice []T, compareFunc func(o1, o2 T) int, low, high int) {
	if low < high {
		pivot := partition(slice, compareFunc, low, high)
		quickSort(slice, compareFunc, low, pivot-1)
		quickSort(slice, compareFunc, pivot+1, high)
	}
}

func partition[T any](slice []T, compareFunc func(o1, o2 T) int, low, high int) int {
	pivotIndex := low
	pivotValue := slice[high]

	for j := low; j < high; j++ {
		if compareFunc(slice[j], pivotValue) <= 0 {
			swap(slice, j, pivotIndex)
			pivotIndex++
		}
	}

	swap(slice, pivotIndex, high)
	return pivotIndex
}

func swap[T any](slice []T, i, j int) {
	var tmp T
	tmp = slice[i]
	slice[i] = slice[j]
	slice[j] = tmp
}

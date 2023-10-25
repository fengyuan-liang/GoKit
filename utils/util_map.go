// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import "github.com/fengyuan-liang/GoKit/collection/maps"

// SliceToMap Slice to Map
//
// example:
//
//	type Person struct {
//		ID int
//	}
//	p1 := &Person{ID: 1}
//	p2 := &Person{ID: 2}
//	p3 := &Person{ID: 3}
//	people := make([]*Person, 0)
//	people = append(people, p1, p2, p3)
//	m := SliceToMap(people, func(element *Person) int {
//		return element.ID
//	})
//	v, _ := m.Get(1)
func SliceToMap[A any, ID comparable](slice []A, f func(element A) ID) maps.IMap[ID, A] {
	enhancedMap := maps.NewEnhancedMap[ID, A]()
	for _, a := range slice {
		enhancedMap.Put(f(a), a)
	}
	return enhancedMap
}

func SliceMap[In any, Out any](slice []In, f func(in In) Out) []Out {
	outs := make([]Out, len(slice))
	for index, in := range slice {
		outs[index] = f(in)
	}
	return outs
}

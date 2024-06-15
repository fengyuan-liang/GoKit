// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceToMap(t *testing.T) {
	type Person struct {
		ID int
	}
	p1 := &Person{ID: 1}
	p2 := &Person{ID: 2}
	p3 := &Person{ID: 3}
	people := make([]*Person, 0)
	people = append(people, p1, p2, p3)
	m := SliceToMap(people, func(element *Person) int {
		return element.ID
	})
	fmt.Printf("%v\n", ObjToJsonStr(m))
	v, _ := m.Get(1)
	assert.Equal(t, p1, v)
	v, _ = m.Get(2)
	assert.Equal(t, p2, v)
	v, _ = m.Get(3)
	assert.Equal(t, p3, v)
}

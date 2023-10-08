// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"testing"
)

func TestCopySlice(t *testing.T) {
	type Person struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	type PersonDTO struct {
		Id string `json:"id"`
	}
	people := []Person{
		{
			Id:   "111",
			Name: "张三",
			Age:  18,
		},
		{
			Id:   "222",
			Name: "李四",
			Age:  19,
		},
	}
	copySlice := CopySlice[Person, PersonDTO](people)
	t.Logf("%v", ObjToJsonStr(copySlice))

}

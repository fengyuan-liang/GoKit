// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayList(t *testing.T) {
	// 测试 NewArrayList 和 Size 方法
	list := NewArrayList[int]()
	if list.Size() != 0 {
		t.Error("NewArrayList should create an empty list with size 0")
	}

	// 测试 Add 和 Get 方法
	list.Add(10)
	list.Add(20)
	list.Add(30)
	if list.Size() != 3 {
		t.Error("Size should return the correct number of elements in the list")
	}
	if list.Get(0) != 10 || list.Get(1) != 20 || list.Get(2) != 30 {
		t.Error("Get should return the correct element at the specified index")
	}

	// 测试 Set 方法
	list.Set(1, 25)
	if list.Get(1) != 25 {
		t.Error("Set should update the element at the specified index")
	}

	// 测试 Remove 方法
	removed := list.Remove(0)
	if removed != 10 || list.Size() != 2 || list.Get(0) != 25 {
		t.Error("Remove should remove the element at the specified index and return the removed element")
	}

	// 测试 Contains 和 IndexOf 方法
	if !list.Contains(25) || list.Contains(10) {
		t.Error("Contains should return true for elements that exist in the list and false for elements that don't")
	}
	if list.IndexOf(25) != 0 || list.IndexOf(10) != -1 {
		t.Error("IndexOf should return the correct index of the specified element or -1 if it doesn't exist")
	}

	// 测试 Clear 和 IsEmpty 方法
	list.Clear()
	if !list.IsEmpty() || list.Size() != 0 {
		t.Error("Clear should remove all elements from the list and reset the size to 0")
	}
}

func TestArrayList_Add(t *testing.T) {
	list := NewArrayList[int]()
	for i := 0; i < 100; i++ {
		list.Add(i)
	}
	t.Logf("%v", list.ToString())
	assert.Equal(t, 100, list.Size())
}

func TestArrayList_Contains(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	assert.Equal(t, true, list.Contains(1))
	assert.Equal(t, false, list.Contains(100))
}

// struct
type Person struct {
	Name string
	Age  int
}

func (p *Person) CompareTo(element *Person) int {
	return p.Age - element.Age
}

func TestArrayList_Add_struct(t *testing.T) {
	list := NewArrayList[*Person]()
	list.Add(&Person{Name: "张三", Age: 19})
	list.Add(&Person{Name: "李四", Age: 18})
	list.Add(&Person{Name: "王五", Age: 20})
	t.Logf("%v", list.ToString())
	assert.Equal(t, 3, list.Size())
}

func TestNewArrayList_Contains_struct(t *testing.T) {
	list := NewArrayList[*Person]()
	p := &Person{Name: "张三", Age: 19}
	list.Add(p)
	list.Add(&Person{Name: "李四", Age: 18})
	list.Add(&Person{Name: "王五", Age: 20})
	assert.Equal(t, true, list.Contains(p))
	assert.Equal(t, false, list.Contains(new(Person)))
}

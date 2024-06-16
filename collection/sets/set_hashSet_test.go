// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sets

import (
	"encoding/json"
	"testing"
)

func TestHashSet_Add(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1, 2, 3)

	if set.Size() != 3 {
		t.Errorf("Expected set size 3, but got %d", set.Size())
	}
}

func TestHashSet_Remove(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1, 2, 3)
	set.Remove(2)

	if set.Size() != 2 {
		t.Errorf("Expected set size 2, but got %d", set.Size())
	}

	if set.Contains(2) {
		t.Errorf("Expected set to not contain 2")
	}
}

func TestHashSet_Contains(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1, 2, 3)

	if !set.Contains(1, 2) {
		t.Errorf("Expected set to contain 1 and 2")
	}

	if set.Contains(4) {
		t.Errorf("Expected set to not contain 4")
	}
}

func TestHashSet_IsEmpty(t *testing.T) {
	set := NewHashSet[int]()

	if !set.IsEmpty() {
		t.Errorf("Expected set to be empty")
	}

	set.Add(1)

	if set.IsEmpty() {
		t.Errorf("Expected set to not be empty")
	}
}

func TestHashSet_Clear(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1, 2, 3)
	set.Clear()

	if set.Size() != 0 {
		t.Errorf("Expected set size 0, but got %d", set.Size())
	}
}

func TestHashSet_serialization(t *testing.T) {
	// Marshal
	set := NewHashSet[int]()
	set.Add(1, 2, 3, 4, 5)
	data, _ := json.Marshal(set)
	println(string(data))
	// UnMarshal
	set.Clear()
	_ = json.Unmarshal([]byte(`[1,2,3,4,5]`), &set)
	println(set.String())
}

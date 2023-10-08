// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package maps

import (
	"fmt"
	"testing"
)

func TestLinkedHashMap(t *testing.T) {
	m := NewLinkedHashMap[string, int]()

	m.Put("one", 1)
	m.Put("two", 2)
	m.Put("three", 3)

	if size := m.Size(); size != 3 {
		t.Errorf("Expected size 3, got %d", size)
	}

	if v, ok := m.Get("two"); !ok || v != 2 {
		t.Errorf("Expected 2, got %d", v)
	}

	if v := m.MustGet("three"); v != 3 {
		t.Errorf("Expected 3, got %d", v)
	}

	if v := m.GetOrDefault("four", 4); v != 4 {
		t.Errorf("Expected 4, got %d", v)
	}

	m.Remove("one")

	if size := m.Size(); size != 2 {
		t.Errorf("Expected size 2, got %d", size)
	}

	if m.ContainsKey("one") {
		t.Errorf("Expected false, got true")
	}

	if !m.ContainsValue(2) {
		t.Errorf("Expected true, got false")
	}

	m.ForEach(func(k string, v int) {
		fmt.Printf("%s: %d\n", k, v)
	})

	if m.IsEmpty() {
		t.Errorf("Expected false, got true")
	}

	m.PutAll(map[string]int{"five": 5, "six": 6})

	if size := m.Size(); size != 4 {
		t.Errorf("Expected size 4, got %d", size)
	}

	keys := m.KeySet()
	expectedKeys := []string{"two", "three", "five", "six"}
	for i, k := range keys {
		if k != expectedKeys[i] {
			t.Errorf("Expected %s, got %s", expectedKeys[i], k)
		}
	}

	values := m.Values()
	expectedValues := []int{2, 3, 5, 6}
	for i, v := range values {
		if v != expectedValues[i] {
			t.Errorf("Expected %d, got %d", expectedValues[i], v)
		}
	}
}

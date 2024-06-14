// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package maps

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEnhancedMap(t *testing.T) {
	m := NewEnhancedMap[string, int]()

	// Test Put and Get
	m.Put("hello", 1)
	v, ok := m.Get("hello")
	if !ok || v != 1 {
		t.Errorf("Put or Get failed. Got %v, expected 1", v)
	}

	// Test Size
	if size := m.Size(); size != 1 {
		t.Errorf("Size failed. Got %v, expected 1", size)
	}

	// Test KeySet
	keySet := m.KeySet()
	if len(keySet) != 1 || keySet[0] != "hello" {
		t.Errorf("KeySet failed. Got %v, expected [\"hello\"]", keySet)
	}

	// Test Values
	values := m.Values()
	if len(values) != 1 || values[0] != 1 {
		t.Errorf("Values failed. Got %v, expected [1]", values)
	}

	// Test ContainsKey
	if !m.ContainsKey("hello") {
		t.Errorf("ContainsKey failed. Expected true, got false.")
	}

	// Test ContainsValue
	if !m.ContainsValue(1) {
		t.Errorf("ContainsValue failed. Expected true, got false.")
	}

	// Test Remove
	m.Remove("hello")
	if size := m.Size(); size != 0 {
		t.Errorf("Remove or Size failed. Got %v, expected 0", size)
	}

	// Test ForEach
	m.Put("world", 2)
	m.ForEach(func(k string, v int) {
		if k != "world" || v != 2 {
			t.Errorf("ForEach failed. Got key: %v, value: %v, expected key: \"world\", value: 2", k, v)
		}
	})
}

func TestEnhancedMapSerialization(t *testing.T) {
	// test Marshal
	m := NewEnhancedMap[string, int]()
	m.Put("one", 1)
	m.Put("two", 2)
	m.Put("three", 3)
	data, _ := json.Marshal(m)
	fmt.Printf("%v\n", string(data))
	// test UnMarshal
	m.Clear()
	_ = json.Unmarshal([]byte(`{"one":1,"three":3,"two":2}`), &m)
	m.ForEach(func(k string, v int) {
		fmt.Printf("k:%v, v:%v\n", k, v)
	})
}

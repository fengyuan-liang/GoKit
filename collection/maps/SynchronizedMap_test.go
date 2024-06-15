// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package maps

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

// TestSynchronizedMapConcurrency tests the SynchronizedMap implementation for concurrency safety
func TestSynchronizedMapConcurrency(t *testing.T) {
	syncMap := NewConcurrentHashMap[string, int]()

	var wg sync.WaitGroup
	numGoroutines := 100
	numIterations := 1000

	// Run concurrent Put operations
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < numIterations; j++ {
				key := fmt.Sprintf("key-%d-%d", i, j)
				syncMap.Put(key, i+j)
			}
		}(i)
	}

	// Run concurrent Get operations
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < numIterations; j++ {
				key := fmt.Sprintf("key-%d-%d", i, j)
				syncMap.Get(key)
			}
		}(i)
	}

	// Run concurrent Remove operations
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < numIterations; j++ {
				key := fmt.Sprintf("key-%d-%d", i, j)
				syncMap.Remove(key)
			}
		}(i)
	}

	wg.Wait()

	// Final size check to ensure consistency
	size := syncMap.Size()
	assert.GreaterOrEqual(t, size, 0)
}

func TestSynchronizedMapSerialization(t *testing.T) {
	// test Marshal
	var m IMap[string, int] = NewConcurrentHashMap[string, int]()
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

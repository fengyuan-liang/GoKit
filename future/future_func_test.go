// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package future

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAnonymousFunc(t *testing.T) {
	futureFunc := FutureFunc[int](func() int {
		return Fibonacci(40)
	})
	result, err := futureFunc.GetWithTimeout(time.Second * 5)
	assert.Equal(t, 102334155, result)
	assert.Nil(t, err)
}

func TestFutureFuncWithHandlerError(t *testing.T) {
	futureFunc := FutureFunc[int](Add, 10, 20)
	result, err := futureFunc.Get()
	assert.Equal(t, 30, result)
	assert.NotNil(t, err)
}

func TestFutureFunc(t *testing.T) {
	x := 10
	var elapsed time.Duration
	start := time.Now()
	future := FutureFunc[int](func() int {
		printTime(t)
		time.Sleep(5 * time.Second)
		fmt.Printf("x = %v\n", x)
		return x * 10
	})
	elapsed = time.Since(start)

	t.Logf("it took %s", elapsed)
	assert.Less(t, elapsed.Milliseconds(), (1 * time.Second).Milliseconds())

	result, _ := future.Get()
	elapsed = time.Since(start)
	assert.Less(t, (5 * time.Second).Milliseconds(), elapsed.Milliseconds())
	assert.Equal(t, 100, result)
	t.Logf("Result: %v\n", result)

	// This assert tests calling result second time doesn't cause any problems
	assert.Equal(t, 100, result)
}

func TestFutureFuncTimeOut(t *testing.T) {
	x := 10
	var elapsed time.Duration
	start := time.Now()
	future := FutureFunc[int](func() int {
		printTime(t)
		time.Sleep(5 * time.Second)
		fmt.Printf("x = %v\n", x)
		return x * 10
	})
	elapsed = time.Since(start)

	t.Logf("it took %s", elapsed)
	assert.Less(t, elapsed.Milliseconds(), (1 * time.Second).Milliseconds())

	result, err := future.GetWithTimeout(3 * time.Second)
	elapsed = time.Since(start)
	assert.Equal(t, nil, result)
	assert.Less(t, elapsed.Milliseconds(), (4 * time.Second).Milliseconds())
	t.Logf("Result: %v\n", result)
	t.Logf("err: %v\n", err.Error())
	assert.Equal(t, nil, result)
}

func TestFuturePanic(t *testing.T) {
	futureFunc := FutureFunc[int](func() int {
		panic("panic")
	})
	if _, err := futureFunc.Get(); err != nil {
		t.Logf("is ErrPanic:%v", err.Error())
	}
}

func printTime(t *testing.T) {
	t.Logf("time: %v\n", time.Now().Unix())
}

// Fibonacci returns the nth Fibonacci number
func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Add(a, b int) (int, error) {
	return a + b, errors.New("you can return error")
}

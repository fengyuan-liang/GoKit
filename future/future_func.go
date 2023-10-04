// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package future

import (
	"errors"
	"fmt"
	"reflect"
	"runtime/debug"
	"time"
)

var (
	// ErrTimeOut timed out
	ErrTimeOut = errors.New("timed out")
)

// Future type holds Result and state
type Future[T any] struct {
	Success          bool
	Done             bool
	Result           T
	InterfaceChannel <-chan T
	err              error
	errChannel       <-chan error
}

// Get return the result when available. This is a blocking call
func (f *Future[T]) Get() (T, error) {
	if f.Done {
		return f.Result, f.err
	}
	f.Result = <-f.InterfaceChannel
	f.err = <-f.errChannel
	f.Success = true
	f.Done = true
	return f.Result, f.err
}

// GetWithTimeout return the result until timeout.
func (f *Future[T]) GetWithTimeout(timeout time.Duration) (T, error) {
	if f.Done {
		return f.Result, f.err
	}
	timeoutChannel := time.After(timeout)
	select {
	case res := <-f.InterfaceChannel:
		f.Result = res
		f.err = <-f.errChannel
		f.Success = true
		f.Done = true
	case <-timeoutChannel:
		f.Done = true
		f.Success = false
		f.err = ErrTimeOut
	}
	return f.Result, f.err
}

// FutureFunc creates a function that returns its response in future
func FutureFunc[T any](implem interface{}, args ...interface{}) *Future[T] {
	valIn := make([]reflect.Value, len(args), len(args))

	fnVal := reflect.ValueOf(implem)

	for idx, elt := range args {
		valIn[idx] = reflect.ValueOf(elt)
	}
	interfaceChannel := make(chan T, 1)
	errChannel := make(chan error, 1)
	go func() {
		defer func() {
			// handle the panic here
			if r := recover(); r != nil {
				// print the stack trace
				debug.PrintStack()
				errChannel <- fmt.Errorf("panic in futureFunc, err: %v", r)
			}
			close(interfaceChannel)
			close(errChannel)
		}()
		res := fnVal.Call(valIn)
		// Up to two return values are supported
		if len(res) > 1 && !res[1].IsNil() {
			// handle err
			errChannel <- res[1].Interface().(error)
		} else {
			// handle err
			errChannel <- nil
		}
		result := res[0].Interface()
		if result != nil {
			interfaceChannel <- result.(T)
		}
	}()

	return &Future[T]{
		Success:          false,
		Done:             false,
		InterfaceChannel: interfaceChannel,
		errChannel:       errChannel,
	}
}

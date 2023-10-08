// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"github.com/jinzhu/copier"
	"reflect"
)

// CopySlice 进行slice的拷贝
func CopySlice[In any, Out any](in []In) []Out {
	outs := make([]Out, len(in))
	outType := reflect.TypeOf((*Out)(nil)).Elem()
	for i := 0; i < len(in); i++ {
		var out Out
		if outType.Kind() == reflect.Ptr {
			out = reflect.New(outType.Elem()).Interface().(Out)
		} else {
			out = reflect.New(outType).Elem().Interface().(Out)
		}
		_ = copier.Copy(&out, in[i])
		outs[i] = out
	}
	return outs
}

// DeepCopySlice 进行slice的拷贝
func DeepCopySlice[In any](in []In) []In {
	outs := make([]In, len(in))
	outType := reflect.TypeOf((*In)(nil)).Elem()
	for i := 0; i < len(in); i++ {
		var out In
		if outType.Kind() == reflect.Ptr {
			out = reflect.New(outType.Elem()).Interface().(In)
		} else {
			out = reflect.New(outType).Elem().Interface().(In)
		}
		_ = copier.Copy(&out, in[i])
		outs[i] = out
	}
	return outs
}

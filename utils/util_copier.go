// Copyright 2023 QINIU. All rights reserved
// @Description: 运行出错的报警
// @Version: 1.0.0
// @Date: 2023/08/17 10:27
// @Author: liangfengyuan@qiniu.com

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

// Copyright 2023 QINIU. All rights reserved
// @Description: 运行出错的报警
// @Version: 1.0.0
// @Date: 2023/08/17 10:27
// @Author: liangfengyuan@qiniu.com

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

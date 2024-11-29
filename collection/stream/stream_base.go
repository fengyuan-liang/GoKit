// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package stream

import (
	"github.com/fengyuan-liang/GoKit/collection"
	"github.com/fengyuan-liang/GoKit/utils/numberUtil"
)

type Stream[In any, Out any] struct {
	Data    []In
	outData []Out
	mapCnt  int
}

func (s *Stream[In, Out]) Filter(filterFunc func(ele In) bool) IStream[In, Out] {
	filteredData := make([]In, 0)
	for _, ele := range s.Data {
		if filterFunc(ele) {
			filteredData = append(filteredData, ele)
		}
	}
	s.Data = filteredData
	return s
}

func (s *Stream[In, Out]) Map(m func(ele In) Out) IStream[In, Out] {
	if s.mapCnt != 0 {
		panic("Only allow map once")
	}
	mappedData := make([]Out, 0)
	for _, datum := range s.Data {
		mappedData = append(mappedData, m(datum))
	}
	s.outData = mappedData
	s.mapCnt += 1
	return s
}

func (s *Stream[In, Out]) CollectToSlice() []Out {
	if s.mapCnt == 0 {
		s.Map(func(element In) Out {
			return (any)(element).(Out)
		})
	}
	return s.outData
}

func (s *Stream[In, Out]) ForEach(action func(ele Out)) {
	if s.mapCnt == 0 {
		s.Map(func(element In) Out {
			return (any)(element).(Out)
		})
	}
	for _, datum := range s.outData {
		action(datum)
	}
}

func (s *Stream[In, Out]) Sort(compareFunc ...collection.CompareFunc[Out]) IStream[In, Out] {
	if s.mapCnt == 0 {
		s.Map(func(element In) Out {
			return (any)(element).(Out)
		})
	}
	if len(s.outData) < 2 {
		return s
	}
	if len(compareFunc) != 0 {
		collection.QuickSort(s.outData, compareFunc[0])
		return s
	}

	collection.QuickSort(s.outData, func(o1, o2 Out) int { return collection.Compare(o1, o2) })

	return s
}

func (s *Stream[In, Out]) Limit(cnt int) IStream[In, Out] {
	s.Data = s.Data[:numberUtil.Min(cnt, len(s.Data))]
	return s
}

func (s *Stream[In, Out]) Skip(cnt int) IStream[In, Out] {
	s.Data = s.Data[numberUtil.Min(cnt, len(s.Data)):]
	return s
}

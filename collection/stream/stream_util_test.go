package stream

import (
	"github.com/fengyuan-liang/GoKit/collection"
	"testing"
)

func TestStream_Sort(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8}
	collection.QuickSort(ints, func(o1, o2 int) int {
		return o2 - o1
	})
	t.Logf("%v", ints)
}

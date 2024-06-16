package sets

import "testing"

func Test01(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1, 2, 3, 4)
	println(set)
}

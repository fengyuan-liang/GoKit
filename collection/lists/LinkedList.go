// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package lists

func NewLinkedList[E any]() IList[E] {
	return new(LinkedList[E])
}

type LinkedList[E any] struct {
	BaseList[E]
}

func (l *LinkedList[E]) Contains(element E) bool {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[E]) Add(element E) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[E]) AddAll(elements []E) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[E]) AddAtIndex(index int, element E) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[E]) Get(index int) E {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[E]) Set(index int, element E) E {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[E]) Remove(index int) (oldElement E) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[E]) IndexOf(element E) int {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[E]) Clear() {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[E]) ToString() string {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList[E]) ToSlice() []E {
	//TODO implement me
	panic("implement me")
}

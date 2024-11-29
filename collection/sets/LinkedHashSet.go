// Copyright The GoKit authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sets

type LinkedHashSet[T comparable] struct {
	innerHashSet ISet[T]
	keys         []T
}

func NewLinkedHashSet[T comparable]() ISet[T] {
	return &LinkedHashSet[T]{
		innerHashSet: NewHashSet[T](),
		keys:         make([]T, 0),
	}
}

func (l *LinkedHashSet[T]) Add(elements ...T) {
	l.innerHashSet.Add(elements...)
}

func (l *LinkedHashSet[T]) Remove(elements ...T) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedHashSet[T]) Contains(elements ...T) bool {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedHashSet[T]) Values() []T {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedHashSet[T]) IsEmpty() bool {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedHashSet[T]) Size() int {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedHashSet[T]) Clear() {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedHashSet[T]) String() string {
	//TODO implement me
	panic("implement me")
}

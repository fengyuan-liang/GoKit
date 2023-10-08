# GoKit
GoKit: is your ultimate toolbox of utilities for seamless development in Go 😉

## import

```shell
go get github.com/fengyuan-liang/GoKit
```


## 1. collection

- maps 
  - EnhancedMap：enhance go raw map
  - LinkedHashMap：LinkedHashMap is a data structure that combines the features of a hash table and a linked list, `providing predictable iteration order based on the insertion sequence`.
  - HashMap：The map with a underlying data structure lower than a red-black tree.
  - TreeMap：TreeMap in Java is a data structure that allows the storage of key-value pairs in a sorted order based on the keys, providing operations like insertion, deletion, and retrieval with logarithmic time complexity.
- lists
  - ArrayList：enhance go slice
  - LinkedList：LinkedList is a data structure that implements a sequence of elements using a doubly-linked list as its underlying structure.

example

```go
func TestLinkedHashMap(t *testing.T) {
    m := maps.NewLinkedHashMap[string, int]()

    m.Put("one", 1)
    m.Put("two", 2)
    m.Put("three", 3)

    m.ForEach(func(k string, v int) {
        t.Logf("%s: %d\n", k, v)
    })
}
```

```go
$ go test -run TestLinkedHashMap
one: 1
two: 2
three: 3
PASS
ok      GoKit/collection/maps   0.166s
```

## 2. stream

In Go, there are various ways to manipulate collections, and the `stream` library provides a convenient and practical approach, particularly for those familiar with Java. By leveraging the power of functional programming, the stream package enables seamless operations on collections, allowing for concise and expressive code. With 'stream', developers can effortlessly perform transformations, filtering, mapping, and aggregations on data, simplifying complex data processing tasks and enhancing code readability.

example

```go
func TestStream_Map(t *testing.T) {
	list := stream.Of[int, int]([]int{1, 2, 3, 4}).Filter(func(element int) bool {
		return element%2 == 0
	}).Filter(func(element int) bool {
		return element%2 == 0
	}).Map(func(element int) int {
		return element * 2
	}).CollectToSlice()
	t.Logf("%v", list)
}
```

```shell
$ go test -run TestStream_Map
[2 6]
PASS
ok      GoKit/collection/stream 0.164s
```

## 3. future

Go-future gives an implementation similar to Java/Scala Futures.

Although there are many ways to handle this behaviour in Golang. This library is useful for people who got used to Java/Scala Future implementation.

example

```go
func TestFutureFunc(t *testing.T) {
	futureFunc := future.FutureFunc[int](func() int {
		time.Sleep(5 * time.Second)
		return 1 * 10
	})
	// do something else here

	// get result when needed
	result, err := futureFunc.Get()
	fmt.Printf("result:%v, err:%v\n", result, err)
}
```

```shell
$ go test -run TestFutureFunc
result:10, err:<nil>
PASS
ok      GoKit/collection/stream 5.177s
```

## 4. utils

The 'utils' package encompasses a majority of commonly used utility methods in Go development. It provides a comprehensive set of tools that are frequently utilized during the development process. 

example

```go
func TestSliceToMap(t *testing.T) {
	type Person struct {
		ID int
	}
	p1 := &Person{ID: 1}
	p2 := &Person{ID: 2}
	p3 := &Person{ID: 3}
	people := make([]*Person, 0)
	people = append(people, p1, p2, p3)
    // k:ID v:pe
	m := utils.SliceToMap(people, func(element *Person) int { return element.ID })
	fmt.Printf("%v\n", utils.ObjToJsonStr(m.RawMap()))
}
```

```shell
$ go test -run TestSliceToMap
{"1":{"ID":1},"2":{"ID":2},"3":{"ID":3}}
PASS
ok      GoKit/utils     0.176s
```


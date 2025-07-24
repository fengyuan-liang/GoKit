# GoKit

[en](https://github.com/fengyuan-liang/GoKit/blob/main/README_en.md)

GoKit是您在Go开发中的终极工具箱😉

## 导入

```
go get github.com/fengyuan-liang/GoKit
```

## 1. collection

- maps
  - EnhancedMap：增强原生Go map
  - LinkedHashMap：LinkedHashMap是将哈希表和链表的特性结合在一起的数据结构，`根据插入顺序提供可预测的迭代顺序`。
  - HashMap：底层数据结构低于红黑树的映射。
  - TreeMap：TreeMap是一种基于红黑树数据结构，它允许按键的排序顺序存储键值对，并提供插入、删除和检索等操作，时间复杂度为对数级别。
  - SynchronizedMap：线程安全map装饰器，可以装饰所有map，使其线程安全。maps下的每一种map都提供了线程安全的版本。
- sets
  - HashSet：基于map实现的HaseSet

- lists
  - ArrayList：增强Go切片
  - LinkedList：LinkedList是一种使用双向链表作为其底层结构来实现元素序列的数据结构。

示例

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

```shell
$ go test -run TestLinkedHashMap
one: 1
two: 2
three: 3
PASS
ok      GoKit/collection/maps   0.166s
```

当然也可以完成序列化和反序列化

```go
func TestLinkedHashMap_Serialization(t *testing.T) {
	// test Marshal
	m := NewLinkedHashMap[string, int]()
	m.Put("one", 1)
	m.Put("two", 2)
	m.Put("three", 3)
	data, _ := json.Marshal(m)
	fmt.Printf("%v\n", string(data))
	// test UnMarshal
	m.Clear()
	_ = json.Unmarshal([]byte(`{"two":2,"one":1,"three":3}`), &m)
	m.ForEach(func(k string, v int) {
		fmt.Printf("k:%v, v:%v\n", k, v)
	})
}
```

```shell
$ go test -run TestLinkedHashMap_Serialization
{"one":1,"two":2,"three":3}
k:one, v:1
k:three, v:3
k:two, v:2
PASS
ok      github.com/fengyuan-liang/GoKit/collection/maps 0.131s
```

线程安全的map

```go
// TestSynchronizedMapConcurrency tests the SynchronizedMap implementation for concurrency safety
func TestSynchronizedMapConcurrency(t *testing.T) {
	syncMap := NewConcurrentHashMap[string, int]()

	var wg sync.WaitGroup
	numGoroutines := 100
	numIterations := 1000

	// Run concurrent Put operations
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < numIterations; j++ {
				key := fmt.Sprintf("key-%d-%d", i, j)
				syncMap.Put(key, i+j)
			}
		}(i)
	}

	// Run concurrent Get operations
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < numIterations; j++ {
				key := fmt.Sprintf("key-%d-%d", i, j)
				syncMap.Get(key)
			}
		}(i)
	}

	// Run concurrent Remove operations
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < numIterations; j++ {
				key := fmt.Sprintf("key-%d-%d", i, j)
				syncMap.Remove(key)
			}
		}(i)
	}

	wg.Wait()

	// Final size check to ensure consistency
	size := syncMap.Size()
	assert.GreaterOrEqual(t, size, 0)
}
```

```shell
$ go test -run TestSynchronizedMapConcurrency
PASS
ok      github.com/fengyuan-liang/GoKit/collection/maps 0.219s
```

## 2. stream

在Go中，有多种方法可以操作集合，而`stream`库提供了一种方便实用的方法，特别适合熟悉Java的人。通过利用函数式编程的威力，stream包能够对集合进行无缝操作，使代码简洁而富有表现力。使用`stream`，开发人员可以轻松进行数据的转换、过滤、映射和聚合，简化复杂的数据处理任务，提高代码的可读性。

示例

```go
func TestStream_Map(t *testing.T) {
	list := Of[int, int]([]int{1, 2, 3, 4, 5, 6, 7, 8}).
		Filter(func(element int) bool { return element%2 == 0 }).
		Skip(1).
		Limit(10).
		Map(func(element int) int { return element * 2 }).
		CollectToSlice()
	t.Logf("%v\n", list)
}
```

```shell
$ go test -run TestStream_Map
[8 12 16]
PASS
ok      GoKit/collection/stream 0.00s
```

## 3. future

Go-future提供了类似于Java/Scala Future的实现。

尽管在Golang中有很多处理此行为的方法，但对于习惯了Java/Scala Future实现的人来说，这个库非常有用。

示例

```go
func TestFutureFunc(t *testing.T) {
	futureFunc := future.FutureFunc[int](func() int {
		time.Sleep(5 * time.Second)
		return 1 * 10
	})
	// 在此处执行其他操作

	// 在需要时获取结果
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

`utils`包涵盖了Go开发中大多数常用的实用方法。它提供了一套全面的工具，这些工具在开发过程中经常被使用。

示例

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
    // k:ID v:person
	m:= utils.SliceToMap(people, func(element *Person) int { return element.ID })
	fmt.Printf("%v\n", utils.ObjToJsonStr(m.RawMap()))
}
```

```shell
$ go test -run TestSliceToMap
{"1":{"ID":1},"2":{"ID":2},"3":{"ID":3}}
PASS
ok      GoKit/utils     0.176s
```

## 5. 更完整的库

当然上述的数据结构只是笔者自用，请勿在生产项目使用，您可以在下面的仓库中找到更加合适的

- https://github.com/emirpasic/gods


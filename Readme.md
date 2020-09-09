## common-lang-in-go

[![API Reference](
https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667
)](https://pkg.go.dev/github.com/sjsdfg/common-lang-in-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/ethereum/go-ethereum)](https://goreportcard.com/report/github.com/sjsdfg/common-lang-in-go)

Java 程序员编写的类似 common-lang 工具包
包名就是要用大驼峰，不想改成 go 规范的包名。自己写的任性

## how to use

```shell script
go get github.com/sjsdfg/common-lang-in-go
```

### StringUtils

- IsEmpty(str string) bool
- IsNotEmpty(str string) bool
- IsAllEmpty(list ...string) bool
- IsAnyEmpty(list ...string) bool
- IsAnyNoneEmpty(list ...string) bool
- IsBlank(str string) bool
- IsNotBlank(str string) bool
- IsZero(str string) bool
- IsNotZero(str string) bool
- IsAnyZero(list ...string) bool
- IsAllZero(list ...string) bool
- IsAnyNoneZero(list ...string) bool
- Equal(str1, str2 string) bool
- EqualIgnoreCase(str1, str2 string) bool
- EqualsAny(str string, list ...string) bool
- EqualsAnyIgnoreCase(str string, list ...string) bool
- IsDigital(str string) bool
- DefaultIfEmpty(str, defaultStr string) string
- If(condition bool, ifTrue, ifFalse string) string
- Truncate(str string, startIndex, endIndex int) string

## CollectionUtils

### match functions

- AllMatch(list interface{}, action matchFunc) bool
- AnyMatch(list interface{}, action matchFunc) bool
- NoneMatch(list interface{}, action matchFunc) bool

**Use Case**
```go
func TestAllMatch(t *testing.T) {
	testCase := []string{"a", "a", "a"}
	assert.Equal(t, true, AllMatch(testCase, func(index int) bool {
		return testCase[index] == "a"
	}))

	testCase = []string{"a", "a", "b"}
	assert.Equal(t, false, AllMatch(testCase, func(index int) bool {
		return testCase[index] == "a"
	}))
}

func TestAnyMatch(t *testing.T) {
	testCase := []string{"a", "a", "b"}
	assert.Equal(t, true, AnyMatch(testCase, func(index int) bool {
		return testCase[index] == "b"
	}))

	testCase = []string{"a", "a", "b"}
	assert.Equal(t, false, AnyMatch(testCase, func(index int) bool {
		return testCase[index] == "c"
	}))
}

func TestNoneMatch(t *testing.T) {
	testCase := []string{"a", "a", "b"}
	assert.Equal(t, true, NoneMatch(testCase, func(index int) bool {
		return testCase[index] == "c"
	}))

	testCase = []string{"a", "a", "b"}
	assert.Equal(t, false, NoneMatch(testCase, func(index int) bool {
		return testCase[index] == "b"
	}))
}
```

### empty judgement

- IsEmpty(collection interface{})
- IsNotEmpty(collection interface{}) bool

### map function

这个是由于没有 Java Stream 里面的 map，导致每次提取某一类型的字段成为数组，写的代码十分冗余。所以抽象出来的。

由于不支持泛型，只能书写基本类型的方法调用。不过我的个人场景就是提取某一个字段（尤其是 id），目前还够用。

- MapToStringSlice(list interface{}, action func(index int) string) []string
- MapToIntSlice(list interface{}, action func(index int) int) []int
- MapToInt64Slice(list interface{}, action func(index int) int64) []int64
- MapToFloat64Slice(list interface{}, action func(index int) float64) []float64
- MapToFloat32Slice(list interface{}, action func(index int) float32) []float32

**Use Case**
```go
func TestMapToStringSlice(t *testing.T) {
	students := createStudents()
	timer := TimeUtils.NewTimer()
	reflectSlice := MapToStringSlice(students, func(i int) string {
		return students[i].name
	})
	t.Logf("MapToStringSlice cost %d nanos", timer.GetDurationInNanos())

	timer.Reset()
	nativeSlice := NativeMapToStringSlice(students)
	t.Logf("NativeMapToStringSlice cost %d nanos", timer.GetDurationInNanos())
	assert.Equal(t, nativeSlice, reflectSlice)
}
```

### iterate

- ForEach(list interface{}, action func(index int))：之前写 Java 写惯了，觉得使用 return 代替 continue 是可读性更高的代码，因为可以提早返回减少代码的嵌套接口。因此抽象出了该方法

```go
func TestForEach(t *testing.T) {
	students := createStudents()
	ForEach(students, func(index int) {
		t.Log(students[index])
	})

	ForEach(students, func(index int) {
		students[index].name = "testing"
	})

	ForEach(students, func(index int) {
		assert.Equal(t, "testing", students[index].name)
	})
}
```

### 注意

以上方法类型为 `FunctionName(list interface{}, action func(index int))` 均可以替换为  `FunctionName(len int, action func(index int))`

比如 ForEach 的实现可以为 
```go
func ForEach(len int, action func(i int)) {
	for i := 0; i < len; i++ {
		action(i)
	}
}
```

这个版本的运行效率更高。但是考虑到边界检查的安全性问题，我还是更倾向于传入 list 进来，通过反射来获取列表的长度。相关代码可查看 https://github.com/sjsdfg/common-lang-in-go/blob/faster/CollectionUtils/collection_utils.go

## Float32Utils

- Max(list ...float32) float32
- Min(list ...float32) float32
- If(condition bool, ifTrue, ifFalse float32) float32
- Abs(a float32) float32
- Equal(a, b float32) bool

## Float64Utils

- Max(list ...float64) float64
- Min(list ...float64) float64
- If(condition bool, ifTrue, ifFalse float64) float64
- Abs(a float64) float64
- Equal(a, b float64) bool

## Int64Utils

- Max(list ...int64) int64
- Min(list ...int64) int64
- If(condition bool, ifTrue, ifFalse int64) int64
- Abs(a int64) int64

## IntUtils

- Max(list ...int) int
- Min(list ...int) int
- If(condition bool, ifTrue, ifFalse int64) int
- Abs(a int) int

## TimeUtils

- Max(list ...time.Time) time.Time
- Min(list ...time.Time) time.Time
- If(condition bool, ifTrue, ifFalse int64) time.Time

上面的 `XxxUtils` 根据类型进行抽象，提供了最大值最小值的方法。并且因为 go 语言没有提供三目表达式，提供了 `If` 方法出来。可以边写出来可读性更高的代码

## 最佳实践

// 待实现
主要是想配合 IDEA 的 live template 和 Postfix completion 进行快捷键的搭配。

这样可以更好的使用该库。

## Benchmark

### CollectionUtils

#### CollectionUtils#IsEmpty
```shell script
goos: darwin
goarch: amd64
pkg: github.com/sjsdfg/common-lang-in-go/CollectionUtils
BenchmarkIsEmpty                    	22247214	        54.3 ns/op
BenchmarkIsEmpty-2                  	23466332	        51.7 ns/op
BenchmarkIsEmpty-4                  	23064122	        51.4 ns/op
BenchmarkIsEmpty-8                  	23634132	        51.5 ns/op
BenchmarkNativeIsEmpty
BenchmarkNativeIsEmpty              	20717202	        53.7 ns/op
BenchmarkNativeIsEmpty-2            	23106975	        50.9 ns/op
BenchmarkNativeIsEmpty-4            	22869487	        51.1 ns/op
BenchmarkNativeIsEmpty-8            	23977856	        51.1 ns/op
PASS
```

#### CollectionUtils#MapToStringSlice

```shell script

BenchmarkMapToStringSlice
BenchmarkMapToStringSlice           	 4942840	       292 ns/op
BenchmarkMapToStringSlice-2         	 4770552	       263 ns/op
BenchmarkMapToStringSlice-4         	 4740819	       302 ns/op
BenchmarkMapToStringSlice-8         	 5007906	       238 ns/op
BenchmarkNativeMapToStringSlice
BenchmarkNativeMapToStringSlice     	10728244	       113 ns/op
BenchmarkNativeMapToStringSlice-2   	10394302	       113 ns/op
BenchmarkNativeMapToStringSlice-4   	10965626	       112 ns/op
BenchmarkNativeMapToStringSlice-8   	10831770	       113 ns/op
PASS
```
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
 
```go
func TestIsEmpty(t *testing.T) {
	assert.Equal(t, true, IsEmpty(Empty))
}

func TestIsNotEmpty(t *testing.T) {
	testCase := "test"
	assert.Equal(t, true, IsNotEmpty(testCase))
}

func TestIsBlank(t *testing.T) {
	assert.Equal(t, true, IsBlank(" "))
	assert.Equal(t, true, IsBlank("	"))
}

func TestIsNotBlank(t *testing.T) {
	assert.Equal(t, true, IsNotBlank(" 123"))
	assert.Equal(t, false, IsNotBlank(" "))
}

func TestIsAllEmpty(t *testing.T) {
	assert.Equal(t, true, IsAllEmpty("", ""))
	assert.Equal(t, false, IsAllEmpty("", "123"))
}

func TestIsAnyEmpty(t *testing.T) {
	assert.Equal(t, true, IsAnyEmpty("", "123"))
	assert.Equal(t, false, IsAnyEmpty("123", "324"))
}

func TestIsAnyNoneEmpty(t *testing.T) {
	assert.Equal(t, true, IsAnyNoneEmpty("", "123"))
	assert.Equal(t, false, IsAnyNoneEmpty())
	assert.Equal(t, false, IsAnyNoneEmpty("", ""))
}

func TestEqual(t *testing.T) {
	assert.Equal(t, true, Equal("123", "123"))
	assert.Equal(t, true, Equal("abc", "abc"))
	assert.Equal(t, false, Equal("abc", "Abc"))
}

func TestEqualIgnoreCase(t *testing.T) {
	assert.Equal(t, true, EqualIgnoreCase("abc", "abc"))
	assert.Equal(t, true, EqualIgnoreCase("abc", "Abc"))
}

func TestEqualsAny(t *testing.T) {
	assert.Equal(t, true, EqualsAny("123", "123", "345", "abc"))
	assert.Equal(t, false, EqualsAny("123", "345", "abc"))
	assert.Equal(t, false, EqualsAny("abc", "345", "Abc"))
}

func TestEqualsAnyIgnoreCase(t *testing.T) {
	assert.Equal(t, true, EqualsAnyIgnoreCase("123", "123", "345", "abc"))
	assert.Equal(t, false, EqualsAnyIgnoreCase("123", "345", "abc"))
	assert.Equal(t, true, EqualsAnyIgnoreCase("abc", "345", "Abc"))
}

func TestIsDigital(t *testing.T) {
	assert.Equal(t, true, IsDigital("123456"))
	assert.Equal(t, false, IsDigital("123asdasd"))
}

func TestDefaultIfEmpty(t *testing.T) {
	var str = "123"
	assert.Equal(t, str, DefaultIfEmpty("", str))
	assert.Equal(t, str, DefaultIfEmpty(str, ""))
}

func TestIsZero(t *testing.T) {
	assert.Equal(t, true, IsZero(Empty))
	assert.Equal(t, true, IsZero("0"))
	assert.Equal(t, false, IsZero("2"))
}

func TestIsAllZero(t *testing.T) {
	assert.Equal(t, true, IsAllZero("", ""))
	assert.Equal(t, true, IsAllZero("", "0"))
	assert.Equal(t, false, IsAllZero("", "123"))
}

func TestTruncate(t *testing.T) {
	testCase := "0123456789"
	assert.Equal(t, "012", Truncate(testCase, 0, 3))
	assert.Equal(t, "012", Truncate(testCase, -5, 3))
	assert.Equal(t, testCase, Truncate(testCase, -5, 20))
}
```

## CollectionUtils

```go

```

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
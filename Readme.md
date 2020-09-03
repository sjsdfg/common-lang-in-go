## common-lang-in-go

[![API Reference](
https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667
)](https://pkg.go.dev/github.com/sjsdfg/common-lang-in-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/ethereum/go-ethereum)](https://goreportcard.com/report/github.com/sjsdfg/common-lang-in-go)

## how to use

```shell script
go get github.com/sjsdfg/common-lang-in-go
```

## Benchmark

### CollectionUtils

```shell script
goos: darwin
goarch: amd64
pkg: github.com/sjsdfg/common-lang-in-go/CollectionUtils
BenchmarkIsEmpty
BenchmarkIsEmpty           	1000000000	         0.000074 ns/op
BenchmarkIsEmpty           	1000000000	         0.000119 ns/op
BenchmarkIsEmpty           	1000000000	         0.000081 ns/op
BenchmarkIsEmpty-2         	1000000000	         0.000134 ns/op
BenchmarkIsEmpty-2         	1000000000	         0.000055 ns/op
BenchmarkIsEmpty-2         	1000000000	         0.000103 ns/op
BenchmarkIsEmpty-4         	1000000000	         0.000061 ns/op
BenchmarkIsEmpty-4         	1000000000	         0.000106 ns/op
BenchmarkIsEmpty-4         	1000000000	         0.000053 ns/op
BenchmarkIsEmpty-8         	1000000000	         0.000051 ns/op
BenchmarkIsEmpty-8         	1000000000	         0.000114 ns/op
BenchmarkIsEmpty-8         	1000000000	         0.000050 ns/op
BenchmarkNativeIsEmpty
BenchmarkNativeIsEmpty     	1000000000	         0.000113 ns/op
BenchmarkNativeIsEmpty     	1000000000	         0.000059 ns/op
BenchmarkNativeIsEmpty     	1000000000	         0.000054 ns/op
BenchmarkNativeIsEmpty-2   	1000000000	         0.000061 ns/op
BenchmarkNativeIsEmpty-2   	1000000000	         0.000049 ns/op
BenchmarkNativeIsEmpty-2   	1000000000	         0.000065 ns/op
BenchmarkNativeIsEmpty-4   	1000000000	         0.000049 ns/op
BenchmarkNativeIsEmpty-4   	1000000000	         0.000059 ns/op
BenchmarkNativeIsEmpty-4   	1000000000	         0.000061 ns/op
BenchmarkNativeIsEmpty-8   	1000000000	         0.000113 ns/op
BenchmarkNativeIsEmpty-8   	1000000000	         0.000062 ns/op
BenchmarkNativeIsEmpty-8   	1000000000	         0.000054 ns/op
PASS
```
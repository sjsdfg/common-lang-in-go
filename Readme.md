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
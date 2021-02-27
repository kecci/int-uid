![GitHub](https://img.shields.io/github/license/kecci/int-uid)
[![Go Report Card](https://goreportcard.com/badge/github.com/kecci/int-uid)](https://goreportcard.com/report/github.com/kecci/int-uid)
![GitHub issues](https://img.shields.io/github/issues/kecci/int-uid)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/kecci/int-uid)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fkecci%2Fint-uid.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fkecci%2Fint-uid?ref=badge_shield)
[![CircleCI](https://circleci.com/gh/kecci/int-uid.svg?style=svg)](https://circleci.com/gh/kecci/int-uid)
# int-uid
<p align="center">
<img src="asset/int-uid.png" alt="Kitten"
	title="A cute kitten" width="200" height="200" />
</p>

Experimental project to generate integer unique identifier.

Docs: https://pkg.go.dev/github.com/kecci/int-uid

## Installation
```sh
$ go get github.com/kecci/int-uid
```

## Usage

examples main.go:

```go
package main

import (
	"fmt"

	uid "github.com/kecci/int-uid"
)

func main() {
	// Unix Nano
	fmt.Println(uid.New().UnixNano().Int64())
	fmt.Println(uid.New().UnixNano().Reverse().Int64())

	// Snowflake
	fmt.Println(uid.New().Snowflake().Int64())
	fmt.Println(uid.New().Snowflake().Reverse().Int64())
}
```

output :
```sh
1614432967521585000
1261257692344161
1365656963861450752
2570541683696565631
```

## Data Types
There are several functions to get int-uid.


| Func | Bit | Length | Go Data Type | SQL Data Type | 
| ---- | --- | ------ | ------------ | ------------- |
| UnixNano() | 64 bit | 19 Character | int64 | BigInt |
| UnixNanoReverse() | 64 bit | 16 Character | int64 | BigInt |

## Benchmark
```sh
goos: darwin
goarch: amd64
pkg: github.com/kecci/int-uid
BenchmarkUnixNano/UnixNano()-8                 1000000000     0.000002 ns/op      0 B/op      0 allocs/op
BenchmarkUnixNano/UnixNano().Reverse()-8       1000000000     0.000006 ns/op      0 B/op      0 allocs/op
BenchmarkSnowflake/Snowflake()-8               1000000000     0.000002 ns/op      0 B/op      0 allocs/op
BenchmarkSnowflake/Snowflake().Reverse()-8     1000000000     0.000002 ns/op      0 B/op      0 allocs/op
PASS
coverage: 100.0% of statements
ok      github.com/kecci/int-uid        0.287s
```

## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fkecci%2Fint-uid.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fkecci%2Fint-uid?ref=badge_large)
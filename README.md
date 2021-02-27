![GitHub](https://img.shields.io/github/license/kecci/int-uid)
[![Go Report Card](https://goreportcard.com/badge/github.com/kecci/int-uid)](https://goreportcard.com/report/github.com/kecci/int-uid)
![GitHub issues](https://img.shields.io/github/issues/kecci/int-uid)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/kecci/int-uid)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fkecci%2Fint-uid.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fkecci%2Fint-uid?ref=badge_shield)
[![CircleCI](https://circleci.com/gh/kecci/int-uid.svg?style=svg)](https://circleci.com/gh/kecci/int-uid)
# int-uid
<p align="center">
<img src="asset/int-uid.png" alt="int-uid"
	title="int-uid" width="200" height="200" />
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
	"time"

	uid "github.com/kecci/int-uid"
)

func main() {
	// Unix Nano
	fmt.Println("UnixNano")
	unix := uid.New().UnixNano()
	unixReverse := unix.Reverse()
	fmt.Println(unix.Int64())
	fmt.Println(unixReverse.Int64())

	// Snowflake
	fmt.Println("Snowflake")
	snow := uid.New().Snowflake()
	snowReverse := snow.Reverse()
	fmt.Println(snow.Int64())
	fmt.Println(snowReverse.Int64())

	// Sonyflake
	fmt.Println("Sonyflake")
	sony := uid.New().Sonyflake(time.Time{})
	sonyReverse := sony.Reverse()
	fmt.Println(sony.Int64())
	fmt.Println(sonyReverse.Int64())
}

```

output :
```sh
UnixNano
1614436846108942000
2498016486344161
Snowflake
1365673231834419200
29144381323765631
Sonyflake
3999050837534991611
1161994357380509993
```

## Data Types
There are several functions to get int-uid.


| Func | Bit | Length | Go Data Type | SQL Data Type | Example |
| ---- | --- | ------ | ------------ | ------------- | ------- |
| UnixNano() | 64 bit | 19 Character | int64 | BigInt | 1614432967521585000 |
| UnixNano().Reverse() | 64 bit | 16 Character | int64 | BigInt | 1261257692344161 |
| Snowflake() | 64 bit | 19 Character | int64 | BigInt | 1365656963861450752 |
| Snowflake().Reverse() | 64 bit | 19 Character | int64 | BigInt | 2570541683696565631 |
| Sonyflake(time.Time) | 64 bit | 19 Character | int64 | BigInt | 3999050837534991611 |
| Sonyflake(time.Time).Reverse() | 64 bit | 19 Character | int64 | BigInt | 1161994357380509993 |

## Benchmark
```sh
goos: darwin
goarch: amd64
pkg: github.com/kecci/int-uid
BenchmarkUnixNano/UnixNano()-8          		1000000000      0.000001 ns/op      0 B/op      0 allocs/op
BenchmarkUnixNano/UnixNano().Reverse()-8        1000000000      0.000005 ns/op      0 B/op      0 allocs/op
BenchmarkSnowflake/Snowflake()-8                1000000000      0.000002 ns/op      0 B/op      0 allocs/op
BenchmarkSnowflake/Snowflake().Reverse()-8      1000000000      0.000002 ns/op      0 B/op      0 allocs/op
BenchmarkSonyflake/Sonyflake()-8                1000000000      0.000049 ns/op      0 B/op      0 allocs/op
BenchmarkSonyflake/Sonyflake().Reverse()-8      1000000000      0.000047 ns/op      0 B/op      0 allocs/op
PASS
coverage: 100.0% of statements
ok      github.com/kecci/int-uid        0.179s
```

## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fkecci%2Fint-uid.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fkecci%2Fint-uid?ref=badge_large)
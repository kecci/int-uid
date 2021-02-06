![GitHub](https://img.shields.io/github/license/kecci/int-uid)
[![Go Report Card](https://goreportcard.com/badge/github.com/kecci/int-uid)](https://goreportcard.com/report/github.com/kecci/int-uid)
![GitHub issues](https://img.shields.io/github/issues/kecci/int-uid)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/kecci/int-uid)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
![GitHub followers](https://img.shields.io/github/followers/kecci?style=social)
![GitHub Repo stars](https://img.shields.io/github/stars/kecci/int-uid?style=social)

# int-uid

Experimental project to generate integer unique identifier.

Docs: https://pkg.go.dev/github.com/kecci/int-uid

## Installation
```sh
$ go get github.com/kecci/int-uid
```

## Usage

main.go:

```go
package main

import (
    "github.com/kecci/int-uid"
)

func main() {
    u := uid.UnixNano()
    ur := uid.UnixNanoReverse()

    fmt.Println(u)
    fmt.Println(ur)
}
```

output :
```sh
1611310691684907000
2858231800752161
```

## Data Types
There are several functions to get int-uid.


| Func | Bit | Length | Go Data Type | SQL Data Type | 
| ---- | --- | ------ | ------------ | ------------- |
| UnixNano() | 64 bit | 19 Character | int64 | BigInt |
| UnixNanoReverse() | 64 bit | 16 Character | int64 | BigInt |

## Benchmark
```sh
BenchmarkUnixNano-8          	1000000000	         0.000001 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnixNanoReverse-8   	1000000000	         0.000007 ns/op	       0 B/op	       0 allocs/op
PASS
coverage: 100.0% of statements
```

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fkecci%2Fint-uid.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fkecci%2Fint-uid?ref=badge_large)
![GitHub](https://img.shields.io/github/license/kecci/int-uid)
[![Go Report Card](https://goreportcard.com/badge/github.com/kecci/int-uid)](https://goreportcard.com/report/github.com/kecci/int-uid)
![GitHub issues](https://img.shields.io/github/issues/kecci/int-uid)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/kecci/int-uid)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
![GitHub followers](https://img.shields.io/github/followers/kecci?style=social)
![GitHub Repo stars](https://img.shields.io/github/stars/kecci/int-uid?style=social)

# int-uid

Experimental project to generate integer unique identifier.

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
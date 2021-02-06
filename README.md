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
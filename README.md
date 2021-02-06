# int-uid

Experimental project to generate integer unique identifier.

## Unix Nano


| Bit | Length | Go Data Type | SQL Data Type | 
| --- | ------ | ------------ | ------------- |
| 64 bit | 19 Character | int64 | BigInt |

Run:
```sh
$ go run unixnano/main.go
```

Output :
```sh
1611310691684907000
```

##  Unix Nano (Reverse)

| Bit | Length | Go Data Type | SQL Data Type | 
| --- | ------ | ------------ | ------------- |
| 64 bit | 16 Character | int64 | BigInt |

Run:
```sh
$ go run unixnano-reverse/main.go
```

Output :
```sh
2858231800752161
```

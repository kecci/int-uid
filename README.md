# int-uid

Experimental project to generate integer unique identifier.

## Unix Nano


| Length | Go Data Type | SQL Data Type | 
| ------ | ------------ | ------------- |
| 19 Character | int64 | BigInt |

Run:
```sh
$ go run unixnano/main.go
```

Output :
```sh
1611310691684907000
```

##  Unix Nano (Reverse)

| Length | Go Data Type | SQL Data Type | 
| ------ | ------------ | ------------- |
| 16 Character | int64 | BigInt |

Run:
```sh
$ go run unixnano-reverse/main.go
```

Output :
```sh
2858231800752161
```

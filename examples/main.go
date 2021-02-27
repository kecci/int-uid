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
	sony := uid.New().Sonyflake(time.Date(1945, 8, 17, 10, 0, 0, 0, time.UTC))
	sonyReverse := sony.Reverse()
	fmt.Println(sony.Int64())
	fmt.Println(sonyReverse.Int64())
}

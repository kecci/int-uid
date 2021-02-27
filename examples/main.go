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
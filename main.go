package main

import (
	"fmt"
	"time"
)

func main() {
	var id int
	idNano := time.Now().UnixNano()

	id = int(idNano)
	
	// Unix Nano
	fmt.Println("time.Now().UnixNano()")
	fmt.Printf("int64 : %d\n", idNano)
	fmt.Printf("int: %d", id)
}

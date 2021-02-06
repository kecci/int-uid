package uid

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// UnixNano generate unixnano
func UnixNano() int64 {
	mu := sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()
	return time.Now().UnixNano()
}

// UnixNanoReverse generate reverse of unixnano
func UnixNanoReverse() int64 {
	id, _ := strconv.Atoi(reverse(fmt.Sprintf("%d", UnixNano())))
	return int64(id)
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

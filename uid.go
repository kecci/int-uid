package uid

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
)

// UID basic info
type UID struct {
	mu sync.Mutex
}

// New initiate UID
func New() *UID {
	return &UID{}
}

// ID custom type used for int-uid
type ID int64

// UnixNano generate unixnano id
func (u *UID) UnixNano() ID {
	u.mu.Lock()
	defer u.mu.Unlock()
	return ID(time.Now().UnixNano())
}

// Snowflake generate snowflake id
func (u *UID) Snowflake() ID {
	u.mu.Lock()
	defer u.mu.Unlock()
	node, _ := snowflake.NewNode(1)
	return ID(node.Generate().Int64())
}

// Reverse ID
func (i ID) Reverse() ID {
	s := fmt.Sprintf("%d", int64(i))
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	id, _ := strconv.Atoi(string(rns))
	return ID(id)
}

// Int64 convert to int64
func (i ID) Int64() int64 {
	return int64(i)
}

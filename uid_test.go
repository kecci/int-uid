package uid_test

import (
	"testing"

	uid "github.com/kecci/int-uid"
	"github.com/stretchr/testify/assert"
)

func TestUnixNano(t *testing.T) {
	t.Run("Not Nil", func(t *testing.T) {
		u := uid.UnixNano()
		assert.NotNil(t, u)
	})

	t.Run("Identical Checker", func(t *testing.T) {
		ch := make(chan int64)
		var u1 int64
		var u2 int64
		for i := 0; i <= 1000000; i++ {
			go func() {
				ch <- uid.UnixNano()
			}()
			u1 = <-ch
			assert.NotEqualValues(t, &u1, &u2)
			u2 = u1
		}
	})
}

func TestUnixNanoReverse(t *testing.T) {
	t.Run("Not Nil", func(t *testing.T) {
		u := uid.UnixNanoReverse()
		assert.NotNil(t, u)
	})

	t.Run("Identical Checker", func(t *testing.T) {
		ch := make(chan int64)
		var u1 int64
		var u2 int64
		for i := 0; i <= 1000000; i++ {
			go func() {
				ch <- uid.UnixNanoReverse()
			}()
			u1 = <-ch
			assert.NotEqualValues(t, &u1, &u2)
			u2 = u1
		}
	})
}

func BenchmarkUnixNano(b *testing.B) {
	u := uid.UnixNano()
	assert.NotNil(b, u)
}

func BenchmarkUnixNanoReverse(b *testing.B) {
	u := uid.UnixNanoReverse()
	assert.NotNil(b, u)
}

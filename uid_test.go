package uid_test

import (
	"testing"

	uid "github.com/kecci/int-uid"
	"github.com/stretchr/testify/assert"
)

/*
Unit Test
*/
func TestUnixNano(t *testing.T) {
	t.Run("Not Nil", func(t *testing.T) {
		u := uid.New().UnixNano().Int64()
		assert.NotNil(t, u)
	})

	t.Run("Identical Checker", func(t *testing.T) {
		ch := make(chan int64)
		var u1 int64
		var u2 int64
		for i := 0; i <= 1000; i++ {
			go func() {
				ch <- uid.New().UnixNano().Int64()
			}()
			u1 = <-ch
			assert.True(t, &u1 != &u2)
			u2 = u1
		}
	})

	t.Run("Reverse Not Nil", func(t *testing.T) {
		u := uid.New().UnixNano().Reverse().Int64()
		assert.NotNil(t, u)
	})

	t.Run("Reverse Identical Checker", func(t *testing.T) {
		ch := make(chan int64)
		var u1 int64
		var u2 int64
		for i := 0; i <= 1000; i++ {
			go func() {
				ch <- uid.New().UnixNano().Reverse().Int64()
			}()
			u1 = <-ch
			assert.True(t, &u1 != &u2)
			u2 = u1
		}
	})
}

func TestSnowflake(t *testing.T) {
	t.Run("Not nil", func(t *testing.T) {
		u := uid.New().Snowflake().Int64()
		assert.NotNil(t, u)
	})

	t.Run("Identical Checker", func(t *testing.T) {
		ch := make(chan int64)
		var u1 int64
		var u2 int64
		for i := 0; i <= 1000; i++ {
			go func() {
				ch <- uid.New().Snowflake().Int64()
			}()
			u1 = <-ch
			assert.True(t, &u1 != &u2)
			u2 = u1
		}
	})

	t.Run("Reverse Not nil", func(t *testing.T) {
		u := uid.New().Snowflake().Reverse().Int64()
		assert.NotNil(t, u)
	})

	t.Run("Reverse Identical Checker", func(t *testing.T) {
		ch := make(chan int64)
		var u1 int64
		var u2 int64
		for i := 0; i <= 1000; i++ {
			go func() {
				ch <- uid.New().Snowflake().Reverse().Int64()
			}()
			u1 = <-ch
			assert.True(t, &u1 != &u2)
			u2 = u1
		}
	})
}

/*
Benchmark Test
*/
func BenchmarkUnixNano(b *testing.B) {
	b.Run("UnixNano()", func(b *testing.B) {
		u := uid.New().UnixNano().Int64()
		assert.NotNil(b, u)
	})

	b.Run("UnixNano().Reverse()", func(b *testing.B) {
		u := uid.New().UnixNano().Reverse().Int64()
		assert.NotNil(b, u)
	})
}

func BenchmarkSnowflake(b *testing.B) {
	b.Run("Snowflake()", func(b *testing.B) {
		u := uid.New().Snowflake().Int64()
		assert.NotNil(b, u)
	})

	b.Run("Snowflake().Reverse()", func(b *testing.B) {
		u := uid.New().Snowflake().Int64()
		assert.NotNil(b, u)
	})
}

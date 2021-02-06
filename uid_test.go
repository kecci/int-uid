package uid_test

import (
	"testing"

	uid "github.com/kecci/int-uid"
	"github.com/stretchr/testify/assert"
)

func TestUnixNano(t *testing.T) {
	u := uid.UnixNano()
	assert.NotNil(t, u)
}

func TestUnixNanoReverse(t *testing.T) {
	u := uid.UnixNanoReverse()
	assert.NotNil(t, u)
}

func BenchmarkUnixNano(b *testing.B) {
	u := uid.UnixNano()
	assert.NotNil(b, u)
}

func BenchmarkUnixNanoReverse(b *testing.B) {
	u := uid.UnixNanoReverse()
	assert.NotNil(b, u)
}

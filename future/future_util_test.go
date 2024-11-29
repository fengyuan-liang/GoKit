package future

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func successfulFuture() int {
	time.Sleep(100 * time.Millisecond)
	return 42
}

func failingFuture() (int, error) {
	time.Sleep(100 * time.Millisecond)
	return 0, errors.New("failed")
}

func TestAllSuccess(t *testing.T) {
	f1 := FutureFunc[int](successfulFuture)
	f2 := FutureFunc[int](successfulFuture)

	allSuccess, err := AllSuccess(f1, f2)
	assert.NoError(t, err)
	assert.True(t, allSuccess)

	f3 := FutureFunc[int](failingFuture)

	allSuccess, err = AllSuccess(f1, f2, f3)
	assert.Error(t, err)
	assert.False(t, allSuccess)
}

func TestAnySuccess(t *testing.T) {
	f1 := FutureFunc[int](successfulFuture)
	f2 := FutureFunc[int](failingFuture)

	anySuccess, err := AnySuccess(f1, f2)
	assert.NoError(t, err)
	assert.True(t, anySuccess)

	f3 := FutureFunc[int](failingFuture)

	anySuccess, err = AnySuccess(f2, f3)
	assert.Error(t, err)
	assert.False(t, anySuccess)
}

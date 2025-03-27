package pointer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPointer(t *testing.T) {
	assert.True(t, IsPointer(New(1)))
	assert.False(t, IsPointer(1))
}

func TestSafe(t *testing.T) {
	var v *int

	assert.Equal(t, 0, Safe(v))

	v = New(1)
	assert.Equal(t, 1, Safe(v))
}

func TestNew(t *testing.T) {
	v := New(1)
	assert.Equal(t, 1, *v)
}

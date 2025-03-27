package ectolinq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTernary(t *testing.T) {
	assert.Equal(t, 1, Ternary(true, 1, 2))
	assert.Equal(t, 2, Ternary(false, 1, 2))
}

func TestDefault(t *testing.T) {
	assert.Equal(t, 1, Default(1, 2))
	assert.Equal(t, 2, Default(0, 2))
}

func TestIsZero(t *testing.T) {
	assert.True(t, IsZero(0))
	assert.True(t, IsZero(""))
	assert.True(t, IsZero(false))
	assert.False(t, IsZero(1))
}

package main

import (
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)

	a, b := alphabeticSort("a", "b")
	assert.Equal("a", a)
	assert.Equal("b", b)

	a, b = alphabeticSort("b", "a")
	assert.Equal("a", a)
	assert.Equal("b", b)

	assert.True(true)
}

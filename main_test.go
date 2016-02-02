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

	r := Routes{}

	r.Increment("a", "b", 1)
	r.Increment("b", "a", 2)
	r.Increment("b", "c", 1)

	assert.Equal(Routes{
		Path: map[string]map[string]int{
			"a": map[string]int{
				"b": 3,
			},
			"b": map[string]int{
				"c": 1,
			},
		},
	}, r)

	r.Increment("a", "b", -1)

	assert.Equal(2, r.Path["a"]["b"])

	assert.True(true)
}

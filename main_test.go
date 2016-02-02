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
		Build: []Path{Path{A: "a", B: "b", Count: 1}, Path{A: "a", B: "b", Count: 2}, Path{A: "b", B: "c", Count: 1}},
	}, r)

	r1 := Routes{}

	r1.Increment("a", "b", 1)

	ok, count := r1.Traverse("ba")
	assert.True(ok)
	assert.Equal(1, count)

	r2 := Routes{}

	r2.Increment("a", "b", 1)

	ok, count = r2.Traverse("aba")
	assert.False(ok)
	assert.Equal(0, count)

	r3 := Routes{}

	r3.Increment("a", "b", 1)
	r3.Increment("b", "c", 3)
	r3.Increment("c", "d", 3)
	r3.Increment("b", "d", 2)
	r3.Increment("d", "e", 1)

	ok, count = r3.Traverse("ab")
	assert.True(ok)
	assert.Equal(1, count)

	r4 := Routes{}

	r4.Increment("a", "b", 1)
	r4.Increment("b", "c", 3)
	r4.Increment("c", "d", 3)
	r4.Increment("b", "d", 2)
	r4.Increment("d", "e", 1)

	ok, count = r4.Traverse("abcdbcdbcde")
	assert.True(ok)
	assert.Equal(1, count)

	r5 := Routes{}

	r5.Increment("a", "b", 1)
	r5.Increment("b", "c", 3)
	r5.Increment("c", "d", 3)
	r5.Increment("b", "d", 2)
	r5.Increment("d", "e", 1)

	ok, count = r5.Traverse("abdbde")
	assert.False(ok)
	assert.Equal(0, count)

	assert.True(true)
}

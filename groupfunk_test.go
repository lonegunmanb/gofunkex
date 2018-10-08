package gofunkex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsMap(t *testing.T) {
	sampleGroup := make(map[int][]int)
	falseGroup := make(map[int]int)
	assert.True(t, isGroup(sampleGroup))
	assert.False(t, isGroup([]int{}))
	assert.False(t, isGroup(""))
	assert.False(t, isGroup(falseGroup))
}

func Test_Sums(t *testing.T) {
	input := make(map[string][]int)
	input["a"] = []int{2, 4}
	input["b"] = []int{1, 3}
	groupFunk := NewGroupFunk(input)
	sums := groupFunk.Sums().Map.(map[string]float64)
	assert.Equal(t, 2, len(sums))
	assert.Equal(t, float64(6), sums["a"])
	assert.Equal(t, float64(4), sums["b"])
}

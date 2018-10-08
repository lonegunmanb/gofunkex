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

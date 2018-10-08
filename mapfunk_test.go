package gofunkex

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_MapElem(t *testing.T) {
	input := getMap()
	mapFunk := NewMapFunk(input)
	stringMap := mapFunk.MapElem(func(i int) string { return strconv.Itoa(i) }).Map.(map[string]string)
	assert.Equal(t, 2, len(stringMap))
	assert.Equal(t, "1", stringMap["a"])
	assert.Equal(t, "2", stringMap["b"])
}

func getMap() map[string]int {
	resultMap := make(map[string]int)
	resultMap["a"] = 1
	resultMap["b"] = 2
	return resultMap
}

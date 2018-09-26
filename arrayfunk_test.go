package gofunkex

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
)

func Test_ArrayShouldPassCheck(t *testing.T) {
	var arr [10]int
	actual := isArray(arr)
	assert.True(t, actual)
}

func Test_SliceShouldPassCheck(t *testing.T) {
	var arr []int
	actual := isArray(arr)
	assert.True(t, actual)
}

func Test_NonArrayOrSliceShouldFailedCheck(t *testing.T) {
	var notArr struct{}
	actual := isArray(notArr)
	assert.False(t, actual)
}

func Test_NilShouldFailedCheck(t *testing.T) {
	var null interface{} = nil
	actual := isArray(null)
	assert.False(t, actual)
}

func Test_SimpleMap(t *testing.T) {
	arr := []int{1, 2}
	expectedArr := []string{"1", "2"}
	funk := NewArrayFunk(arr)
	stringFunk := funk.Map(func(i int) string { return strconv.Itoa(i) })
	actualArr := stringFunk.Arr.([]string)
	assert.True(t, reflect.DeepEqual(expectedArr, actualArr))
}

func Test_CheckPredicateType_Nil_Predicate(t *testing.T) {
	arr := []int{1}
	assert.Panics(t, nil, arr)
}

func Test_CheckPredicateType_Receive_NonFunc(t *testing.T) {
	arr := []int{1}
	assert.Panics(t, func() {
		checkPredicateType(1, arr)
	})
}

func Test_CheckPredicateType_Return_Not_Bool(t *testing.T) {
	arr := []int{1}
	assert.Panics(t, func() {
		checkPredicateType(func(i int) int { return i }, arr)
	})
}

func Test_CheckPredicateType_IncompatiblePredicateType(t *testing.T) {
	arr := []int{1}
	assert.Panics(t, func() {
		checkPredicateType(func(s string) bool { return true }, arr)
	})
}

func Test_Any(t *testing.T) {
	arr := []int{1, 2, 3}
	arrFunk := NewArrayFunk(arr)
	assert.True(t, arrFunk.AnyMeets(func(i int) bool { return i%2 == 0 }))
	assert.False(t, arrFunk.AnyMeets(func(i int) bool { return i > 3 }))
}

func Test_All(t *testing.T) {
	arr := []int{1, 2, 3}
	arrFunk := NewArrayFunk(arr)
	assert.False(t, arrFunk.AllMeets(func(i int) bool { return i%2 == 0 }))
	assert.True(t, arrFunk.AllMeets(func(i int) bool { return i <= 3 }))
}

func Test_Filter(t *testing.T) {
	arr := []int{1, 2, 3}
	expected := []int{2}
	arrFunk := NewArrayFunk(arr)
	filteredFunk := arrFunk.Filter(func(i int) bool { return i%2 == 0 })
	assert.True(t, reflect.DeepEqual(expected, filteredFunk.Arr))
}

func Test_Contains(t *testing.T) {
	arr := []int{1, 2, 3}
	arrFunk := NewArrayFunk(arr)
	assert.True(t, arrFunk.Contains(1))
	assert.False(t, arrFunk.Contains(0))
}

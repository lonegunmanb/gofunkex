package gofunkex

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
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

func Test_Distinct(t *testing.T) {
	arr := []string{"a", "b", "b"}
	expected := []string{"a", "b"}
	arrFunk := NewArrayFunk(arr)
	actual := arrFunk.Distinct().Arr
	assert.True(t, reflect.DeepEqual(expected, actual))
}

func Test_Length(t *testing.T) {
	arr := []int{1, 2, 3}
	arrFunk := NewArrayFunk(arr)
	assert.Equal(t, 3, arrFunk.Length())
}

func Test_Length_Empty_Slice(t *testing.T) {
	var arr []int
	arrFunk := NewArrayFunk(arr)
	assert.Equal(t, 0, arrFunk.Length())
}

func Test_Empty(t *testing.T) {
	arr := []int{1, 2, 3}
	arrFunk := NewArrayFunk(arr)
	emptyFunk := arrFunk.Filter(func(i int) bool { return i > 3 })
	assert.True(t, emptyFunk.Empty())
}

func Test_Empty_Slice_Should_Return_True(t *testing.T) {
	var arr []int
	assert.True(t, NewArrayFunk(arr).Empty())
}

func Test_Count(t *testing.T) {
	arr := []int{1, 2, 3}
	arrFunk := NewArrayFunk(arr)
	assert.Equal(t, 1, arrFunk.Count(func(i int) bool { return i%2 == 0 }))
}

func Test_Count_Empty_Slice_Should_Return_Zero(t *testing.T) {
	var arr []int
	assert.Equal(t, 0, NewArrayFunk(arr).Count(func(i int) bool { return true }))
}

func Test_Head(t *testing.T) {
	arr := []int{1, 2, 3}
	arrFunk := NewArrayFunk(arr)
	assert.Equal(t, 1, arrFunk.Head())
}

func Test_Head_Empty_Slice_Should_Return_Nil(t *testing.T) {
	var arr []int
	head := NewArrayFunk(arr).Head()
	assert.Nil(t, head)
}

func Test_Last(t *testing.T) {
	arr := []int{1, 2, 3}
	arrFunk := NewArrayFunk(arr)
	assert.Equal(t, 3, arrFunk.Last())
}

func Test_Last_Empty_Slice_Should_Return_Nil(t *testing.T) {
	var arr []int
	last := NewArrayFunk(arr).Last()
	assert.Nil(t, last)
}

func Test_Initial(t *testing.T) {
	arr := []int{1, 2, 3}
	expected := []int{1, 2}
	arrFunk := NewArrayFunk(arr)
	assert.True(t, reflect.DeepEqual(expected, arrFunk.Initial().Arr))
}

func Test_Initial_Empty_Slice_Should_Return_Nil(t *testing.T) {
	var arr []int
	initial := NewArrayFunk(arr).Initial()
	assert.Nil(t, initial.Arr)
}

func Test_Tail(t *testing.T) {
	arr := []int{1, 2, 3}
	expected := []int{2, 3}
	arrFunk := NewArrayFunk(arr)
	assert.True(t, reflect.DeepEqual(expected, arrFunk.Tail().Arr))
}

func Test_Tail_Empty_Slice_Should_Return_Nil(t *testing.T) {
	var arr []int
	assert.Nil(t, NewArrayFunk(arr).Tail().Arr)
}

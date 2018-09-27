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

func Test_Pass_Nil_Slice_To_New_ArrayFunk_Should_Return_Empty_Slice(t *testing.T) {
	var arr []int
	actual := NewArrayFunk(arr).Arr.([]int)
	assert.NotNil(t, actual)
	assert.Empty(t, actual)
}

func Test_SimpleMap(t *testing.T) {
	arr := []int{1, 2}
	expectedArr := []string{"1", "2"}
	funk := NewArrayFunk(arr)
	stringFunk := funk.Map(func(i int) string { return strconv.Itoa(i) })
	actualArr := stringFunk.Arr.([]string)
	assert.True(t, reflect.DeepEqual(expectedArr, actualArr))
}

func Test_Map_On_Empty_Slice(t *testing.T) {
	var arr []int
	actual := NewArrayFunk(arr).Map(func(i int) string { return strconv.Itoa(i) }).Arr.([]string)
	assertEmptyStringSlice(t, actual)
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

func Test_Any_Empty_Slice_Should_Return_False(t *testing.T) {
	var arr []int
	assert.False(t, NewArrayFunk(arr).AnyMeets(func(i int) bool { return true }))
}

func Test_All(t *testing.T) {
	arr := []int{1, 2, 3}
	arrFunk := NewArrayFunk(arr)
	assert.False(t, arrFunk.AllMeets(func(i int) bool { return i%2 == 0 }))
	assert.True(t, arrFunk.AllMeets(func(i int) bool { return i <= 3 }))
}

func Test_All_Empty_Slice_Should_Return_False(t *testing.T) {
	var arr []int
	assert.False(t, NewArrayFunk(arr).AllMeets(func(i int) bool { return true }))
}

func Test_Filter(t *testing.T) {
	arr := []int{1, 2, 3}
	expected := []int{2}
	arrFunk := NewArrayFunk(arr)
	filteredFunk := arrFunk.Filter(func(i int) bool { return i%2 == 0 })
	assert.True(t, reflect.DeepEqual(expected, filteredFunk.Arr))
}

func Test_Filter_Empty_Slice_Should_Return_Empty_Slice(t *testing.T) {
	var arr []int
	assertEmptyIntSlice(t, NewArrayFunk(arr).Filter(func(i int) bool { return true }).Arr.([]int))
}

func Test_Contains(t *testing.T) {
	arr := []int{1, 2, 3}
	arrFunk := NewArrayFunk(arr)
	assert.True(t, arrFunk.Contains(1))
	assert.False(t, arrFunk.Contains(0))
}

func Test_Contains_Empty_Slice_Should_Return_False(t *testing.T) {
	var arr []int
	assert.False(t, NewArrayFunk(arr).Contains(1))
}

func Test_Distinct(t *testing.T) {
	arr := []string{"a", "b", "b"}
	expected := []string{"a", "b"}
	arrFunk := NewArrayFunk(arr)
	actual := arrFunk.Distinct().Arr
	assert.True(t, reflect.DeepEqual(expected, actual))
}

func Test_Distinct_On_Empty_Slice_Should_Return_Empty_Slice(t *testing.T) {
	var arr []int
	assertEmptyIntSlice(t, NewArrayFunk(arr).Distinct().Arr.([]int))
}

func Test_Length(t *testing.T) {
	arr := []int{1, 2, 3}
	arrFunk := NewArrayFunk(arr)
	assert.Equal(t, 3, arrFunk.Length())
}

func Test_Length_Empty_Slice_Should_Return_Zero(t *testing.T) {
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

func Test_Initial_Empty_Slice_Should_Return_Empty(t *testing.T) {
	var arr []int
	initial := NewArrayFunk(arr).Initial()
	assertEmptyIntSlice(t, initial.Arr.([]int))
}

func Test_Tail(t *testing.T) {
	arr := []int{1, 2, 3}
	expected := []int{2, 3}
	arrFunk := NewArrayFunk(arr)
	assert.True(t, reflect.DeepEqual(expected, arrFunk.Tail().Arr))
}

func Test_Tail_Empty_Slice_Should_Return_Nil(t *testing.T) {
	var arr []int
	assertEmptyIntSlice(t, NewArrayFunk(arr).Tail().Arr.([]int))
}

func Test_Flatten(t *testing.T) {
	arr := [][]int{{1, 2}, {3, 4}}
	expected := []int{1, 2, 3, 4}
	arrFunk := NewArrayFunk(arr)
	actual := arrFunk.Flatten().Arr
	assert.True(t, reflect.DeepEqual(expected, actual))
}

func Test_Flatten_On_One_Dimension_Slice(t *testing.T) {
	arr := []int{1, 2, 3}
	actual := NewArrayFunk(arr).Flatten().Arr
	assert.True(t, reflect.DeepEqual(arr, actual))
}

func Test_Take(t *testing.T) {
	arr := []int{1, 2, 3}
	expected := arr[:2]
	assert.True(t, reflect.DeepEqual(expected, NewArrayFunk(arr).Take(2).Arr))
}

func Test_Take_More_Than_Len_Should_Return_Origin_Slice(t *testing.T) {
	arr := []int{1, 2, 3}
	expected := arr
	actual := NewArrayFunk(arr).Take(4).Arr
	assert.True(t, reflect.DeepEqual(expected, actual))
}

func Test_Take_Empty_Slice_Should_Return_Empty_Slice(t *testing.T) {
	var arr []int
	actual := NewArrayFunk(arr).Take(10).Arr.([]int)
	assertEmptyIntSlice(t, actual)
}

func Test_Take_Negative_Should_Return_Empty_Slice(t *testing.T) {
	arr := []int{1, 2, 3}
	assertEmptyIntSlice(t, NewArrayFunk(arr).Take(-1).Arr.([]int))
}

func Test_Skip(t *testing.T) {
	arr := []int{1, 2, 3}
	expected := arr[1:]
	assert.True(t, reflect.DeepEqual(expected, NewArrayFunk(arr).Skip(1).Arr))
}

func Test_Skip_More_Than_Len_Should_Return_Empty_Slice(t *testing.T) {
	arr := []int{1, 2, 3}
	actual := NewArrayFunk(arr).Skip(4).Arr.([]int)
	assertEmptyIntSlice(t, actual)
}

func Test_Skip_Empty_Slice_Should_Return_Empty_Slice(t *testing.T) {
	var arr []int
	actual := NewArrayFunk(arr).Skip(1).Arr.([]int)
	assertEmptyIntSlice(t, actual)
}

func Test_Skip_Negative_Should_Skip_Zero(t *testing.T) {
	arr := []int{1, 2, 3}
	actual := NewArrayFunk(arr).Skip(-1).Arr
	assert.True(t, reflect.DeepEqual(arr, actual))
}

func assertEmptyIntSlice(t *testing.T, slice []int) {
	assert.NotNil(t, slice)
	assert.Equal(t, 0, len(slice))
}

func assertEmptyStringSlice(t *testing.T, slice []string) {
	assert.NotNil(t, slice)
	assert.Equal(t, 0, len(slice))
}

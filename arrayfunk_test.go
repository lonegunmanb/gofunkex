package gofunkex

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
)

func Test_ArrayShouldPassCheck(t *testing.T){
	var arr [10]int
	actual := isArray(arr)
	assert.True(t, actual)
}

func Test_SliceShouldPassCheck(t *testing.T){
	var arr []int
	actual := isArray(arr)
	assert.True(t, actual)
}

func Test_NonArrayOrSliceShouldFailedCheck(t *testing.T){
	var notArr struct{}
	actual := isArray(notArr)
	assert.False(t, actual)
}

func Test_NilShouldFailedCheck(t *testing.T){
	var null interface{} = nil
	actual := isArray(null)
	assert.False(t, actual)
}

func Test_SimpleMap(t *testing.T){
	arr := []int{1,2}
	expectedArr := []string{"1", "2"}
	funk := NewArrayFunk(arr)
	stringFunk := funk.Map(func(i int) string {return strconv.Itoa(i)})
	actualArr := stringFunk.Arr.([]string)
	assert.True(t, reflect.DeepEqual(expectedArr, actualArr))
}

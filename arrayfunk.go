package gofunkex

import (
	"fmt"
	"go-funk"
	"reflect"
)

type ArrayFunk struct{
	Arr interface{}
}

func NewArrayFunk(arr interface{}) ArrayFunk{
	if !isArray(arr) {
		panic(fmt.Sprintf("non nil array or slice required, got %s-%s", arr, reflect.TypeOf(arr).Kind().String()))
	}
	return ArrayFunk{arr}
}

func isArray(arr interface{}) bool{
	if arr == nil{
		return false
	}
	kind := reflect.TypeOf(arr).Kind()
	return kind == reflect.Array || kind == reflect.Slice
}

func (something ArrayFunk) Map(mapper interface{}) ArrayFunk {
	return NewArrayFunk(funk.Map(something.Arr, mapper))
}

package gofunkex

import (
	"fmt"
	"go-funk"
	"reflect"
)

type ArrayFunk struct {
	Arr interface{}
}

func NewArrayFunk(arr interface{}) ArrayFunk {
	if !isArray(arr) {
		panic(fmt.Sprintf("non nil array or slice required, got %s-%s", arr, reflect.TypeOf(arr).Kind().String()))
	}
	return ArrayFunk{arr}
}

func isArray(arr interface{}) bool {
	if arr == nil {
		return false
	}
	kind := reflect.TypeOf(arr).Kind()
	return kind == reflect.Array || kind == reflect.Slice
}

func (something ArrayFunk) Map(mapper interface{}) ArrayFunk {
	return NewArrayFunk(funk.Map(something.Arr, mapper))
}

func (something ArrayFunk) AnyMeets(predicate interface{}) bool {
	arr := something.Arr
	checkPredicateType(predicate, arr)
	funcValue := reflect.ValueOf(predicate)
	arrValue := reflect.ValueOf(arr)
	for i := 0; i < arrValue.Len(); i++ {
		elem := arrValue.Index(i)
		result := funcValue.Call([]reflect.Value{elem})[0].Interface().(bool)
		if result {
			return true
		}
	}
	return false
}
func (something ArrayFunk) AllMeets(predicate interface{}) bool {
	arr := something.Arr
	checkPredicateType(predicate, arr)
	funcValue := reflect.ValueOf(predicate)
	arrValue := reflect.ValueOf(arr)
	for i := 0; i < arrValue.Len(); i++ {
		elem := arrValue.Index(i)
		result := funcValue.Call([]reflect.Value{elem})[0].Interface().(bool)
		if !result {
			return false
		}
	}
	return true
}

func checkPredicateType(predicate interface{}, arr interface{}) {
	if predicate == nil {
		panic("Predicate required")
	}
	if !funk.IsFunction(predicate, 1, 1) {
		panic("Second argument must be function")
	}
	funcValue := reflect.ValueOf(predicate)
	funcType := funcValue.Type()
	if funcType.Out(0).Kind() != reflect.Bool {
		panic("Return argument should be a boolean")
	}
	elementType := reflect.TypeOf(arr).Elem()
	if elementType != funcType.In(0) {
		panic("Incompatible array and predicate input type")
	}
}

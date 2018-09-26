package gofunkex

import (
	"reflect"

	"github.com/thoas/go-funk"
)

type ArrayFunk struct {
	Arr interface{}
}

func NewArrayFunk(arr interface{}) ArrayFunk {
	if !isArray(arr) {
		panic("non nil array or slice required")
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
func (something ArrayFunk) Filter(predicate interface{}) ArrayFunk {
	return NewArrayFunk(funk.Filter(something.Arr, predicate))
}

func (something ArrayFunk) Contains(item interface{}) bool {
	return funk.Contains(something.Arr, item)
}

func (something ArrayFunk) Distinct() ArrayFunk {
	arr := something.Arr
	arrValue := reflect.ValueOf(arr)
	set := make(map[interface{}]struct{})
	for i := 0; i < arrValue.Len(); i++ {
		elem := arrValue.Index(i).Interface()
		_, found := set[elem]
		if !found {
			set[elem] = struct{}{}
		}
	}
	neatSlice := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(arr).Elem()), 0, len(set))
	for key := range set {
		neatSlice = reflect.Append(neatSlice, reflect.ValueOf(key))
	}
	return NewArrayFunk(neatSlice.Interface())
}
func (something ArrayFunk) Empty() bool {
	return something.Length() == 0
}

func (something ArrayFunk) Length() int {
	arrValue := reflect.ValueOf(something.Arr)
	return arrValue.Len()
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

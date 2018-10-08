package gofunkex

import (
	"reflect"

	"github.com/thoas/go-funk"
)

type SliceFunk struct {
	Arr interface{}
}

func NewSliceFunk(arr interface{}) SliceFunk {
	if !isSlice(arr) {
		panic("non nil slice required")
	}
	arrValue := reflect.ValueOf(arr)
	if arrValue.IsNil() {
		arr = reflect.MakeSlice(arrValue.Type(), 0, 0).Interface()
	}
	return SliceFunk{arr}
}

func isSlice(arr interface{}) bool {
	if arr == nil {
		return false
	}
	kind := reflect.TypeOf(arr).Kind()
	return kind == reflect.Slice
}

func (something SliceFunk) Map(mapper interface{}) SliceFunk {
	return NewSliceFunk(funk.Map(something.Arr, mapper))
}

func (something SliceFunk) AnyMeets(predicate interface{}) bool {
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
func (something SliceFunk) AllMeets(predicate interface{}) bool {
	arr := something.Arr
	checkPredicateType(predicate, arr)
	funcValue := reflect.ValueOf(predicate)
	arrValue := reflect.ValueOf(arr)
	if arrValue.Len() == 0 {
		return false
	}
	for i := 0; i < arrValue.Len(); i++ {
		elem := arrValue.Index(i)
		result := funcValue.Call([]reflect.Value{elem})[0].Interface().(bool)
		if !result {
			return false
		}
	}
	return true
}
func (something SliceFunk) Filter(predicate interface{}) SliceFunk {
	return NewSliceFunk(funk.Filter(something.Arr, predicate))
}

func (something SliceFunk) Contains(item interface{}) bool {
	return funk.Contains(something.Arr, item)
}

func (something SliceFunk) Distinct() SliceFunk {
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
	return NewSliceFunk(neatSlice.Interface())
}
func (something SliceFunk) Empty() bool {
	return something.Length() == 0
}

func (something SliceFunk) Length() int {
	arrValue := reflect.ValueOf(something.Arr)
	return arrValue.Len()
}

func (something SliceFunk) Count(predicate interface{}) int {
	return something.Filter(predicate).Length()
}

func (something SliceFunk) Head() interface{} {
	return funk.Head(something.Arr)
}

func (something SliceFunk) Last() interface{} {
	return funk.Last(something.Arr)
}
func (something SliceFunk) Initial() SliceFunk {
	return NewSliceFunk(funk.Initial(something.Arr))
}
func (something SliceFunk) Tail() SliceFunk {
	return NewSliceFunk(funk.Tail(something.Arr))
}
func (something SliceFunk) Flatten() SliceFunk {
	return NewSliceFunk(funk.FlattenDeep(something.Arr))
}

func (something SliceFunk) Take(i int) SliceFunk {
	return NewSliceFunk(takeSlice(something, 0, i).Interface())
}

func (something SliceFunk) Skip(i int) SliceFunk {
	arrValue := reflect.ValueOf(something.Arr)
	length := arrValue.Len()

	return NewSliceFunk(takeSlice(something, i, length).Interface())
}

func (something SliceFunk) Concat(funk2 SliceFunk) SliceFunk {
	concatTypeCheck(something, funk2)
	sliceType := reflect.SliceOf(reflect.ValueOf(something.Arr).Type().Elem())
	newSlice := reflect.MakeSlice(sliceType, 0, something.Length()+funk2.Length())
	newSlice = concatTwoSlices(newSlice, something, funk2)
	return NewSliceFunk(newSlice.Interface())
}
func (something SliceFunk) GroupBy(groupKeySelector interface{}) MapFunk {
	arr := something.Arr
	if groupKeySelector == nil {
		panic("GroupKeySelector required")
	}
	if !funk.IsFunction(groupKeySelector, 1, 1) {
		panic("GroupKeySelector must be a function")
	}
	selectorType := reflect.TypeOf(groupKeySelector)
	elemType := reflect.TypeOf(arr).Elem()
	if elemType != selectorType.In(0) {
		panic("Incompatible slice and key selector type")
	}
	sliceType := reflect.TypeOf(something.Arr)
	keyType := selectorType.Out(0)
	mapType := reflect.MapOf(keyType, sliceType)
	resultMap := reflect.MakeMap(mapType)
	sliceValue := reflect.ValueOf(arr)
	selectorValue := reflect.ValueOf(groupKeySelector)
	for i := 0; i < sliceValue.Len(); i++ {
		ele := sliceValue.Index(i)
		key := selectorValue.Call([]reflect.Value{ele})[0]
		groupSlice := resultMap.MapIndex(key)
		if !groupSlice.IsValid() {
			groupSlice = reflect.MakeSlice(sliceType, 0, 1)
		}
		groupSlice = reflect.Append(groupSlice, ele)
		resultMap.SetMapIndex(key, groupSlice)
	}
	return NewMapFunk(resultMap.Interface())
}

func concatTwoSlices(newSlice reflect.Value, funk1 SliceFunk, funk2 SliceFunk) reflect.Value {
	newSlice = reflect.AppendSlice(newSlice, reflect.ValueOf(funk1.Arr))
	newSlice = reflect.AppendSlice(newSlice, reflect.ValueOf(funk2.Arr))
	return newSlice
}

func concatTypeCheck(funk1 SliceFunk, funk2 SliceFunk) {
	type1 := reflect.ValueOf(funk1.Arr).Type().Elem()
	type2 := reflect.ValueOf(funk2.Arr).Type().Elem()
	if type1 != type2 {
		panic("Cannot concat different types")
	}
}

func takeSlice(something SliceFunk, low, high int) reflect.Value {
	arrValue := reflect.ValueOf(something.Arr)
	length := arrValue.Len()
	low = adjustBoundary(low, length)
	high = adjustBoundary(high, length)
	newSlice := arrValue.Slice(low, high)
	return newSlice
}

func adjustBoundary(boundary int, length int) int {
	if boundary > length {
		boundary = length
	} else if boundary < 0 {
		boundary = 0
	}
	return boundary
}

func checkPredicateType(predicate interface{}, arr interface{}) {
	if predicate == nil {
		panic("Predicate required")
	}
	if !funk.IsFunction(predicate, 1, 1) {
		panic("Predicate must be function")
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

package gofunkex

import (
	"go-funk"
	"reflect"
)

type GroupFunk struct {
	MapFunk
}

func NewGroupFunk(mapValue interface{}) GroupFunk {
	if !isGroup(mapValue) {
		panic("non nil group(key-slice) required")
	}
	return GroupFunk{MapFunk{mapValue}}
}

func (this GroupFunk) Sums() MapFunk {
	return this.MapElem(func(slice interface{}) float64 {
		return funk.Sum(slice)
	})
}

func (this GroupFunk) MapGroup(mapper interface{}) GroupFunk {
	if !funk.IsFunction(mapper, 1, 1) {
		panic("Mapper must be a function with one slice input and return something")
	}
	mapperType := reflect.TypeOf(mapper)
	keyType := reflect.TypeOf(this.Map).Key()
	dstType := reflect.SliceOf(mapperType.Out(0))
	resultMapType := reflect.MapOf(keyType, dstType)
	resultMap := reflect.MakeMap(resultMapType)
	mapValue := reflect.ValueOf(this.Map)
	for _, key := range mapValue.MapKeys() {
		slice := mapValue.MapIndex(key)
		r := funk.Map(slice.Interface(), mapper)
		resultMap.SetMapIndex(key, reflect.ValueOf(r))
	}
	return NewGroupFunk(resultMap.Interface())
}

func isGroup(mapValue interface{}) bool {
	if !isMap(mapValue) {
		return false
	}
	mapType := reflect.TypeOf(mapValue)
	return mapType.Elem().Kind() == reflect.Slice
}

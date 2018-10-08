package gofunkex

import (
	"go-funk"
	"reflect"
)

type MapFunk struct {
	Map interface{}
}

func NewMapFunk(mapValue interface{}) MapFunk {
	if !isMap(mapValue) {
		panic("non nil map required")
	}
	return MapFunk{mapValue}
}
func (this MapFunk) MapElem(mapper interface{}) MapFunk {
	return this.mapElem(nil, mapper)
}

func (this MapFunk) mapElem(dstType reflect.Type, mapper interface{}) MapFunk {
	if !funk.IsFunction(mapper, 1, 1) {
		panic("Mapper must be a function with one slice input and return something")
	}
	mapperType := reflect.TypeOf(mapper)
	mapperValue := reflect.ValueOf(mapper)
	keyType := reflect.TypeOf(this.Map).Key()
	if dstType == nil {
		dstType = mapperType.Out(0)
	}
	resultMapType := reflect.MapOf(keyType, dstType)
	resultMap := reflect.MakeMap(resultMapType)
	mapValue := reflect.ValueOf(this.Map)
	for _, key := range mapValue.MapKeys() {
		slice := mapValue.MapIndex(key)
		dst := mapperValue.Call([]reflect.Value{slice})[0]
		resultMap.SetMapIndex(key, dst)
	}
	return NewMapFunk(resultMap.Interface())
}

func isMap(mapValue interface{}) bool {
	if mapValue == nil {
		return false
	}
	mapType := reflect.TypeOf(mapValue)
	return mapType.Kind() == reflect.Map
}

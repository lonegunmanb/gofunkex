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
func (this *MapFunk) MapElem(aggregator interface{}) MapFunk {
	if !funk.IsFunction(aggregator, 1, 1) {
		panic("Aggregator must be a function with one slice input and return something")
	}
	aggregatorType := reflect.TypeOf(aggregator)
	aggregatorValue := reflect.ValueOf(aggregator)
	keyType := reflect.TypeOf(this.Map).Key()
	resultMapType := reflect.MapOf(keyType, aggregatorType.Out(0))
	resultMap := reflect.MakeMap(resultMapType)
	mapValue := reflect.ValueOf(this.Map)
	for _, key := range mapValue.MapKeys() {
		slice := mapValue.MapIndex(key)
		aggregate := aggregatorValue.Call([]reflect.Value{slice})[0]
		resultMap.SetMapIndex(key, aggregate)
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

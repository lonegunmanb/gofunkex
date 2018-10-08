package gofunkex

import "reflect"

type MapFunk struct {
	Map interface{}
}

func NewMap(mapValue interface{}) MapFunk {
	if !isMap(mapValue) {
		panic("non nil map required")
	}
	return MapFunk{mapValue}
}
func isMap(mapValue interface{}) bool {
	if mapValue == nil {
		return false
	}
	mapType := reflect.TypeOf(mapValue)
	return mapType.Kind() == reflect.Map
}

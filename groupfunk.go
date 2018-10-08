package gofunkex

import "reflect"

type GroupFunk struct {
	Map interface{}
}

func NewGroupFunk(mapValue interface{}) GroupFunk {
	if !isGroup(mapValue) {

	}
	return GroupFunk{mapValue}
}
func isGroup(mapValue interface{}) bool {
	if mapValue == nil {
		return false
	}
	mapType := reflect.TypeOf(mapValue)
	kind := mapType.Kind()
	if kind != reflect.Map {
		return false
	}
	return mapType.Elem().Kind() == reflect.Slice
}

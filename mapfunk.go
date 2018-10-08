package gofunkex

import "reflect"

type MapFunk struct {
	Arr interface{}
}

func NewMapFunk(arr interface{}) MapFunk {
	if !isMap(arr) {

	}
	return MapFunk{arr}
}
func isMap(arr interface{}) bool {
	if arr == nil {
		return false
	}
	kind := reflect.TypeOf(arr).Kind()
	return kind == reflect.Map
}

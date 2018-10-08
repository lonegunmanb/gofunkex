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

func (this *GroupFunk) Sums() MapFunk {
	keyType := reflect.TypeOf(this.Map).Key()
	resultMapType := reflect.MapOf(keyType, reflect.TypeOf(float64(0)))
	resultMap := reflect.MakeMap(resultMapType)
	mapValue := reflect.ValueOf(this.Map)
	for _, key := range mapValue.MapKeys() {
		slice := mapValue.MapIndex(key)
		sum := funk.Sum(slice.Interface())
		resultMap.SetMapIndex(key, reflect.ValueOf(sum))
	}
	return NewMap(resultMap.Interface())
}

func isGroup(mapValue interface{}) bool {
	if !isMap(mapValue) {
		return false
	}
	mapType := reflect.TypeOf(mapValue)
	return mapType.Elem().Kind() == reflect.Slice
}

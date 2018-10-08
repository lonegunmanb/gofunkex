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
	return this.MapElem(func(slice interface{}) float64 {
		return funk.Sum(slice)
	})
}

func isGroup(mapValue interface{}) bool {
	if !isMap(mapValue) {
		return false
	}
	mapType := reflect.TypeOf(mapValue)
	return mapType.Elem().Kind() == reflect.Slice
}

package gofunkex

type MapFunk struct {
	Arr interface{}
}

func NewMapFunk(arr interface{}) MapFunk {
	return MapFunk{arr}
}

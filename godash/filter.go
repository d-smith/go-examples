package godash

import (
	"reflect"
)

type Collection []interface{}

type PropSpec struct {
	Name string
	Value interface{}
}

func (c Collection) Filter(propSpec PropSpec) Collection {
	var nc Collection
	for _, item := range c {
		if propTest(item, propSpec) == true {
			nc = append(nc, item)
		}
	}
	return nc
}

func propTest(i interface{}, propSpec PropSpec) bool {
	//Is it a struct?
	theType := reflect.TypeOf(i)
	if theType.Kind() != reflect.Interface {
		return false
	}
	return true
}

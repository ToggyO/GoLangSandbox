package v1

import "reflect"

type param struct {
	Name  string
	Type  reflect.Type
	Value reflect.Value
}

func newParamFromInterface(parameter interface{}) param {
	t := reflect.TypeOf(parameter)
	v := reflect.ValueOf(parameter)
	return param{
		Name:  t.Name(),
		Type:  t,
		Value: v,
	}
}

func newParamFromValue(value reflect.Value) param {
	t := value.Type()
	return param{
		Type:  t,
		Value: value,
		Name:  t.Name(),
	}
}

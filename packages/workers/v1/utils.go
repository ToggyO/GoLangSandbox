package v1

import (
	"errors"
	"fmt"
	"reflect"
)

func isFunction(function interface{}) error {
	fType := reflect.TypeOf(function)
	if fType == nil {
		return errors.New("Can't invoke an untyped nil")
	}

	if fType.Kind() != reflect.Func {
		return errors.New(fmt.Sprintf("Сan't invoke non-function %v (type %v)", function, fType))
	}

	return nil
}

func getFunctionInParams(fType reflect.Type) []reflect.Value {
	inParams := fType.NumIn()
	args := make([]reflect.Value, inParams)

	for i := 0; i < inParams; i++ {
		inV := fType.In(i)
		args = append(args, reflect.ValueOf(inV))
	}

	return args
}

func transformParamsToValues(params []param) []reflect.Value {
	l := len(params)
	args := make([]reflect.Value, l)

	for i := 0; i < l; i++ {
		value := params[i].Value
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}
		args[i] = value
	}

	return args
}

func getValueFromReflectValue(values []reflect.Value) []interface{} {
	l := len(values)
	args := make([]interface{}, l)

	for i := 0; i < l; i++ {
		args[i] = values[i].Interface()
	}

	return args
}

func invokeTask(task task) []reflect.Value {
	return reflect.ValueOf(task.Job.Function).Call(transformParamsToValues(task.Params))
}

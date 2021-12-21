package main

import (
	"errors"
	"fmt"
	"reflect"
)

func add(a, b interface{}) (interface{}, error) {
	typeA := reflect.TypeOf(a).Kind()
	typeB := reflect.TypeOf(b).Kind()
	if typeA != typeB {
		return nil, errors.New("Both parameters should be of the same type")
	}
	switch typeA {
	case reflect.Int:
		return a.(int) + b.(int), nil
	case reflect.Int8:
		return a.(int8) + b.(int8), nil
	case reflect.Int16:
		return a.(int16) + b.(int16), nil
	case reflect.Int32:
		return a.(int32) + b.(int32), nil
	case reflect.Int64:
		return a.(int64) + b.(int64), nil
	case reflect.Float32:
		return a.(float32) + b.(float32), nil
	case reflect.Float64:
		return a.(float64) + b.(float64), nil
	case reflect.String:
		return a.(string) + b.(string), nil
	default:
		return nil, errors.New("Parameters type is not supported")
	}
}
func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(3.5, 4.2))
	fmt.Println(add("1", "2"))
}
